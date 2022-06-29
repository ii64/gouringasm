package gouringasm

import "syscall"

type Uring struct {
	h *IoUring
}

func New(entries, flags uint32) (*Uring, error) {
	ring := &Uring{}
	ring.h = new(IoUring)
	// ring.h = allocIoUringHeap()
	ret := IoUringQueueInit(entries, ring.h, flags)
	if ret < 0 {
		return nil, syscall.Errno(-ret)
	}
	return ring, nil
}

func NewWithParams(entries uint32, params *IoUringParams) (*Uring, error) {
	ring := &Uring{}
	ring.h = new(IoUring)
	ret := IoUringQueueInitParams(entries, ring.h, params)
	if ret < 0 {
		return nil, syscall.Errno(-ret)
	}
	return ring, nil
}

func (ring *Uring) Close() {
	if ring.h != nil {
		IoUringQueueExit(ring.h)
	}
}

func (ring *Uring) Ring() *IoUring {
	return ring.h
}

func (ring *Uring) GetSQE() *IoUringSQE {
	return IoUringGetSQE(ring.h)
}

func (ring *Uring) Submit() (int32, error) {
	ret := IoUringSubmit(ring.h)
	if ret < 0 {
		return 0, syscall.Errno(-ret)
	}
	return ret, nil
}

func (ring *Uring) SubmitAndWait(waitNr uint32) (int32, error) {
	ret := IoUringSubmitAndWait(ring.h, waitNr)
	if ret < 0 {
		return 0, syscall.Errno(-ret)
	}
	return ret, nil
}

func (ring *Uring) WaitCQE(cqePtr **IoUringCQE) error {
	ret := IoUringWaitCQE(ring.h, cqePtr)
	if ret < 0 {
		return syscall.Errno(-ret)
	}
	return nil
}
