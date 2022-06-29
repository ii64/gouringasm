package gouringasm

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSched(t *testing.T) {
	const task = 10_000
	h, err := New(256, 0)
	assert.NoError(t, err)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		var sqe *IoUringSQE
		var err error
		i := 0
		for i < task {
			sqe = h.GetSQE()
			if sqe == nil {
				// keep trying..
				continue
			}
			IoUringPrepRW(IORING_OP_NOP, sqe, 0, nil, 0, 0)
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
		var cqe *IoUringCQE
		var err error
		for i < task {
			err = h.WaitCQE(&cqe)
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
