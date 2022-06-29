#include "include/exports.h"
#include <liburing.h>
#include <liburing/io_uring.h>

void IoUringPrepRW(
    __uint64 op, struct io_uring_sqe *sqe,
    __uint64 fd, void *addr,
    __uint64 len, __uint64 offset)
{
    io_uring_prep_rw(
        (int)op,
        sqe, (int)fd,
        (const void *)addr,
        (unsigned int)len, (__u64)offset);
}

struct io_uring_sqe *IoUringGetSQE(
    struct io_uring *ring)
{
    return io_uring_get_sqe(ring);
}