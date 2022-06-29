#include "include/exports.h"
#include <liburing.h>
#include <liburing/io_uring.h>

__uint64 IoUringQueueInit(
    __uint64 entries,
    struct io_uring *ring,
    __uint64 flags)
{
    return (__uint64)io_uring_queue_init(
        (unsigned)entries,
        ring,
        (unsigned)flags);
}

__uint64 IoUringQueueInitParams(
    __uint64 entries,
    struct io_uring *ring,
    struct io_uring_params *params)
{

    return (__uint64)io_uring_queue_init_params(
        (unsigned)entries,
        ring,
        params);

}

void IoUringQueueExit(
    struct io_uring *ring)
{
    io_uring_queue_exit(ring);
}

__uint64 IoUringSubmit(struct io_uring *ring)
{
    return (__uint64)io_uring_submit(ring);
}

__uint64 IoUringSubmitAndWait(
    struct io_uring *ring,
    __uint64 waitNr)
{
    return (__uint64)io_uring_submit_and_wait(
        ring, (unsigned int)waitNr);
}