#include "include/exports.h"
#include <bits/types/struct_iovec.h>
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

void IouringPrepSplice(struct io_uring_sqe *sqe, PrepSpliceArgs *args)
{
    io_uring_prep_splice(sqe,
        args->fd_in, args->off_in,
        args->fd_out,args->off_out,
        args->nbytes, args->splice_flags);
}

void IoUringPrepTee(struct io_uring_sqe *sqe,
    __uint64 fd_in, __uint64 fd_out,
    __uint64 nbytes, __uint64 splice_flags)
{
    io_uring_prep_tee(sqe,
        (int)fd_in,
        (int)fd_out, (unsigned int)nbytes,
        (unsigned int)splice_flags);
}



void IoUringPrepReadv(struct io_uring_sqe *sqe,
    __uint64 fd, struct iovec *iovecs,
    __uint64 nr_vecs, __uint64 offset)
{
    io_uring_prep_readv(sqe, (int)fd,
        (const struct iovec *)iovecs,
        (unsigned int)nr_vecs, (__u64)offset);
}
void IoUringPrepReadv2(struct io_uring_sqe *sqe,
    __uint64 fd, struct iovec *iovecs,
    __uint64 nr_vecs, __uint64 offset,
    __uint64 flags)
{
    io_uring_prep_readv2(sqe, (int)fd,
        (const struct iovec *)iovecs,
        (unsigned int)nr_vecs, (__u64)offset,
        (int)flags);
}
// io_uring_prep_read_fixed



void IoUringPrepWritev(struct io_uring_sqe *sqe,
    __uint64 fd, struct iovec *iovecs,
    __uint64 nr_vecs, __uint64 offset)
{
    io_uring_prep_writev(sqe, (int)fd,
        (const struct iovec *)iovecs,
        (unsigned int)nr_vecs, (__u64)offset);
}
void IoUringPrepWritev2(struct io_uring_sqe *sqe,
    __uint64 fd, struct iovec *iovecs,
    __uint64 nr_vecs, __uint64 offset,
    __uint64 flags)
{
    io_uring_prep_writev2(sqe,
        (int)fd,
        (const struct iovec *)iovecs,
        (unsigned int)nr_vecs, (__u64)offset,
        (int)flags);
}
// io_uring_prep_write_fixed


void IoUringPrepRecvmsg(struct io_uring_sqe *sqe,
    __uint64 fd, struct msghdr *msg, __uint64 flags)
{
    io_uring_prep_recvmsg(sqe,
        (int)fd, (struct msghdr *)msg, (unsigned int)flags);
}

void IoUringPrepSendmsg(struct io_uring_sqe *sqe,
    __uint64 fd, struct msghdr *msg, __uint64 flags)
{
    io_uring_prep_sendmsg(sqe,
        (int)fd, (const struct msghdr *)msg, (unsigned int)flags);
}

void IoUringPrepPollAdd(struct io_uring_sqe *sqe, __uint64 fd, __uint64 poll_mask)
{
    io_uring_prep_poll_add(sqe, (int)fd, (unsigned int)poll_mask);
}

void IoUringPrepPollMultishot(struct io_uring_sqe *sqe, __uint64 fd, __uint64 poll_mask)
{
    io_uring_prep_poll_multishot(sqe, (int)fd, (unsigned int)poll_mask);
}

void IoUringPrepPollRemove(struct io_uring_sqe *sqe, __uint64 user_data)
{
    io_uring_prep_poll_remove(sqe, user_data);
}
void IoUringPrepPollUpdate(struct io_uring_sqe *sqe, __uint64 old_user_data,
                                             __uint64 new_user_data,
                                             __uint64 poll_mask,
                                             __uint64 flags)
{
    io_uring_prep_poll_update(sqe, old_user_data, new_user_data,
        (unsigned int)poll_mask, (unsigned int)flags);
}

void IoUringPrepFsync(struct io_uring_sqe *sqe, __uint64 fd, __uint64 fsync_flags)
{
    io_uring_prep_fsync(sqe, (int)fd, (unsigned int)fsync_flags);
}

void IoUringPrepNop(struct io_uring_sqe *sqe)
{
    io_uring_prep_nop(sqe);
}


void IoUringPrepTimeout(struct io_uring_sqe *sqe,
                        struct __kernel_timespec *ts,
                        __uint64 count, __uint64 flags)
{
    io_uring_prep_timeout(sqe, ts, (unsigned int)count, (unsigned int)flags);
}

void IoUringPrepTimeoutRemove(struct io_uring_sqe *sqe,  __uint64 user_data,
                            __uint64 flags)
{
    io_uring_prep_timeout_remove(sqe, user_data, (unsigned int)flags);
}

void IoUringPrepTimeoutUpdate(struct io_uring_sqe *sqe,
                                struct __kernel_timespec *ts,
                                __uint64 user_data,
                                __uint64 flags)
{
    io_uring_prep_timeout_update(sqe, ts, user_data, (unsigned int)flags);
}

void IoUringPrepAccept(struct io_uring_sqe *sqe,
    __uint64 fd, struct sockaddr *addr,
    socklen_t *addrlen, __uint64 flags)
{
    io_uring_prep_accept(sqe, (int)fd, addr, addrlen, (int)flags);
}

// io_uring_prep_accept_direct
void IoUringPrepMultishotAccept(struct io_uring_sqe *sqe,
    __uint64 fd, struct sockaddr *addr,
    socklen_t *addrlen, __uint64 flags)
{
    io_uring_prep_multishot_accept(sqe, (int)fd,
        addr, addrlen, (int)flags);
}

void IoUringPrepMultishotAcceptDirect(struct io_uring_sqe *sqe, __uint64 fd,
                                      struct sockaddr *addr, socklen_t *addrlen,
                                      __uint64 flags)
{
    io_uring_prep_multishot_accept_direct(sqe, (int)fd, addr, addrlen, (int)flags);
}



void IoUringPrepCancel64(struct io_uring_sqe *sqe,
                        __uint64 user_data, __uint64 flags)
{
    io_uring_prep_cancel64(sqe, user_data, (int)flags);
}

void IoUringPrepCancel(struct io_uring_sqe *sqe,
                        void *user_data, __uint64 flags)
{
    io_uring_prep_cancel(sqe, user_data, (int)flags);
}

void IoUringPrepCancelFd(struct io_uring_sqe *sqe, __uint64 fd,
                            __uint64 flags)
{
    io_uring_prep_cancel_fd(sqe, (int)fd, (unsigned int)flags);
}

void IoUringPrepLinkTimeout(struct io_uring_sqe *sqe,
                            struct __kernel_timespec *ts,
                            __uint64 flags)
{
    io_uring_prep_link_timeout(sqe, ts, (unsigned int)flags);
}

void IoUringPrepConnect(struct io_uring_sqe *sqe,
    __uint64 fd,
    struct sockaddr *addr,
    __uint64 addrlen)
{
    io_uring_prep_connect(sqe,
        (int)fd, (const struct sockaddr *)addr, (socklen_t)addrlen);
}

void IoUringPrepFilesUpdate(struct io_uring_sqe *sqe,
                            int *fds, __uint64 nr_fds,
                            __uint64 offset)
{
    io_uring_prep_files_update(sqe, fds, (unsigned int)nr_fds, (int)offset);
}

void IoUringPrepOpenat(struct io_uring_sqe *sqe, __uint64 dfd,
                        const char *path, __uint64 flags,
                        __uint64 mode)
{
    io_uring_prep_openat(sqe, (int)dfd, path, (int)flags, (mode_t)mode);
}

void IoUringPrepOpenatDirect(struct io_uring_sqe *sqe,
                            __uint64 dfd, const char *path,
                            __uint64 flags, __uint64 mode,
                            __uint64 file_index)
{
    io_uring_prep_openat_direct(sqe, (int)dfd, path, (int)flags, (mode_t)mode, (unsigned int)file_index);
}

void IoUringPrepClose(struct io_uring_sqe *sqe,
    __uint64 fd)
{
    io_uring_prep_close(sqe, (int)fd);
}

void IoUringPrepCloseDirect(struct io_uring_sqe *sqe,
                            __uint64 file_index)
{
    io_uring_prep_close_direct(sqe, (unsigned int)file_index);
}


void IoUringPrepRead(struct io_uring_sqe *sqe,
    __uint64 fd, void *buf, __uint64 nbytes,
    __uint64 offset)
{
    io_uring_prep_read(sqe, (int)fd, buf, (unsigned int)nbytes, (__u64)offset);
}

void IoUringPrepWrite(struct io_uring_sqe *sqe,
    __uint64 fd, void *buf, __uint64 nbytes,
    __uint64 offset)
{
    io_uring_prep_write(sqe, (int)fd, buf, (unsigned int)nbytes, (__u64)offset);
}


void IoUringPrepStatx(struct io_uring_sqe *sqe, __uint64 dfd,
                        const char *path, __uint64 flags,
                        __uint64 mask,
                        struct statx *statxbuf)
{
    io_uring_prep_statx(sqe, (int)dfd, path, (int)flags, (unsigned int)mask, statxbuf);
}

void IoUringPrepFadvise(struct io_uring_sqe *sqe, int fd,
                        __uint64 offset, __uint64 len, __uint64 advice)
{
    io_uring_prep_fadvise(sqe, fd, (__u64)offset, (off_t)len, (int)advice);
}

void IoUringPrepMadvise(struct io_uring_sqe *sqe, void *addr,
                        __uint64 length, __uint64 advice)
{
    io_uring_prep_madvise(sqe, addr, (off_t)length, (int)advice);
}

void IoUringPrepSend(struct io_uring_sqe *sqe,
    __uint64 sockfd, void *buf, size_t len, __uint64 flags)
{
    io_uring_prep_send(sqe, (int)sockfd,
        (const void *)buf, len, (int)flags);
}

void IoUringPrepRecv(struct io_uring_sqe *sqe,
    __uint64 sockfd, void *buf,
    __uint64 len, __uint64 flags)
{
    io_uring_prep_recv(sqe, (int)sockfd,
        buf, (size_t)len, (int)flags);
}


void IoUringPrepOpenat2(struct io_uring_sqe *sqe, __uint64 dfd,
                        const char *path, struct open_how *how)
{
    io_uring_prep_openat2(sqe, (int)dfd, path, how);
}

void IoUringPrepOpenat2Direct(struct io_uring_sqe *sqe,
                                __uint64 dfd, const char *path,
                                struct open_how *how,
                                __uint64 file_index)
{
    io_uring_prep_openat2_direct(sqe, (int)dfd, path, how, (unsigned int)file_index);
}

void IoUringPrepEpollCtl(struct io_uring_sqe *sqe,
    __uint64 epfd, __uint64 fd, __uint64 op,
    struct epoll_event *ev)
{
    io_uring_prep_epoll_ctl(sqe, (int)epfd, (int)fd,
        (int)op, ev);
}

// io_uring_prep_provide_buffers
// io_uring_prep_remove_buffers
// io_uring_prep_shutdown
// io_uring_prep_unlinkat
// io_uring_prep_unlink
// io_uring_prep_renameat
// io_uring_prep_rename
// io_uring_prep_sync_file_range
// io_uring_prep_mkdirat
// io_uring_prep_mkdir
// io_uring_prep_symlinkat
// io_uring_prep_symlink
// io_uring_prep_linkat
// io_uring_prep_link
// io_uring_prep_msg_ring
// io_uring_prep_getxattr
// io_uring_prep_setxattr
// io_uring_prep_fgetxattr
// io_uring_prep_fsetxattr
// io_uring_prep_socket
// io_uring_prep_socket_direct
// io_uring_prep_socket_direct_alloc