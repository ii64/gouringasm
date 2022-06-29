package gouringasm

import "unsafe"

type (
	__ushort    = uint16
	__uint      = uint32
	__uchar     = uint8
	__ulong     = uint32
	__ulonglong = uint64
)

type __kernel_rwf_t = __int
type __s32 = __int
type __u16 = __ushort
type __u32 = __uint
type __u64 = __ulonglong
type __u8 = __uchar
type __int int32
type __size_t = __ulong

type IoUring struct {
	sq            IoUringSQ
	cq            IoUringCQ
	flags         __uint
	ring_fd       __int
	features      __uint
	enter_ring_fd __int
	int_flags     __u8
	pad           [3]__u8
	pad2          __uint
}

type IoUringCQ struct {
	khead         *__uint
	ktail         *__uint
	kring_mask    *__uint
	kring_entries *__uint
	kflags        *__uint
	koverflow     *__uint
	cqes          *IoUringCQE
	ring_sz       __size_t
	ring_ptr      unsafe.Pointer
	pad           [4]__uint
}

type IoUringCQE struct {
	user_data __u64
	res       __s32
	flags     __u32
}

type IoUringSQ struct {
	khead         *__uint
	ktail         *__uint
	kring_mask    *__uint
	kring_entries *__uint
	kflags        *__uint
	kdropped      *__uint
	array         *__uint
	sqes          *IoUringSQE
	sqe_head      __uint
	sqe_tail      __uint
	ring_sz       __size_t
	ring_ptr      unsafe.Pointer
	pad           [4]__uint
}

type IoUringSQE struct {
	opcode      __u8
	flags       __u8
	ioprio      __u16
	fd          __s32
	anon0       [8]byte
	anon1       [8]byte
	len         __u32
	anon2       [4]byte
	user_data   __u64
	anon3       [2]byte
	personality __u16
	anon4       [4]byte
	addr3       __u64
	__pad2      [1]__u64
}

type IoUringParams struct {
	sq_entries     __u32
	cq_entries     __u32
	flags          __u32
	sq_thread_cpu  __u32
	sq_thread_idle __u32
	features       __u32
	wq_fd          __u32
	resv           [3]__u32
	sq_off         IoSQRingOffsets
	cq_off         IoCQRingOffsets
}

type IoSQRingOffsets struct {
	head         __u32
	tail         __u32
	ring_mask    __u32
	ring_entries __u32
	flags        __u32
	dropped      __u32
	array        __u32
	resv1        __u32
	resv2        __u64
}

type IoCQRingOffsets struct {
	head         __u32
	tail         __u32
	ring_mask    __u32
	ring_entries __u32
	overflow     __u32
	cqes         __u32
	flags        __u32
	resv1        __u32
	resv2        __u64
}
