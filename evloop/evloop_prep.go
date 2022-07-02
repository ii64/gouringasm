package evloop

import (
	"syscall"
	"unsafe"

	"github.com/ii64/gouringasm"
)

func (ev *Eventloop) QueueAccept(fd int) (*syscall.RawSockaddrAny, *int32, *EventloopUserdata) {
	sqe, ud := ev.GetSQE()
	rsa := new(syscall.RawSockaddrAny)
	var rsaSz int32 = syscall.SizeofSockaddrAny
	gouringasm.IoUringPrepRW(gouringasm.IORING_OP_ACCEPT, sqe, fd,
		unsafe.Pointer(rsa), 0, uint64(uintptr(unsafe.Pointer(&rsaSz))))
	return rsa, &rsaSz, ud
}

func (ev *Eventloop) QueueWrite(fd int, b []byte) *EventloopUserdata {
	sqe, ud := ev.GetSQE()
	gouringasm.IoUringPrepRW(gouringasm.IORING_OP_WRITE, sqe, fd,
		unsafe.Pointer(&b[0]), len(b), 0)
	return ud
}

func (ev *Eventloop) QueueRead(fd int, b []byte) *EventloopUserdata {
	sqe, ud := ev.GetSQE()
	gouringasm.IoUringPrepRW(gouringasm.IORING_OP_READ, sqe, fd,
		unsafe.Pointer(&b[0]), len(b), 0)
	return ud
}

func (ev *Eventloop) QueueClose(fd int) *EventloopUserdata {
	sqe, ud := ev.GetSQE()
	gouringasm.IoUringPrepRW(gouringasm.IORING_OP_CLOSE, sqe, fd,
		nil, 0, 0)
	return ud
}
