package native

import "unsafe"

// !! full store to register (64-bit) !!

//go:nosplit
//go:noescape
func IoUringQueueInit(
	entries uintptr,
	ring unsafe.Pointer,
	flags uintptr) (ret uintptr)

//go:nosplit
//go:noescape
func IoUringQueueInitParams(
	entries uintptr,
	ring unsafe.Pointer,
	params unsafe.Pointer) (ret uintptr)

//go:nosplit
//go:noescape
func IoUringQueueExit(
	ring unsafe.Pointer) (ret uintptr)

//go:nosplit
//go:noescape
func IoUringSubmit(
	ring unsafe.Pointer) (ret uintptr)

//go:nosplit
//go:noescape
func IoUringSubmitAndWait(
	ring unsafe.Pointer,
	waitNr uintptr) (ret uintptr)

// !! ---- sqe ----

//go:nosplit
//go:noescape
func IoUringGetSQE(ring unsafe.Pointer) (sqe unsafe.Pointer)

// ** Prep

//go:nosplit
//go:noescape
func IoUringPrepRW(
	op uintptr, sqe unsafe.Pointer,
	fd uintptr,
	addr unsafe.Pointer,
	len uintptr, offset uintptr)

// //go:nosplit
// //go:noescape
// func IoUringPrepSplice(
// 	sqe unsafe.Pointer)

//go:nosplit
//go:noescape
func IoUringPrepTee(
	sqe unsafe.Pointer,
	fdIn uintptr, fdOut uintptr,
	nbytes uintptr, spliceFlags uintptr)

//go:nosplit
//go:noescape
func IoUringPrepReadv(
	sqe unsafe.Pointer,
	fd uintptr, iovecs unsafe.Pointer,
	nrVecs uintptr, offset uintptr)

//go:nosplit
//go:noescape
func IoUringPrepReadv2(
	sqe unsafe.Pointer,
	fd uintptr, iovecs unsafe.Pointer,
	nrVecs uintptr, offset uintptr,
	flags uintptr)

//go:nosplit
//go:noescape
func IoUringPrepWritev(
	sqe unsafe.Pointer,
	fd uintptr, iovecs unsafe.Pointer,
	nrVecs uintptr, offset uintptr)

//go:nosplit
//go:noescape
func IoUringPrepWritev2(
	sqe unsafe.Pointer,
	fd uintptr, iovecs unsafe.Pointer,
	nrVecs uintptr, offset uintptr,
	flags uintptr)

//go:nosplit
//go:noescape
func IoUringPrepRecvmsg(
	sqe unsafe.Pointer,
	fd uintptr, msg unsafe.Pointer,
	flags uintptr)

//go:nosplit
//go:noescape
func IoUringPrepSendmsg(
	sqe unsafe.Pointer,
	fd uintptr, msg unsafe.Pointer,
	flags uintptr)

//go:nosplit
//go:noescape
func IoUringPrepFsync(
	sqe unsafe.Pointer,
	fd uintptr, flags uintptr)

//go:nosplit
//go:noescape
func IoUringPrepNop(
	sqe unsafe.Pointer)

//go:nosplit
//go:noescape
func IoUringPrepAccept(
	sqe unsafe.Pointer,
	fd uintptr, addr unsafe.Pointer,
	addrlen unsafe.Pointer,
	flags uintptr)

//go:nosplit
//go:noescape
func IoUringPrepMultishotAccept(
	sqe unsafe.Pointer,
	fd uintptr, addr unsafe.Pointer,
	addrlen unsafe.Pointer,
	flags uintptr)

//go:nosplit
//go:noescape
func IoUringPrepConnect(
	sqe unsafe.Pointer,
	addr unsafe.Pointer,
	addrlen uintptr)

//go:nosplit
//go:noescape
func IoUringPrepClose(
	sqe unsafe.Pointer, fd uintptr)

//go:nosplit
//go:noescape
func IoUringPrepRead(
	sqe unsafe.Pointer,
	buf unsafe.Pointer,
	nbytes uintptr,
	offset uintptr)

//go:nosplit
//go:noescape
func IoUringPrepWrite(
	sqe unsafe.Pointer,
	buf unsafe.Pointer,
	nbytes uintptr,
	offset uintptr)

//go:nosplit
//go:noescape
func IoUringPrepSend(
	sqe unsafe.Pointer,
	sockfd uintptr,
	buf unsafe.Pointer,
	flags uintptr)

//go:nosplit
//go:noescape
func IoUringPrepRecv(
	sqe unsafe.Pointer,
	sockfd uintptr,
	buf unsafe.Pointer,
	flags uintptr)

//go:nosplit
//go:noescape
func IoUringPrepEpollCtl(
	sqe unsafe.Pointer,
	epfd uintptr,
	fd uintptr,
	op uintptr,
	ev unsafe.Pointer)

// !! ---- cqe ----

//go:nosplit
//go:noescape
func IoUringWaitCQE(
	ring unsafe.Pointer,
	cqePtr unsafe.Pointer) (ret uintptr)

//go:nosplit
//go:noescape
func IoUringCQESeen(
	ring unsafe.Pointer,
	cqe unsafe.Pointer)
