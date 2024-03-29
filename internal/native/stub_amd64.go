// +build !noasm !appengine
// Code generated by golinker, DO NOT EDIT.
// Command: golinker -fallback-rawbytes-x86 -rawbytes-x86 -extsymstub -stub=./internal/native/stub.go -out=./internal/native -entryname=__native_entry__ ./native/dist/libnative-amd64.a
package native

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __native_entry__() uintptr

var (
	__subr_io_uring_queue_mmap = __native_entry__() + 16
	__subr_io_uring_mmap = __native_entry__() + 160
	__subr_io_uring_ring_dontfork = __native_entry__() + 1168
	__subr___sys_madvise = __native_entry__() + 1520
	__subr_io_uring_queue_init_params = __native_entry__() + 1584
	__subr_____sys_io_uring_setup = __native_entry__() + 1728
	__subr___sys_close = __native_entry__() + 1776
	__subr_io_uring_queue_init = __native_entry__() + 1824
	__subr_io_uring_queue_exit = __native_entry__() + 1904
	__subr___sys_munmap = __native_entry__() + 2080
	__subr_io_uring_unmap_rings = __native_entry__() + 2128
	__subr_io_uring_get_probe_ring = __native_entry__() + 2240
	__subr_uring_malloc = __native_entry__() + 2384
	__subr_uring_free = __native_entry__() + 2416
	__subr_io_uring_get_probe = __native_entry__() + 2448
	__subr_io_uring_free_probe = __native_entry__() + 2576
	__subr_io_uring_mlock_size_params = __native_entry__() + 2608
	__subr_roundup_pow2 = __native_entry__() + 3056
	__subr_get_page_size = __native_entry__() + 3104
	__subr_rings_size = __native_entry__() + 3120
	__subr_io_uring_mlock_size = __native_entry__() + 3360
	__subr___sys_mmap = __native_entry__() + 3424
	__subr_IS_ERR = __native_entry__() + 3536
	__subr_PTR_ERR = __native_entry__() + 3584
	__subr___fls = __native_entry__() + 3600
	__subr_npages = __native_entry__() + 3664
	__subr___io_uring_get_cqe = __native_entry__() + 3728
	__subr__io_uring_get_cqe = __native_entry__() + 3824
	__subr_io_uring_peek_batch_cqe = __native_entry__() + 4256
	__subr_io_uring_cq_ready = __native_entry__() + 4656
	__subr_cq_ring_needs_flush = __native_entry__() + 4704
	__subr_____sys_io_uring_enter = __native_entry__() + 4752
	__subr___io_uring_flush_sq = __native_entry__() + 4816
	__subr_io_uring_wait_cqes = __native_entry__() + 5008
	__subr_io_uring_wait_cqes_new = __native_entry__() + 5200
	__subr___io_uring_submit_timeout = __native_entry__() + 5328
	__subr_io_uring_submit_and_wait_timeout = __native_entry__() + 5504
	__subr_io_uring_wait_cqe_timeout = __native_entry__() + 5760
	__subr_io_uring_submit = __native_entry__() + 5824
	__subr___io_uring_submit_and_wait = __native_entry__() + 5856
	__subr_io_uring_submit_and_wait = __native_entry__() + 5920
	__subr_io_uring_get_sqe = __native_entry__() + 5968
	__subr__io_uring_get_sqe = __native_entry__() + 6000
	__subr___io_uring_sqring_wait = __native_entry__() + 6192
	__subr___io_uring_peek_cqe = __native_entry__() + 6288
	__subr_cq_ring_needs_enter = __native_entry__() + 6640
	__subr_sq_ring_needs_enter = __native_entry__() + 6720
	__subr_____sys_io_uring_enter2 = __native_entry__() + 6848
	__subr_io_uring_cq_advance = __native_entry__() + 6944
	__subr_io_uring_prep_timeout = __native_entry__() + 7008
	__subr_io_uring_prep_rw = __native_entry__() + 7088
	__subr___io_uring_submit = __native_entry__() + 7264
	__subr_io_uring_register_buffers_update_tag = __native_entry__() + 7456
	__subr_____sys_io_uring_register = __native_entry__() + 7568
	__subr_io_uring_register_buffers_tags = __native_entry__() + 7632
	__subr_io_uring_register_buffers_sparse = __native_entry__() + 7728
	__subr_io_uring_register_buffers = __native_entry__() + 7808
	__subr_io_uring_unregister_buffers = __native_entry__() + 7904
	__subr_io_uring_register_files_update_tag = __native_entry__() + 8000
	__subr_io_uring_register_files_update = __native_entry__() + 8112
	__subr_io_uring_register_files_sparse = __native_entry__() + 8192
	__subr_increase_rlimit_nofile = __native_entry__() + 8368
	__subr_io_uring_register_files_tags = __native_entry__() + 8480
	__subr_io_uring_register_files = __native_entry__() + 8672
	__subr_io_uring_unregister_files = __native_entry__() + 8816
	__subr_io_uring_register_eventfd = __native_entry__() + 8912
	__subr_io_uring_unregister_eventfd = __native_entry__() + 9008
	__subr_io_uring_register_eventfd_async = __native_entry__() + 9104
	__subr_io_uring_register_probe = __native_entry__() + 9200
	__subr_io_uring_register_personality = __native_entry__() + 9296
	__subr_io_uring_unregister_personality = __native_entry__() + 9344
	__subr_io_uring_register_restrictions = __native_entry__() + 9392
	__subr_io_uring_enable_rings = __native_entry__() + 9488
	__subr_io_uring_register_iowq_aff = __native_entry__() + 9536
	__subr_io_uring_unregister_iowq_aff = __native_entry__() + 9600
	__subr_io_uring_register_iowq_max_workers = __native_entry__() + 9648
	__subr_io_uring_register_ring_fd = __native_entry__() + 9712
	__subr_io_uring_unregister_ring_fd = __native_entry__() + 9840
	__subr_io_uring_register_buf_ring = __native_entry__() + 9984
	__subr_io_uring_unregister_buf_ring = __native_entry__() + 10048
	__subr___sys_getrlimit = __native_entry__() + 10128
	__subr___sys_setrlimit = __native_entry__() + 10176
	__subr_memset = __native_entry__() + 10224
	__subr___uring_malloc = __native_entry__() + 10320
	__subr___sys_mmap_1 = __native_entry__() + 10448
	__subr_IS_ERR_1 = __native_entry__() + 10560
	__subr___uring_free = __native_entry__() + 10608
	__subr___sys_munmap_1 = __native_entry__() + 10720
	__subr_IoUringQueueInit = __native_entry__() + 10768
	__subr_IoUringQueueInitParams = __native_entry__() + 10784
	__subr_IoUringQueueExit = __native_entry__() + 10800
	__subr_IoUringSubmit = __native_entry__() + 10816
	__subr_IoUringSubmitAndWait = __native_entry__() + 10832
	__subr_IoUringWaitCQE = __native_entry__() + 10848
	__subr_IoUringCQESeen = __native_entry__() + 11008
	__subr_IoUringPrepRW = __native_entry__() + 11040
	__subr_IouringPrepSplice = __native_entry__() + 11104
	__subr_IoUringPrepTee = __native_entry__() + 11184
	__subr_IoUringPrepReadv = __native_entry__() + 11232
	__subr_IoUringPrepReadv2 = __native_entry__() + 11296
	__subr_IoUringPrepWritev = __native_entry__() + 11344
	__subr_IoUringPrepWritev2 = __native_entry__() + 11408
	__subr_IoUringPrepRecvmsg = __native_entry__() + 11456
	__subr_IoUringPrepSendmsg = __native_entry__() + 11520
	__subr_IoUringPrepPollAdd = __native_entry__() + 11584
	__subr_IoUringPrepPollMultishot = __native_entry__() + 11632
	__subr_IoUringPrepPollRemove = __native_entry__() + 11680
	__subr_IoUringPrepPollUpdate = __native_entry__() + 11744
	__subr_IoUringPrepFsync = __native_entry__() + 11808
	__subr_IoUringPrepNop = __native_entry__() + 11856
	__subr_IoUringPrepTimeout = __native_entry__() + 11920
	__subr_IoUringPrepTimeoutRemove = __native_entry__() + 11984
	__subr_IoUringPrepTimeoutUpdate = __native_entry__() + 12048
	__subr_IoUringPrepAccept = __native_entry__() + 12112
	__subr_IoUringPrepMultishotAccept = __native_entry__() + 12176
	__subr_IoUringPrepMultishotAcceptDirect = __native_entry__() + 12240
	__subr_IoUringPrepCancel64 = __native_entry__() + 12304
	__subr_IoUringPrepCancel = __native_entry__() + 12368
	__subr_IoUringPrepCancelFd = __native_entry__() + 12432
	__subr_IoUringPrepLinkTimeout = __native_entry__() + 12496
	__subr_IoUringPrepConnect = __native_entry__() + 12560
	__subr_IoUringPrepFilesUpdate = __native_entry__() + 12624
	__subr_IoUringPrepOpenat = __native_entry__() + 12688
	__subr_IoUringPrepOpenatDirect = __native_entry__() + 12752
	__subr_IoUringPrepClose = __native_entry__() + 12816
	__subr_IoUringPrepCloseDirect = __native_entry__() + 12864
	__subr_IoUringPrepRead = __native_entry__() + 12912
	__subr_IoUringPrepWrite = __native_entry__() + 12976
	__subr_IoUringPrepStatx = __native_entry__() + 13040
	__subr_IoUringPrepFadvise = __native_entry__() + 13088
	__subr_IoUringPrepMadvise = __native_entry__() + 13152
	__subr_IoUringPrepSend = __native_entry__() + 13216
	__subr_IoUringPrepRecv = __native_entry__() + 13280
	__subr_IoUringPrepOpenat2 = __native_entry__() + 13344
	__subr_IoUringPrepOpenat2Direct = __native_entry__() + 13392
	__subr_IoUringPrepEpollCtl = __native_entry__() + 13456
	__subr_IoUringGetSQE = __native_entry__() + 13520
)

const (
	__stack_io_uring_queue_mmap = 32
	__stack_io_uring_mmap = 48
	__stack_io_uring_ring_dontfork = 32
	__stack___sys_madvise = 0
	__stack_io_uring_queue_init_params = 32
	__stack_____sys_io_uring_setup = 0
	__stack___sys_close = 0
	__stack_io_uring_queue_init = 144
	__stack_io_uring_queue_exit = 32
	__stack___sys_munmap = 0
	__stack_io_uring_unmap_rings = 16
	__stack_io_uring_get_probe_ring = 48
	__stack_uring_malloc = 16
	__stack_uring_free = 16
	__stack_io_uring_get_probe = 240
	__stack_io_uring_free_probe = 16
	__stack_io_uring_mlock_size_params = 384
	__stack_roundup_pow2 = 16
	__stack_get_page_size = 0
	__stack_rings_size = 48
	__stack_io_uring_mlock_size = 128
	__stack___sys_mmap = 0
	__stack_IS_ERR = 0
	__stack_PTR_ERR = 0
	__stack___fls = 0
	__stack_npages = 16
	__stack___io_uring_get_cqe = 64
	__stack__io_uring_get_cqe = 64
	__stack_io_uring_peek_batch_cqe = 64
	__stack_io_uring_cq_ready = 0
	__stack_cq_ring_needs_flush = 0
	__stack_____sys_io_uring_enter = 32
	__stack___io_uring_flush_sq = 0
	__stack_io_uring_wait_cqes = 64
	__stack_io_uring_wait_cqes_new = 96
	__stack___io_uring_submit_timeout = 48
	__stack_io_uring_submit_and_wait_timeout = 112
	__stack_io_uring_wait_cqe_timeout = 32
	__stack_io_uring_submit = 16
	__stack___io_uring_submit_and_wait = 32
	__stack_io_uring_submit_and_wait = 16
	__stack_io_uring_get_sqe = 16
	__stack__io_uring_get_sqe = 0
	__stack___io_uring_sqring_wait = 16
	__stack___io_uring_peek_cqe = 64
	__stack_cq_ring_needs_enter = 16
	__stack_sq_ring_needs_enter = 0
	__stack_____sys_io_uring_enter2 = 0
	__stack_io_uring_cq_advance = 0
	__stack_io_uring_prep_timeout = 32
	__stack_io_uring_prep_rw = 0
	__stack___io_uring_submit = 32
	__stack_io_uring_register_buffers_update_tag = 80
	__stack_____sys_io_uring_register = 0
	__stack_io_uring_register_buffers_tags = 64
	__stack_io_uring_register_buffers_sparse = 48
	__stack_io_uring_register_buffers = 32
	__stack_io_uring_unregister_buffers = 16
	__stack_io_uring_register_files_update_tag = 80
	__stack_io_uring_register_files_update = 48
	__stack_io_uring_register_files_sparse = 64
	__stack_increase_rlimit_nofile = 32
	__stack_io_uring_register_files_tags = 80
	__stack_io_uring_register_files = 32
	__stack_io_uring_unregister_files = 16
	__stack_io_uring_register_eventfd = 32
	__stack_io_uring_unregister_eventfd = 16
	__stack_io_uring_register_eventfd_async = 32
	__stack_io_uring_register_probe = 32
	__stack_io_uring_register_personality = 16
	__stack_io_uring_unregister_personality = 16
	__stack_io_uring_register_restrictions = 32
	__stack_io_uring_enable_rings = 16
	__stack_io_uring_register_iowq_aff = 32
	__stack_io_uring_unregister_iowq_aff = 16
	__stack_io_uring_register_iowq_max_workers = 16
	__stack_io_uring_register_ring_fd = 32
	__stack_io_uring_unregister_ring_fd = 32
	__stack_io_uring_register_buf_ring = 32
	__stack_io_uring_unregister_buf_ring = 64
	__stack___sys_getrlimit = 0
	__stack___sys_setrlimit = 0
	__stack_memset = 0
	__stack___uring_malloc = 32
	__stack___sys_mmap_1 = 0
	__stack_IS_ERR_1 = 0
	__stack___uring_free = 32
	__stack___sys_munmap_1 = 0
	__stack_IoUringQueueInit = 0
	__stack_IoUringQueueInitParams = 0
	__stack_IoUringQueueExit = 0
	__stack_IoUringSubmit = 0
	__stack_IoUringSubmitAndWait = 0
	__stack_IoUringWaitCQE = 0
	__stack_IoUringCQESeen = 0
	__stack_IoUringPrepRW = 0
	__stack_IouringPrepSplice = 0
	__stack_IoUringPrepTee = 0
	__stack_IoUringPrepReadv = 0
	__stack_IoUringPrepReadv2 = 0
	__stack_IoUringPrepWritev = 0
	__stack_IoUringPrepWritev2 = 0
	__stack_IoUringPrepRecvmsg = 0
	__stack_IoUringPrepSendmsg = 0
	__stack_IoUringPrepPollAdd = 0
	__stack_IoUringPrepPollMultishot = 0
	__stack_IoUringPrepPollRemove = 0
	__stack_IoUringPrepPollUpdate = 0
	__stack_IoUringPrepFsync = 0
	__stack_IoUringPrepNop = 0
	__stack_IoUringPrepTimeout = 0
	__stack_IoUringPrepTimeoutRemove = 0
	__stack_IoUringPrepTimeoutUpdate = 0
	__stack_IoUringPrepAccept = 0
	__stack_IoUringPrepMultishotAccept = 0
	__stack_IoUringPrepMultishotAcceptDirect = 0
	__stack_IoUringPrepCancel64 = 0
	__stack_IoUringPrepCancel = 0
	__stack_IoUringPrepCancelFd = 0
	__stack_IoUringPrepLinkTimeout = 0
	__stack_IoUringPrepConnect = 0
	__stack_IoUringPrepFilesUpdate = 0
	__stack_IoUringPrepOpenat = 0
	__stack_IoUringPrepOpenatDirect = 0
	__stack_IoUringPrepClose = 0
	__stack_IoUringPrepCloseDirect = 0
	__stack_IoUringPrepRead = 0
	__stack_IoUringPrepWrite = 0
	__stack_IoUringPrepStatx = 0
	__stack_IoUringPrepFadvise = 0
	__stack_IoUringPrepMadvise = 0
	__stack_IoUringPrepSend = 0
	__stack_IoUringPrepRecv = 0
	__stack_IoUringPrepOpenat2 = 0
	__stack_IoUringPrepOpenat2Direct = 0
	__stack_IoUringPrepEpollCtl = 0
	__stack_IoUringGetSQE = 0
)

const (
	_ = __stack_io_uring_queue_mmap
	_ = __stack_io_uring_mmap
	_ = __stack_io_uring_ring_dontfork
	_ = __stack___sys_madvise
	_ = __stack_io_uring_queue_init_params
	_ = __stack_____sys_io_uring_setup
	_ = __stack___sys_close
	_ = __stack_io_uring_queue_init
	_ = __stack_io_uring_queue_exit
	_ = __stack___sys_munmap
	_ = __stack_io_uring_unmap_rings
	_ = __stack_io_uring_get_probe_ring
	_ = __stack_uring_malloc
	_ = __stack_uring_free
	_ = __stack_io_uring_get_probe
	_ = __stack_io_uring_free_probe
	_ = __stack_io_uring_mlock_size_params
	_ = __stack_roundup_pow2
	_ = __stack_get_page_size
	_ = __stack_rings_size
	_ = __stack_io_uring_mlock_size
	_ = __stack___sys_mmap
	_ = __stack_IS_ERR
	_ = __stack_PTR_ERR
	_ = __stack___fls
	_ = __stack_npages
	_ = __stack___io_uring_get_cqe
	_ = __stack__io_uring_get_cqe
	_ = __stack_io_uring_peek_batch_cqe
	_ = __stack_io_uring_cq_ready
	_ = __stack_cq_ring_needs_flush
	_ = __stack_____sys_io_uring_enter
	_ = __stack___io_uring_flush_sq
	_ = __stack_io_uring_wait_cqes
	_ = __stack_io_uring_wait_cqes_new
	_ = __stack___io_uring_submit_timeout
	_ = __stack_io_uring_submit_and_wait_timeout
	_ = __stack_io_uring_wait_cqe_timeout
	_ = __stack_io_uring_submit
	_ = __stack___io_uring_submit_and_wait
	_ = __stack_io_uring_submit_and_wait
	_ = __stack_io_uring_get_sqe
	_ = __stack__io_uring_get_sqe
	_ = __stack___io_uring_sqring_wait
	_ = __stack___io_uring_peek_cqe
	_ = __stack_cq_ring_needs_enter
	_ = __stack_sq_ring_needs_enter
	_ = __stack_____sys_io_uring_enter2
	_ = __stack_io_uring_cq_advance
	_ = __stack_io_uring_prep_timeout
	_ = __stack_io_uring_prep_rw
	_ = __stack___io_uring_submit
	_ = __stack_io_uring_register_buffers_update_tag
	_ = __stack_____sys_io_uring_register
	_ = __stack_io_uring_register_buffers_tags
	_ = __stack_io_uring_register_buffers_sparse
	_ = __stack_io_uring_register_buffers
	_ = __stack_io_uring_unregister_buffers
	_ = __stack_io_uring_register_files_update_tag
	_ = __stack_io_uring_register_files_update
	_ = __stack_io_uring_register_files_sparse
	_ = __stack_increase_rlimit_nofile
	_ = __stack_io_uring_register_files_tags
	_ = __stack_io_uring_register_files
	_ = __stack_io_uring_unregister_files
	_ = __stack_io_uring_register_eventfd
	_ = __stack_io_uring_unregister_eventfd
	_ = __stack_io_uring_register_eventfd_async
	_ = __stack_io_uring_register_probe
	_ = __stack_io_uring_register_personality
	_ = __stack_io_uring_unregister_personality
	_ = __stack_io_uring_register_restrictions
	_ = __stack_io_uring_enable_rings
	_ = __stack_io_uring_register_iowq_aff
	_ = __stack_io_uring_unregister_iowq_aff
	_ = __stack_io_uring_register_iowq_max_workers
	_ = __stack_io_uring_register_ring_fd
	_ = __stack_io_uring_unregister_ring_fd
	_ = __stack_io_uring_register_buf_ring
	_ = __stack_io_uring_unregister_buf_ring
	_ = __stack___sys_getrlimit
	_ = __stack___sys_setrlimit
	_ = __stack_memset
	_ = __stack___uring_malloc
	_ = __stack___sys_mmap_1
	_ = __stack_IS_ERR_1
	_ = __stack___uring_free
	_ = __stack___sys_munmap_1
	_ = __stack_IoUringQueueInit
	_ = __stack_IoUringQueueInitParams
	_ = __stack_IoUringQueueExit
	_ = __stack_IoUringSubmit
	_ = __stack_IoUringSubmitAndWait
	_ = __stack_IoUringWaitCQE
	_ = __stack_IoUringCQESeen
	_ = __stack_IoUringPrepRW
	_ = __stack_IouringPrepSplice
	_ = __stack_IoUringPrepTee
	_ = __stack_IoUringPrepReadv
	_ = __stack_IoUringPrepReadv2
	_ = __stack_IoUringPrepWritev
	_ = __stack_IoUringPrepWritev2
	_ = __stack_IoUringPrepRecvmsg
	_ = __stack_IoUringPrepSendmsg
	_ = __stack_IoUringPrepPollAdd
	_ = __stack_IoUringPrepPollMultishot
	_ = __stack_IoUringPrepPollRemove
	_ = __stack_IoUringPrepPollUpdate
	_ = __stack_IoUringPrepFsync
	_ = __stack_IoUringPrepNop
	_ = __stack_IoUringPrepTimeout
	_ = __stack_IoUringPrepTimeoutRemove
	_ = __stack_IoUringPrepTimeoutUpdate
	_ = __stack_IoUringPrepAccept
	_ = __stack_IoUringPrepMultishotAccept
	_ = __stack_IoUringPrepMultishotAcceptDirect
	_ = __stack_IoUringPrepCancel64
	_ = __stack_IoUringPrepCancel
	_ = __stack_IoUringPrepCancelFd
	_ = __stack_IoUringPrepLinkTimeout
	_ = __stack_IoUringPrepConnect
	_ = __stack_IoUringPrepFilesUpdate
	_ = __stack_IoUringPrepOpenat
	_ = __stack_IoUringPrepOpenatDirect
	_ = __stack_IoUringPrepClose
	_ = __stack_IoUringPrepCloseDirect
	_ = __stack_IoUringPrepRead
	_ = __stack_IoUringPrepWrite
	_ = __stack_IoUringPrepStatx
	_ = __stack_IoUringPrepFadvise
	_ = __stack_IoUringPrepMadvise
	_ = __stack_IoUringPrepSend
	_ = __stack_IoUringPrepRecv
	_ = __stack_IoUringPrepOpenat2
	_ = __stack_IoUringPrepOpenat2Direct
	_ = __stack_IoUringPrepEpollCtl
	_ = __stack_IoUringGetSQE
)
var (
	_ = __subr_io_uring_queue_mmap
	_ = __subr_io_uring_mmap
	_ = __subr_io_uring_ring_dontfork
	_ = __subr___sys_madvise
	_ = __subr_io_uring_queue_init_params
	_ = __subr_____sys_io_uring_setup
	_ = __subr___sys_close
	_ = __subr_io_uring_queue_init
	_ = __subr_io_uring_queue_exit
	_ = __subr___sys_munmap
	_ = __subr_io_uring_unmap_rings
	_ = __subr_io_uring_get_probe_ring
	_ = __subr_uring_malloc
	_ = __subr_uring_free
	_ = __subr_io_uring_get_probe
	_ = __subr_io_uring_free_probe
	_ = __subr_io_uring_mlock_size_params
	_ = __subr_roundup_pow2
	_ = __subr_get_page_size
	_ = __subr_rings_size
	_ = __subr_io_uring_mlock_size
	_ = __subr___sys_mmap
	_ = __subr_IS_ERR
	_ = __subr_PTR_ERR
	_ = __subr___fls
	_ = __subr_npages
	_ = __subr___io_uring_get_cqe
	_ = __subr__io_uring_get_cqe
	_ = __subr_io_uring_peek_batch_cqe
	_ = __subr_io_uring_cq_ready
	_ = __subr_cq_ring_needs_flush
	_ = __subr_____sys_io_uring_enter
	_ = __subr___io_uring_flush_sq
	_ = __subr_io_uring_wait_cqes
	_ = __subr_io_uring_wait_cqes_new
	_ = __subr___io_uring_submit_timeout
	_ = __subr_io_uring_submit_and_wait_timeout
	_ = __subr_io_uring_wait_cqe_timeout
	_ = __subr_io_uring_submit
	_ = __subr___io_uring_submit_and_wait
	_ = __subr_io_uring_submit_and_wait
	_ = __subr_io_uring_get_sqe
	_ = __subr__io_uring_get_sqe
	_ = __subr___io_uring_sqring_wait
	_ = __subr___io_uring_peek_cqe
	_ = __subr_cq_ring_needs_enter
	_ = __subr_sq_ring_needs_enter
	_ = __subr_____sys_io_uring_enter2
	_ = __subr_io_uring_cq_advance
	_ = __subr_io_uring_prep_timeout
	_ = __subr_io_uring_prep_rw
	_ = __subr___io_uring_submit
	_ = __subr_io_uring_register_buffers_update_tag
	_ = __subr_____sys_io_uring_register
	_ = __subr_io_uring_register_buffers_tags
	_ = __subr_io_uring_register_buffers_sparse
	_ = __subr_io_uring_register_buffers
	_ = __subr_io_uring_unregister_buffers
	_ = __subr_io_uring_register_files_update_tag
	_ = __subr_io_uring_register_files_update
	_ = __subr_io_uring_register_files_sparse
	_ = __subr_increase_rlimit_nofile
	_ = __subr_io_uring_register_files_tags
	_ = __subr_io_uring_register_files
	_ = __subr_io_uring_unregister_files
	_ = __subr_io_uring_register_eventfd
	_ = __subr_io_uring_unregister_eventfd
	_ = __subr_io_uring_register_eventfd_async
	_ = __subr_io_uring_register_probe
	_ = __subr_io_uring_register_personality
	_ = __subr_io_uring_unregister_personality
	_ = __subr_io_uring_register_restrictions
	_ = __subr_io_uring_enable_rings
	_ = __subr_io_uring_register_iowq_aff
	_ = __subr_io_uring_unregister_iowq_aff
	_ = __subr_io_uring_register_iowq_max_workers
	_ = __subr_io_uring_register_ring_fd
	_ = __subr_io_uring_unregister_ring_fd
	_ = __subr_io_uring_register_buf_ring
	_ = __subr_io_uring_unregister_buf_ring
	_ = __subr___sys_getrlimit
	_ = __subr___sys_setrlimit
	_ = __subr_memset
	_ = __subr___uring_malloc
	_ = __subr___sys_mmap_1
	_ = __subr_IS_ERR_1
	_ = __subr___uring_free
	_ = __subr___sys_munmap_1
	_ = __subr_IoUringQueueInit
	_ = __subr_IoUringQueueInitParams
	_ = __subr_IoUringQueueExit
	_ = __subr_IoUringSubmit
	_ = __subr_IoUringSubmitAndWait
	_ = __subr_IoUringWaitCQE
	_ = __subr_IoUringCQESeen
	_ = __subr_IoUringPrepRW
	_ = __subr_IouringPrepSplice
	_ = __subr_IoUringPrepTee
	_ = __subr_IoUringPrepReadv
	_ = __subr_IoUringPrepReadv2
	_ = __subr_IoUringPrepWritev
	_ = __subr_IoUringPrepWritev2
	_ = __subr_IoUringPrepRecvmsg
	_ = __subr_IoUringPrepSendmsg
	_ = __subr_IoUringPrepPollAdd
	_ = __subr_IoUringPrepPollMultishot
	_ = __subr_IoUringPrepPollRemove
	_ = __subr_IoUringPrepPollUpdate
	_ = __subr_IoUringPrepFsync
	_ = __subr_IoUringPrepNop
	_ = __subr_IoUringPrepTimeout
	_ = __subr_IoUringPrepTimeoutRemove
	_ = __subr_IoUringPrepTimeoutUpdate
	_ = __subr_IoUringPrepAccept
	_ = __subr_IoUringPrepMultishotAccept
	_ = __subr_IoUringPrepMultishotAcceptDirect
	_ = __subr_IoUringPrepCancel64
	_ = __subr_IoUringPrepCancel
	_ = __subr_IoUringPrepCancelFd
	_ = __subr_IoUringPrepLinkTimeout
	_ = __subr_IoUringPrepConnect
	_ = __subr_IoUringPrepFilesUpdate
	_ = __subr_IoUringPrepOpenat
	_ = __subr_IoUringPrepOpenatDirect
	_ = __subr_IoUringPrepClose
	_ = __subr_IoUringPrepCloseDirect
	_ = __subr_IoUringPrepRead
	_ = __subr_IoUringPrepWrite
	_ = __subr_IoUringPrepStatx
	_ = __subr_IoUringPrepFadvise
	_ = __subr_IoUringPrepMadvise
	_ = __subr_IoUringPrepSend
	_ = __subr_IoUringPrepRecv
	_ = __subr_IoUringPrepOpenat2
	_ = __subr_IoUringPrepOpenat2Direct
	_ = __subr_IoUringPrepEpollCtl
	_ = __subr_IoUringGetSQE
)
