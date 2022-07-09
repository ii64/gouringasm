package gouringasm

import (
	"syscall"
	"unsafe"
)

func PrepRW(op IoUringOp, sqe *IoUringSqe, fd int,
	addr unsafe.Pointer, len int, offset uint64) {
	// Go asm epilogiue...
	//	    	.       70ms   4b05c0:     CALL github.com/ii64/gouringasm/internal/native.IoUringPrepRW.abi0(SB) prep.go:12
	//      	.          .   4b05c5:     XORPS X15, X15                                                       prep.go:12
	// ->*	790ms      790ms   4b05c9:     MOVQ FS:0xfffffff8, R14
	// native.IoUringPrepRW(uintptr(op), unsafe.Pointer(sqe), uintptr(fd), addr, uintptr(len), uintptr(offset))

	// inlined
	sqe.Opcode = op
	sqe.Flags = 0
	sqe.IoPrio = 0
	sqe.Fd = int32(fd)
	sqe.IoUringSqe_Union1 = IoUringSqe_Union1(offset)                    // union1
	sqe.IoUringSqe_Union2 = *(*IoUringSqe_Union2)(unsafe.Pointer(&addr)) // union2
	sqe.Len = uint32(len)
	sqe.IoUringSqe_Union3 = 0 // sqe.SetOpFlags(0) // union3
	sqe.IoUringSqe_Union4 = 0 // sqe.SetBufIndex(0) // union4
	sqe.Personality = 0
	sqe.IoUringSqe_Union5 = 0 // sqe.SetFileIndex(0) // union5
	sqe.Addr3 = 0
	sqe.__pad2[0] = 0
}

func PrepNop(sqe *IoUringSqe) {
	// native.IoUringPrepNop(unsafe.Pointer(sqe))
	PrepRW(IORING_OP_NOP, sqe, -1, nil, 0, 0)
}

func PrepTimeout(sqe *IoUringSqe, ts *syscall.Timespec, count uint32, flags uint32) {
	PrepRW(IORING_OP_TIMEOUT, sqe, -1, unsafe.Pointer(ts), 1, uint64(count))
	sqe.SetTimeoutFlags(flags)
}

func PrepTimeoutRemove(sqe *IoUringSqe, userDaata uint64, flags uint32) {
	PrepRW(IORING_OP_TIMEOUT_REMOVE, sqe, -1, nil, 0, 0)
	sqe.SetAddr_Value(userDaata)
	sqe.SetTimeoutFlags(flags)
}

func PrepTimeoutUpdate(sqe *IoUringSqe, ts *syscall.Timespec, userData uint64, flags uint32) {
	PrepRW(IORING_OP_TIMEOUT_REMOVE, sqe, -1, nil, 0, 0)
	sqe.SetAddr_Value(userData)
	sqe.SetTimeoutFlags(flags | IORING_TIMEOUT_UPDATE)
}

// ** "Syscall" OP

func PrepRead(sqe *IoUringSqe, fd int, buf *byte, nb int, offset uint64) {
	PrepRW(IORING_OP_READ, sqe, fd, unsafe.Pointer(buf), nb, offset)
}
func PrepReadv(sqe *IoUringSqe, fd int,
	iov *syscall.Iovec, nrVecs int,
	offset uint64) {
	PrepRW(IORING_OP_READV, sqe, fd, unsafe.Pointer(iov), nrVecs, offset)
}
func PrepReadv2(sqe *IoUringSqe, fd int,
	iov *syscall.Iovec, nrVecs int,
	offset uint64, flags uint32) {
	PrepReadv(sqe, fd, iov, nrVecs, offset)
	sqe.SetRwFlags(flags)
}

func PrepWrite(sqe *IoUringSqe, fd int, buf *byte, nb int, offset uint64) {
	PrepRW(IORING_OP_WRITE, sqe, fd, unsafe.Pointer(buf), nb, offset)
}
func PrepWritev(sqe *IoUringSqe, fd int,
	iov *syscall.Iovec, nrVecs int,
	offset uint64) {
	PrepRW(IORING_OP_WRITEV, sqe, fd, unsafe.Pointer(iov), nrVecs, offset)
}
func PrepWritev2(sqe *IoUringSqe, fd int,
	iov *syscall.Iovec, nrVecs int,
	offset uint64, flags uint32) {
	PrepWritev(sqe, fd, iov, nrVecs, offset)
	sqe.SetRwFlags(flags)
}

func PrepAccept(sqe *IoUringSqe, fd int, rsa *syscall.RawSockaddrAny, rsaSz *uintptr, flags uint32) {
	// *rsaSz = syscall.SizeofSockaddrAny // leave this out to caller?
	PrepRW(IORING_OP_ACCEPT, sqe, fd, unsafe.Pointer(rsa), 0, uint64(uintptr(unsafe.Pointer(rsaSz))))
	sqe.SetAcceptFlags(uint32(flags))
}

func PrepClose(sqe *IoUringSqe, fd int) {
	PrepRW(IORING_OP_CLOSE, sqe, fd, nil, 0, 0)
}

func PrepRecvmsg(sqe *IoUringSqe, fd int, msg *syscall.Msghdr, flags uint32) {
	PrepRW(IORING_OP_RECVMSG, sqe, fd, unsafe.Pointer(msg), 1, 0)
	sqe.SetMsgFlags(uint32(flags))
}

func PrepSendmsg(sqe *IoUringSqe, fd int, msg *syscall.Msghdr, flags uint32) {
	PrepRW(IORING_OP_SENDMSG, sqe, fd, unsafe.Pointer(msg), 1, 0)
	sqe.SetMsgFlags(uint32(flags))
}

func PrepEpollCtl(sqe *IoUringSqe, epfd int, op int, fd int, event *syscall.EpollEvent) {
	PrepRW(IORING_OP_EPOLL_CTL, sqe, epfd, unsafe.Pointer(event), op, uint64(fd))
}

// ** Multishot

func PrepMultishotAccept(sqe *IoUringSqe, fd int, rsa *syscall.RawSockaddrAny, rsaSz *uintptr, flags uint32) {
	PrepAccept(sqe, fd, rsa, rsaSz, flags)
	sqe.IoPrio |= IORING_ACCEPT_MULTISHOT
}
