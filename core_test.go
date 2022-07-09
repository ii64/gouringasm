package gouringasm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type genericTestingT interface {
	assert.TestingT
	require.TestingT
}

func testNewIoUring(t genericTestingT, entries uint32, flags uint32) *IoUring {
	h, err := New(entries, flags)
	require.NoError(t, err)
	require.NotNil(t, h)
	return h
}

func testNewIoUringWithParams(t genericTestingT, entries uint32, p *IoUringParams) *IoUring {
	h, err := NewWithParams(entries, p)
	require.NoError(t, err)
	require.NotNil(t, h)
	return h
}

func TestNew(t *testing.T) {
	h := testNewIoUring(t, 256, 0)
	defer h.Close()

	assert.NotNil(t, h.Sq.Head)
	assert.NotNil(t, h.Sq.Sqes)

	assert.NotNil(t, h.Cq.Head)
	assert.NotNil(t, h.Cq.Cqes)
}
func TestNewWithParams(t *testing.T) {
	h := testNewIoUringWithParams(t, 256, &IoUringParams{
		Flags:        IORING_SETUP_SQPOLL,
		SqThreadCpu:  4,
		SqThreadIdle: 10_000,
	})
	assert.NotNil(t, h)

	assert.NotNil(t, h.Sq.Head)
	assert.NotNil(t, h.Sq.Sqes)

	assert.NotNil(t, h.Cq.Head)
	assert.NotNil(t, h.Cq.Cqes)
}
