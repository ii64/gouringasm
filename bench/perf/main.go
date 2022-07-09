package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"syscall"
	"time"

	"github.com/ii64/gouringasm"
)

var fs = flag.NewFlagSet("perf", flag.ExitOnError)

var (
	entries      uint
	sqPoll       bool
	sqThreadCpu  uint
	sqThreadIdle uint

	N    uint
	noti uint

	pprofCpuFilename = "pprof.cpu"
)

func init() {
	fs.UintVar(&entries, "entries", 256, "Entries")
	fs.BoolVar(&sqPoll, "sqpoll", false, "Enable SQPOLL")
	fs.UintVar(&sqThreadCpu, "sqthreadcpu", 4, "SQ Thread CPU")
	fs.UintVar(&sqThreadIdle, "sqthreadidle", 10_000, "SQ Thread idle")

	fs.UintVar(&N, "n", 10_000, "N times")
	fs.UintVar(&noti, "noti", 10_000, "Notify per attempt N")

	fs.StringVar(&pprofCpuFilename, "pprofCpu", pprofCpuFilename, "pprof cpu output file")
}

func main() {
	err := fs.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	// check entries size
	if entries > uint(^uint32(0)) {
		panic("entries overflow.")
	}

	params := &gouringasm.IoUringParams{}
	if sqPoll {
		params.Flags |= gouringasm.IORING_SETUP_SQPOLL
		params.SqThreadCpu = uint32(sqThreadCpu)
		params.SqThreadIdle = uint32(sqThreadIdle)
	}

	h, err := gouringasm.NewWithParams(uint32(entries), params)
	if err != nil {
		panic(err)
	}
	defer h.Close()

	f, err := os.Create(pprofCpuFilename)
	if err != nil {
		panic(err)
	}

	fmt.Println("performing...")

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var i, j uint
	var sqe *gouringasm.IoUringSqe
	var cqe *gouringasm.IoUringCqe
	var submitted int

	startTime := time.Now()
	for i = 0; i < N; i++ {
		if i%noti == 0 { // notify
			fmt.Printf("n:%d e:%s\n", j, time.Now().Sub(startTime))
		}

		for j = 0; j < entries; j++ {
			for {
				// sqe could be nil if SQ is already full so we spin until we got one
				sqe = h.GetSqe()
				if sqe != nil {
					break
				}
				runtime.Gosched()
			}
			gouringasm.PrepNop(sqe)
			sqe.UserData.SetUint64(uint64(i + j))
		}
		submitted, err = h.Submit()
		if err != nil {
			panic(err)
		}

		if i%noti == 0 { // notify
			fmt.Printf(" >> submitted %d\n", submitted)
		}

		for j = 0; j < entries; j++ {
			err = h.WaitCqe(&cqe)
			if err == syscall.EINTR {
				continue
			}
			if err != nil {
				panic(err)
			}
			if cqe == nil {
				panic("cqe is nil!")
			}
			if cqe.Res < 0 {
				panic(syscall.Errno(-cqe.Res))
			}

			h.SeenCqe(cqe)
		}
	}
	_ = submitted
	_ = err
}
