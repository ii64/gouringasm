package evloop

import (
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"

	"github.com/ii64/gouringasm"
)

type EventloopUserdata struct {
	resolved uint64

	Res int32

	Callback func(*EventloopUserdata)
}

func (ud *EventloopUserdata) Resolve() {
	atomic.StoreUint64(&ud.resolved, 1)
	// no need ack.
}

func (ud *EventloopUserdata) Wait() {
	for !atomic.CompareAndSwapUint64(&ud.resolved, 1, 0) {
		runtime.Gosched()
	}
}

func (ud *EventloopUserdata) zero() {
	ud.resolved = 0
	ud.Res = 0
}
func (ud *EventloopUserdata) uptr() uint64 {
	return uint64(uintptr(unsafe.Pointer(ud)))
}

type Eventloop struct {
	ring               *gouringasm.Uring
	evloopUserdataPool sync.Pool
	spinCheck          *uint64
}

func New(ring *gouringasm.Uring) *Eventloop {
	evloop := &Eventloop{
		ring: ring,
		evloopUserdataPool: sync.Pool{
			New: func() any {
				return new(EventloopUserdata)
			},
		},
		spinCheck: new(uint64),
	}
	return evloop
}

func (ev *Eventloop) Close() {
	atomic.StoreUint64(ev.spinCheck, 1)
	// no need ack
}

func (ev *Eventloop) Submit() (int32, error) {
	return ev.ring.Submit()
}

func (ev *Eventloop) SubmitAndWait(waitNr uint32) (int32, error) {
	return ev.ring.SubmitAndWait(waitNr)
}

func (ev *Eventloop) Serve() (err error) {
	var cqe *gouringasm.IoUringCQE
	for !atomic.CompareAndSwapUint64(ev.spinCheck, 1, 0) {
		err = ev.ring.WaitCQE(&cqe)
		if err == syscall.EINTR {
			runtime.Gosched()
			continue
		} else if err != nil {
			return
		}

		ud := ev.peekUserdata(cqe.UserData)
		ev.handle(cqe, ud)
		ev.freeUserdata(ud)

		ev.ring.SeenCQE(cqe)
	}
	return
}

func (ev *Eventloop) handle(cqe *gouringasm.IoUringCQE, ud *EventloopUserdata) {
	if cqe == nil || ud == nil {
		return
	}

	ud.Res = int32(cqe.Res)
	ud.Resolve()
	if f := ud.Callback; f != nil {
		f(ud)
	}
}

// --- userdata ---

func (ev *Eventloop) GetUserdata() *EventloopUserdata {
	ud := ev.evloopUserdataPool.Get().(*EventloopUserdata)
	ud.zero()
	return ud
}

func (ev *Eventloop) peekUserdata(id uint64) *EventloopUserdata {
	if id == 0 {
		return nil
	}
	return (*EventloopUserdata)(unsafe.Pointer(uintptr(id)))
}

func (ev *Eventloop) freeUserdata(ud *EventloopUserdata) {
	if ud == nil {
		return
	}
	ev.evloopUserdataPool.Put(ud)
}

// --- sqe ---

func (ev *Eventloop) GetSQE() (*gouringasm.IoUringSQE, *EventloopUserdata) {
	for {
		sqe := ev.ring.GetSQE() // atomic op
		if sqe == nil {
			runtime.Gosched()
			continue
		}
		ud := ev.GetUserdata()
		gouringasm.IoUringSetUserdata(sqe, ud.uptr())
		return sqe, ud
	}
}
