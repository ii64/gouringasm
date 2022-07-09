package gouringasm

import "unsafe"

type uint32Array = unsafe.Pointer // *uint32

func uint32Array_Index(u uint32Array, i uintptr) *uint32 {
	return (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + SizeofUint32*i))
}

type ioUringSqeArray = unsafe.Pointer // *IoUringSqe

func ioUringSqeArray_Index(u ioUringSqeArray, i uintptr) *IoUringSqe {
	return (*IoUringSqe)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + SizeofIoUringSqe*i))
}

//

type ioUringCqeArray = unsafe.Pointer // *IoUringCqe

func ioUringCqeArray_Index(u ioUringCqeArray, i uintptr) *IoUringCqe {
	return (*IoUringCqe)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + SizeofIoUringCqe*i))
}

//

type userdata [8]byte // uint64

func (u *userdata) SetUint64(v uint64) {
	u.SetUintptr(uintptr(v))
}
func (u *userdata) SetUintptr(v uintptr) {
	putUintptr(unsafe.Pointer(u), v)
}
func (u *userdata) SetUnsafe(ptr unsafe.Pointer) {
	putUnsafe(unsafe.Pointer(u), ptr)
}

func (u userdata) GetUnsafe() unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&u))
}
func (u userdata) GetUintptr() uintptr {
	return uintptr(u.GetUnsafe())
}
func (u userdata) GetUint64() uint64 {
	return uint64(u.GetUintptr())
}
func (u userdata) IsZero() bool {
	return u.GetUint64() == 0
}

// ---

func putUnsafe(ptr unsafe.Pointer, v unsafe.Pointer) {
	*(*unsafe.Pointer)(ptr) = v
}

func putUintptr(ptr unsafe.Pointer, v uintptr) {
	*(*uintptr)(ptr) = v
}
func putUint64(ptr unsafe.Pointer, v uint64) {
	*(*uint64)(ptr) = v
}
func putUint32(ptr unsafe.Pointer, v uint32) {
	*(*uint32)(ptr) = v
}
func putUint16(ptr unsafe.Pointer, v uint16) {
	*(*uint16)(ptr) = v
}
func putUint8(ptr unsafe.Pointer, v uint8) {
	*(*uint8)(ptr) = v
}

func putInt64(ptr unsafe.Pointer, v int64) {
	*(*int64)(ptr) = v
}
func putInt32(ptr unsafe.Pointer, v int32) {
	*(*int32)(ptr) = v
}
func putInt16(ptr unsafe.Pointer, v int16) {
	*(*int16)(ptr) = v
}
func putInt8(ptr unsafe.Pointer, v int8) {
	*(*int8)(ptr) = v
}
