package gouringasm

import "unsafe"

const (
	SizeofUnsigned   = unsafe.Sizeof(uint32(0))
	SizeofUint32     = unsafe.Sizeof(uint32(0))
	SizeofIoUringSqe = unsafe.Sizeof(IoUringSqe{})
	SizeofIoUringCqe = unsafe.Sizeof(IoUringCqe{})
)

//go:generate stringer -type=IoUringOp -trimprefix=IORING_OP_
type IoUringOp uint8

//go:generate stringer -typeIoUringSetupFlag -trimprefix=IORING_SETUP_
type IoUringSetupFlag uint32

type IoUring struct {
	Sq          IoUringSq
	Cq          IoUringCq
	Flags       uint32
	RingFd      uint32
	Features    uint32
	EnterRingFd uint32
	IntFlags    uint8
	pad         [3]uint8
	pad2        uint32
}

type IoUringCq struct {
	Head        unsafe.Pointer
	Tail        unsafe.Pointer
	RingMask    unsafe.Pointer
	RingEntries unsafe.Pointer
	Flags       unsafe.Pointer
	Overflow    unsafe.Pointer
	Cqes        ioUringCqeArray
	RingSz      uint32
	RingPtr     unsafe.Pointer
	pad         [4]__uint
}

type IoUringSq struct {
	Head        unsafe.Pointer
	Tail        unsafe.Pointer
	RingMask    unsafe.Pointer
	RingEntries unsafe.Pointer
	Flags       unsafe.Pointer
	Dropped     unsafe.Pointer
	Array       uint32Array
	Sqes        ioUringSqeArray
	SqeHead     uint32
	SqeTail     uint32
	RingSz      uint32
	RingPtr     unsafe.Pointer
	pad         [4]__uint
}
