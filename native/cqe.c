#include "include/exports.h"
#include <liburing.h>
#include <liburing/io_uring.h>

__uint64 IoUringWaitCQE(
    struct io_uring *ring,
    struct io_uring_cqe **cqe_ptr)
{
    return (__uint64)io_uring_wait_cqe(
        ring, cqe_ptr);
}

void IoUringCQESeen(
    struct io_uring *ring,
    struct io_uring_cqe *cqe)
{
    io_uring_cqe_seen(ring, cqe);
}