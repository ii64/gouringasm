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
func IoUringPrepRW(
	op uintptr, sqe unsafe.Pointer,
	fd uintptr,
	addr unsafe.Pointer,
	len uintptr, offset uintptr)

//go:nosplit
//go:noescape
func IoUringGetSQE(ring unsafe.Pointer) (sqe unsafe.Pointer)

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
