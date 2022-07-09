package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	_ "net/http/pprof"
	_ "unsafe"

	"github.com/ii64/gouringasm"
)

var (
	h  *gouringasm.IoUring
	fd int

	connPoolx = sync.Pool{
		New: func() any { return newConn() },
	}
	userdataPoolx = sync.Pool{
		New: func() any { return newUserdata() },
	}
)

type eventloop struct {
	OnConnect func(c *Conn)
	OnRead    func(c *Conn)
	OnWrite   func(c *Conn)
	OnClose   func(c *Conn)
}

var evloop = eventloop{
	OnConnect: func(c *Conn) {
		// fmt.Println("got connection")
		c.Read(c.buf)
	},
	OnRead: func(c *Conn) {
		// fmt.Printf("readed %d\n", c.buflen)

		hdr := []byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nServer: net-tcp\r\nContent-Length: %d\r\nConnection: close\r\n\r\n", c.buflen))
		c.buf = append(hdr, c.buf[:c.buflen]...)
		c.buflen = len(hdr) + c.buflen

		c.Write(c.buf[:c.buflen])
	},
	OnWrite: func(c *Conn) {
		c.Close()
	},
	OnClose: func(c *Conn) {
		// connPool.Put(c)
	},
}

// **

type Conn struct {
	fd    int
	rsa   *syscall.RawSockaddrAny
	rsaSz *uintptr

	buf    []byte
	buflen int
}

func newConn() *Conn {
	conn := new(Conn)
	conn.setup()
	return conn
}

func getConn() *Conn {
	// return connPool.Get().(*Conn).setup()
	return newConn()
}

func (c *Conn) setup() *Conn {
	c.rsa = &syscall.RawSockaddrAny{}
	c.rsaSz = new(uintptr)
	*c.rsaSz = syscall.SizeofSockaddrAny
	c.buf = make([]byte, 1024, 1024)
	return c
}

func (c *Conn) Read(b []byte) {
	queueRead(c, b)
}

func (c *Conn) Write(b []byte) {
	queueWrite(c, b)
}

func (c *Conn) Close() {
	queueClose(c)
}

// **

type userdata struct {
	Op gouringasm.IoUringOp
	*Conn
}

func (u *userdata) zero() *userdata {
	u.Op = 0
	u.Conn = nil
	return u
}

func (u *userdata) uptr() unsafe.Pointer {
	return unsafe.Pointer(u)
}

// **

func newUserdata() *userdata {
	ud := new(userdata)
	return ud
}

func getUserdata(c *Conn) *userdata {
	// ud := userdataPool.Get().(*userdata).zero()
	ud := newUserdata()
	ud.Conn = c
	if c == nil {
		ud.Conn = getConn()
	}
	return ud
}

func putUserdata(u *userdata) {
	// userdataPool.Put(u)
}

// **

func getSQE() (sqe *gouringasm.IoUringSqe) {
	for {
		sqe = h.GetSqe()
		if sqe != nil {
			return
		}
		runtime.Gosched()
	}
}

func queueAccept() (*userdata, error) {
	sqe := getSQE()
	sqe.Flags |= gouringasm.IOSQE_IO_LINK
	ud := getUserdata(nil)
	// gouringasm.PrepMultishotAccept(sqe, fd, rsa, rsaSz, 0)
	gouringasm.PrepAccept(sqe, fd, ud.rsa, ud.rsaSz, 0)
	ud.Op = sqe.Opcode
	sqe.UserData.SetUnsafe(ud.uptr())

	_, err := h.Submit()
	return ud, err
}

func queueRead(c *Conn, b []byte) (*userdata, error) {
	sqe := getSQE()
	sqe.Flags |= gouringasm.IOSQE_IO_LINK
	ud := getUserdata(c)
	gouringasm.PrepRead(sqe, c.fd, &b[0], len(b), 0)
	ud.Op = sqe.Opcode
	sqe.UserData.SetUnsafe(ud.uptr())
	_, err := h.Submit()
	return ud, err
}

func queueWrite(c *Conn, b []byte) (*userdata, error) {
	sqe := getSQE()
	sqe.Flags |= gouringasm.IOSQE_IO_LINK
	ud := getUserdata(c)
	gouringasm.PrepWrite(sqe, c.fd, &b[0], len(b), 0)
	ud.Op = sqe.Opcode
	sqe.UserData.SetUnsafe(ud.uptr())
	_, err := h.Submit()
	return ud, err
}

func queueClose(c *Conn) (*userdata, error) {
	sqe := getSQE()
	sqe.Flags |= gouringasm.IOSQE_IO_LINK
	ud := getUserdata(c)
	gouringasm.PrepClose(sqe, c.fd)
	ud.Op = sqe.Opcode
	sqe.UserData.SetUnsafe(ud.uptr())
	_, err := h.Submit()
	return ud, err
}

// **

func main() {
	fmt.Printf("pid: %d\n", os.Getpid())

	var err error
	// h, err = gouringasm.New(256, 0)
	h, err = gouringasm.NewWithParams(256, &gouringasm.IoUringParams{
		Flags:        gouringasm.IORING_SETUP_SQPOLL,
		SqThreadCpu:  10,
		SqThreadIdle: 10_000,
	})
	if err != nil {
		panic(err)
	}
	defer h.Close()

	fd, err = tcpLister(":9001")
	if err != nil {
		panic(err)
	}

	fmt.Printf("fd: %+#v\n", fd)

	queueAccept()

	go func() {
		http.HandleFunc("/queue_accept", func(w http.ResponseWriter, r *http.Request) {
			queueAccept()
			w.WriteHeader(200)
		})
		http.ListenAndServe(":9002", nil)
	}()

	go func() {
		for {
			sqMask := *(*uint32)(h.Sq.RingMask)
			sqFlag := atomic.LoadUint32((*uint32)(h.Sq.Flags))
			sqeKHead := atomic.LoadUint32((*uint32)(h.Sq.Head)) & sqMask
			sqeKTail := atomic.LoadUint32((*uint32)(h.Sq.Tail)) & sqMask

			cqMask := *(*uint32)(h.Cq.RingMask)
			cqFlag := atomic.LoadUint32((*uint32)(h.Cq.Flags))
			cqeKHead := atomic.LoadUint32((*uint32)(h.Cq.Head)) & cqMask
			cqeKTail := atomic.LoadUint32((*uint32)(h.Cq.Tail)) & cqMask

			fmt.Printf("sqe head:%d tail: %d flag: %d\n", sqeKHead, sqeKTail, sqFlag)
			fmt.Printf("cqe head:%d tail: %d flag: %d\n", cqeKHead, cqeKTail, cqFlag)
			<-time.After(time.Second * 1)
		}
	}()

	var cqe *gouringasm.IoUringCqe
	var ud *userdata
	for {
		err = h.WaitCqe(&cqe)
		if err == syscall.EINTR {
			// fmt.Println("EINTR")
			runtime.Gosched()
			goto exit3
		}
		if err != nil {
			panic(err)
		}
		if cqe == nil {
			panic("cqe is nil")
		}
		if cqe.Res < 0 {
			errno := -cqe.Res
			err = syscall.Errno(errno)
			ud = (*userdata)(cqe.UserData.GetUnsafe())
			fmt.Printf("cqe: %+#v\n", cqe)
			if err == syscall.EBADF {
				queueAccept()
				goto exit2
			} else if err == syscall.EAGAIN {
				goto exit2
			} else if err == syscall.EPIPE {
				goto exit2
			} else if err == syscall.ECONNRESET {
				goto exit2
			} else if err != nil {
				panic(err)
			}

		}
		if cqe.UserData.GetUint64() < 1 {
			continue
		}

		// ** handle

		ud = (*userdata)(cqe.UserData.GetUnsafe())
		if ud.Conn == nil {
			goto exit2
		}
		switch ud.Op {
		case gouringasm.IORING_OP_ACCEPT:
			ud.Conn.fd = int(cqe.Res)
			evloop.OnConnect(ud.Conn)
			queueAccept()
		case gouringasm.IORING_OP_READ:
			ud.Conn.buflen = int(cqe.Res)
			evloop.OnRead(ud.Conn)
		case gouringasm.IORING_OP_WRITE:
			ud.Conn.buflen = int(cqe.Res)
			evloop.OnWrite(ud.Conn)
		case gouringasm.IORING_OP_CLOSE:
			evloop.OnClose(ud.Conn)
		default:
			fmt.Printf("cqe: %+#v ud: %v\n", cqe, ud)
		}

		goto exit1
	exit1:
		putUserdata(ud)
	exit2:
		h.SeenCqe(cqe)
	exit3:
		// runtime.GC()
	}

}

func stdOpenSocket(addr string) (fd int, err error) {
	var ln net.Listener
	ln, err = net.Listen("tcp", addr)
	if err != nil {
		return
	}
	tcpLn, ok := ln.(*net.TCPListener)
	if !ok {
		err = fmt.Errorf("not a TCPListener")
		return
	}
	var f *os.File
	f, err = tcpLn.File()
	if err != nil {
		return
	}
	fd = int(f.Fd())
	return
}

func openSocket(addr string) (fd int, err error) {
	var host string
	var port string
	host, port, err = net.SplitHostPort(addr)
	if err != nil {
		return
	}

	var portN uint64
	portN, err = strconv.ParseUint(port, 10, 16)
	if err != nil {
		return
	}

	var portE [2]byte
	binary.LittleEndian.PutUint16(portE[:], uint16(portN))
	portN = uint64(binary.BigEndian.Uint16(portE[:]))

	fd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM|syscall.SOCK_CLOEXEC, 0)
	if err != nil {
		return
	}

	const SO_REUSEPORT = 15

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, SO_REUSEPORT, 1)
	if err != nil {
		return
	}
	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	if err != nil {
		return
	}

	rs4 := &syscall.RawSockaddrInet4{
		Family: syscall.AF_INET,
		Port:   uint16(portN),
		Addr:   [4]byte{},
		Zero:   [8]uint8{},
	}
	copy(rs4.Addr[:], net.ParseIP(host))
	rsa := &syscall.RawSockaddrAny{
		Addr: *(*syscall.RawSockaddr)(unsafe.Pointer(rs4)),
		Pad:  [96]int8{},
	}

	var sa syscall.Sockaddr
	sa, err = anyToSockaddr(rsa)
	if err != nil {
		return
	}

	laddr := sockaddrToTCP(sa)
	fmt.Printf("laddr: %s\n", laddr)

	err = syscall.Bind(fd, sa)
	if err != nil {
		return
	}

	err = syscall.Listen(fd, 128)
	if err != nil {
		return
	}

	return
}

func tcpLister(addr string) (fd int, err error) {

	if false {
		fd, err = stdOpenSocket(addr)
	} else {
		fd, err = openSocket(addr)
	}

	syscall.CloseOnExec(fd)
	err = syscall.SetNonblock(fd, true)

	var val int
	val, err = syscall.GetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_TYPE)
	if val != syscall.SOCK_STREAM {
		panic("not SOCK_STREAM")
	}
	return
}

//go:linkname anyToSockaddr syscall.anyToSockaddr
func anyToSockaddr(rsa *syscall.RawSockaddrAny) (syscall.Sockaddr, error)

//go:linkname sockaddrToTCP net.sockaddrToTCP
func sockaddrToTCP(sa syscall.Sockaddr) net.Addr

//go:linkname fcntl syscall.fcntl
func fcntl(fd int, cmd int, arg int) (val int, err error)
