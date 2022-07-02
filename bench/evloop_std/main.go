package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"

	"log"
	_ "net/http/pprof"
	_ "unsafe"

	"github.com/ii64/gouringasm"
	"github.com/ii64/gouringasm/evloop"
)

var l = log.New(os.Stdout, "http_simple ", 0)

func main() {

	ln, err := Listen("tcp", ":9002")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	var buf = make([]byte, 4096)
	var conn net.Conn
	var n int

	if true {
		for {
			err = http.Serve(ln, http.DefaultServeMux)
			if err != nil {
				continue
			}
		}
	} else {
		for {
			conn, err = ln.Accept()
			if err != nil {
				panic(err)
			}

			n, err = conn.Read(buf)
			if err != nil {
				l.Fatal(err)
			}

			h := []byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nServer: gouringasm-http\r\nConnection: close\r\nContent-Length: %d\r\n\r\n", n))

			copy(buf[len(h):], buf[:n])
			copy(buf[:len(h)], h)

			conn.Write(buf[:len(h)+n])
			conn.Close()

		}
	}
}

func Listen(network, address string) (net.Listener, error) {
	ln, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	switch ln := ln.(type) {
	case *net.TCPListener:
		f, err := ln.File()
		if err != nil {
			return nil, err
		}

		h, err := gouringasm.NewWithParams(256, &gouringasm.IoUringParams{
			Flags:        gouringasm.IORING_SETUP_SQPOLL,
			SqThreadCPU:  4,
			SqThreadIdle: 10_000,
		})
		if err != nil {
			return nil, err
		}
		ev := evloop.New(h)
		go func() {
			err := ev.Serve()
			if err != nil {
				panic(err)
			}
		}()
		return &uringTCPListener{
			ln:   ln,
			f:    f,
			r:    h,
			ev:   ev,
			lnFd: int(f.Fd()),
		}, nil
	}

	return nil, fmt.Errorf("not implemented")
}

type uringTCPListener struct {
	ln *net.TCPListener
	f  *os.File

	r  *gouringasm.Uring
	ev *evloop.Eventloop

	lnFd int
}

func (l *uringTCPListener) Accept() (net.Conn, error) {
	rsa, rsaSz, ud := l.ev.QueueAccept(l.lnFd)
	_, err := l.ev.Submit()
	if err != nil {
		return nil, err
	}
	ud.Wait()
	if ud.Res < 0 {
		return nil, syscall.Errno(-ud.Res)
	}

	_ = rsaSz

	sa, err := anyToSockaddr(rsa)
	if err != nil {
		return nil, err
	}

	raddr := sockaddrToTCP(sa)

	return &uringTCPConn{
		ln:    l,
		sa:    sa,
		fd:    int(ud.Res),
		ev:    l.ev,
		raddr: raddr,
	}, nil
}

func (l *uringTCPListener) Close() error {
	return l.f.Close()
}

func (l *uringTCPListener) Addr() net.Addr {
	return l.ln.Addr()
}

//

type uringTCPConn struct {
	ln    net.Listener
	sa    syscall.Sockaddr
	fd    int
	raddr net.Addr

	ev *evloop.Eventloop
}

func (c *uringTCPConn) Read(b []byte) (n int, err error) {
	ud := c.ev.QueueRead(c.fd, b)
	_, err = c.ev.Submit()
	if err != nil {
		return
	}
	ud.Wait()
	if ud.Res < 0 {
		err = syscall.Errno(-ud.Res)
		if err == syscall.EAGAIN {
			return c.Read(b)
		}
		return
	}
	n = int(ud.Res)
	return
}
func (c *uringTCPConn) Write(b []byte) (n int, err error) {
	ud := c.ev.QueueWrite(c.fd, b)
	_, err = c.ev.Submit()
	if err != nil {
		return
	}
	ud.Wait()
	if ud.Res < 0 {
		err = syscall.Errno(-ud.Res)
		if err == syscall.EAGAIN {
			return c.Write(b)
		}
		return
	}
	n = int(ud.Res)
	return
}

func (c *uringTCPConn) Close() (err error) {
	ud := c.ev.QueueClose(c.fd)
	_, err = c.ev.Submit()
	if err != nil {
		return
	}
	ud.Wait()
	if ud.Res < 0 {
		err = syscall.Errno(-ud.Res)
		return
	}
	return
}

func (c *uringTCPConn) LocalAddr() net.Addr {
	return c.ln.Addr()
}

func (c *uringTCPConn) RemoteAddr() net.Addr {
	return c.raddr
}

func (c *uringTCPConn) SetDeadline(t time.Time) error {
	// todo
	return nil
}

func (c *uringTCPConn) SetReadDeadline(t time.Time) error {
	// todo
	return nil
}

func (c *uringTCPConn) SetWriteDeadline(t time.Time) error {
	// todo
	return nil
}

//go:linkname anyToSockaddr syscall.anyToSockaddr
func anyToSockaddr(rsa *syscall.RawSockaddrAny) (syscall.Sockaddr, error)

//go:linkname sockaddrToTCP net.sockaddrToTCP
func sockaddrToTCP(sa syscall.Sockaddr) net.Addr
