#include "include/exports.h"
#include <liburing.h>
#include <liburing/io_uring.h>

struct io_uring_sqe *IoUringGetSQE(
    struct io_uring *ring)
{
    return io_uring_get_sqe(ring);
}