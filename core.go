package gouringasm

import "syscall"

// type Uring struct {
// 	*IoUring
// }

func New(entries, flags uint32) (*IoUring, error) {
	ring := new(IoUring)
	// ring.h = allocIoUringHeap()
	ret := IoUringQueueInit(entries, ring, flags)
	if ret < 0 {
		return nil, syscall.Errno(-ret)
	}
	return ring, nil
}

func NewWithParams(entries uint32, params *IoUringParams) (*IoUring, error) {
	ring := &IoUring{}
	ret := IoUringQueueInitParams(entries, ring, params)
	if ret < 0 {
		return nil, syscall.Errno(-ret)
	}
	return ring, nil
}

func (ring *IoUring) Close() {
	IoUringQueueExit(ring)
}

func (ring *IoUring) GetSqe() *IoUringSqe {
	return IoUringGetSqe(ring)
}

func (ring *IoUring) WaitCqe(cqePtr **IoUringCqe) error {
	ret := IoUringWaitCqe(ring, cqePtr)
	if ret < 0 {
		return syscall.Errno(-ret)
	}
	return nil
}

func (ring *IoUring) SeenCqe(cqe *IoUringCqe) {
	IoUringCqeSeen(ring, cqe)
}

func (ring *IoUring) Submit() (int, error) {
	ret := IoUringSubmit(ring)
	if ret < 0 {
		return 0, syscall.Errno(-ret)
	}
	return int(ret), nil
}

func (ring *IoUring) SubmitAndWait(waitNr uint32) (int, error) {
	ret := IoUringSubmitAndWait(ring, waitNr)
	if ret < 0 {
		return 0, syscall.Errno(-ret)
	}
	return int(ret), nil
}
