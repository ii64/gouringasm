package gouringasm

import (
	"sync/atomic"
	"unsafe"

	"github.com/ii64/gouringasm/internal/native"
)

func IoUringQueueInit(entries uint32, ring *IoUring, flags uint32) (r int32) {
	// runtime_entersyscall()
	r = int32(native.IoUringQueueInit(
		uintptr(entries),
		unsafe.Pointer(ring),
		uintptr(flags)))
	// runtime_exitsyscall()
	return
}

func IoUringQueueInitParams(entries uint32, ring *IoUring, params *IoUringParams) (r int32) {
	// runtime_entersyscall()
	r = int32(native.IoUringQueueInitParams(
		uintptr(entries),
		unsafe.Pointer(ring),
		unsafe.Pointer(params)))
	// runtime_exitsyscall()
	return
}

func IoUringQueueExit(ring *IoUring) {
	// runtime_entersyscall()
	native.IoUringQueueExit(unsafe.Pointer(ring))
	// runtime_exitsyscall()
}

func IoUringSubmit(ring *IoUring) (r int32) {
	// runtime_entersyscall()
	r = int32(native.IoUringSubmit(unsafe.Pointer(ring)))
	// runtime_exitsyscall()
	return
}

func IoUringSubmitAndWait(ring *IoUring, waitNr uint32) (r int32) {
	// runtime_entersyscall()
	r = int32(native.IoUringSubmitAndWait(unsafe.Pointer(ring), uintptr(waitNr)))
	// runtime_exitsyscall()
	return
}

// !! --- sqe ---

func IoUringGetSqe(ring *IoUring) (sqe *IoUringSqe) {
	// runtime_entersyscall()
	sqe = (*IoUringSqe)(native.IoUringGetSQE(unsafe.Pointer(ring)))
	// runtime_exitsyscall()
	return
}

func IoUringSetUserdata(sqe *IoUringSqe, userdata uint64) {
	sqe.UserData.SetUint64(userdata)
}

// !! --- cqe ---

func IoUringWaitCqe(ring *IoUring, cqePtr **IoUringCqe) (r int32) {
	// runtime_entersyscall()
	r = int32(native.IoUringWaitCQE(unsafe.Pointer(ring), unsafe.Pointer(cqePtr)))
	// runtime_exitsyscall()
	return
}

func IoUringCqeSeen(ring *IoUring, cqe *IoUringCqe) {
	// runtime_entersyscall()
	atomic.AddUint32((*uint32)(ring.Cq.Head), 1)
	// native.IoUringCQESeen(unsafe.Pointer(ring), unsafe.Pointer(cqe))
	// runtime_exitsyscall()
}
