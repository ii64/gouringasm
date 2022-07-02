package main

import (
	"fmt"
	"os"
	"sync"
	"time"
	"unsafe"

	"github.com/ii64/gouringasm"
	"github.com/ii64/gouringasm/evloop"
)

func main() {

	h, err := gouringasm.NewWithParams(256, &gouringasm.IoUringParams{
		Flags:        gouringasm.IORING_SETUP_SQPOLL,
		SqThreadCPU:  4,
		SqThreadIdle: 10_000,
	})
	if err != nil {
		panic(err)
	}
	ev := evloop.New(h)
	defer func() {
		ev.Close()
		h.Close()
	}()

	go ev.Serve()

	b := []byte("hello world\n")

	wr := func(sqe *gouringasm.IoUringSQE, ud *evloop.EventloopUserdata) {
		gouringasm.IoUringPrepRW(gouringasm.IORING_OP_WRITE, sqe, int(os.Stdout.Fd()),
			unsafe.Pointer(&b[0]), len(b), 0)
	}

	var wg sync.WaitGroup
	var wgSetup sync.WaitGroup
	const N = 10
	wg.Add(N)
	wgSetup.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()

			sqe, ud := ev.GetSQE()
			wr(sqe, ud)
			wgSetup.Done()

			startWait := time.Now()
			ud.Wait()
			elapsed := time.Since(startWait)
			fmt.Println("resx", ud.Res, elapsed)
		}()
	}

	wgSetup.Wait()

	var submitted int32
	submitted, err = ev.Submit()
	if err != nil {
		panic(err)
	}
	fmt.Println("submitted", submitted)

	wg.Wait()
}
