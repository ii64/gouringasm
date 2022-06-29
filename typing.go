package gouringasm

import (
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

func IoUringPrepRW(op IoUringOpcode, sqe *IoUringSQE, fd int, addr unsafe.Pointer, len int, offset uint64) {
	// runtime_entersyscall()
	native.IoUringPrepRW(uintptr(op), unsafe.Pointer(sqe), uintptr(fd), addr, uintptr(len), uintptr(offset))
	// runtime_exitsyscall()
}

func IoUringGetSQE(ring *IoUring) (sqe *IoUringSQE) {
	// runtime_entersyscall()
	sqe = (*IoUringSQE)(native.IoUringGetSQE(unsafe.Pointer(ring)))
	// runtime_exitsyscall()
	return
}

func IoUringSetUserdata(sqe *IoUringSQE, userdata uint64) {
	sqe.user_data = userdata
}

// !! --- cqe ---

func IoUringWaitCQE(ring *IoUring, cqePtr **IoUringCQE) (r int32) {
	// runtime_entersyscall()
	r = int32(native.IoUringWaitCQE(unsafe.Pointer(ring), unsafe.Pointer(cqePtr)))
	// runtime_exitsyscall()
	return
}
