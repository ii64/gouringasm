#ifndef _H_NATIVE
#define _H_NATIVE

#include <liburing/io_uring.h>
#include <stdint.h>
#include <liburing.h>

typedef int64_t __uint64;

typedef struct {
    int fd_in;  int64_t off_in;
    int fd_out; int64_t off_out;
    unsigned int nbytes;
    unsigned int splice_flags;
} PrepSpliceArgs;


#endif