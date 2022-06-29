package gouringasm

import (
	"os"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	h, err := New(256, 0)
	assert.NoError(t, err)
	assert.NotNil(t, h)
	assert.NotNil(t, h.h)

	assert.NotNil(t, h.h.sq.khead)
	assert.NotNil(t, h.h.sq.sqes)

	assert.NotNil(t, h.h.cq.khead)
	assert.NotNil(t, h.h.cq.cqes)
}

func TestSubmitAndGetCQE(t *testing.T) {
	const n = 10
	h, err := New(256, 0)
	assert.NoError(t, err)

	b := []byte("submit test 1011!\n")
	ud := uint64(0xcafebabe)
	for i := 0; i < n; i++ {
		sqe := h.GetSQE()
		IoUringPrepRW(23, sqe, int(os.Stdout.Fd()),
			unsafe.Pointer(&b[0]), len(b), 0)
		sqe.user_data = ud
	}

	submitted, err := h.Submit()
	assert.NoError(t, err)
	assert.Equal(t, int32(n), submitted)

	var cqe *IoUringCQE
	err = h.WaitCQE(&cqe)
	assert.NoError(t, err)
	assert.NotNil(t, cqe)

	assert.Equal(t, ud, cqe.user_data)
}
