package gouringasm

import (
	"sync"
	"testing"
)

func TestSched(t *testing.T) {
	const task = 10_000
	h := testNewIoUring(t, 256, 0)
	defer h.Close()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		var sqe *IoUringSqe
		var err error
		i := 0
		for i < task {
			sqe = h.GetSqe()
			if sqe == nil {
				// keep trying..
				continue
			}
			PrepRW(IORING_OP_NOP, sqe, 0, nil, 0, 0)
			IoUringSetUserdata(sqe, uint64(i))
			_, err = h.Submit()
			if err != nil {
				panic(err)
			}
			i++
		}
		println(i)
	}()
	go func() {
		defer wg.Done()
		i := 0
		var cqe *IoUringCqe
		var err error
		for i < task {
			err = h.WaitCqe(&cqe)
			if err != nil {
				panic(err)
			}
			if cqe == nil {
				panic("cqe returned is nil")
			}
			i++
		}
		println(i)
	}()

	wg.Wait()
}
