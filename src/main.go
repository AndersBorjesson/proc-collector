package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AndersBorjesson/snifferlib"
)

// func Test() {
// 	nf, err := netflow.New(
// 		netflow.WithCaptureTimeout(5 * time.Second),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = nf.Start()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer nf.Stop()

// 	<-nf.Done()

// 	var (
// 		limit     = 50
// 		recentSec = 5
// 	)

// 	rank, err := nf.GetProcessRank(limit, recentSec)
// 	if err != nil {
// 		panic(err)
// 	}

// 	bs, err := json.MarshalIndent(rank, "", "    ")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(string(bs))

// }

func Collect(collectProcFsFreq int64) {
	// ConvertMemdump2Json()
	// os.Exit(0)
	// parse("testout.log")
	// os.Exit(0)
	// Test()
	// os.Exit(0)
	defer fmt.Println("Defferd")
	comm := NewComm()
	measure := NewMeasure(comm)

	l := NewDumper(".", 50000)
	l.Start()
	defer l.Stop()
	// a, _ := reflector.Reflect(&message{})
	// fmt.Println(a)
	// os.Exit(0)
	go measure.Start()
	go trigger(comm)
	go Recieve(comm, &l)
	done := make(chan bool, 1)
	waitSig(done)
	<-done
}

func Sniff() {
	A := snifferlib.NewSnifferLib()
	defer A.Close()
	b := A.GetStats()
	fmt.Println(b)
	// Collect()
	time.Sleep(2 * time.Second)
	b = A.GetStats()
	fmt.Println(b)
}
func main() {

	flagParse()

}
func serialize(m message) {
	enc := json.NewEncoder(os.Stdout)
	err := enc.Encode(m)
	if err != nil {
		log.Fatal("encode:", err)
	}

	fmt.Println("datagram>>", "<<datagram")
}

func Recieve(c Comm, l *Dumper) {

	for true {
		a := <-c.datagram
		// fmt.Println(a.Type, a.Time)
		// fmt.Println(a.ProcStat)
		// fmt.Println("Recieved")
		(*l).AddLog(transform(a))
		// transform(a)
	}
}

func transform(m message) *ParquetMessage {
	tmp := ParquetMessage{Type: m.Type, Time: m.Time,
		ProcStat: ProcStat{
			PID:                 m.ProcStat.PID,
			Comm:                m.ProcStat.Comm,
			State:               m.ProcStat.State,
			Session:             m.ProcStat.Session,
			Flags:               m.ProcStat.Flags,
			MinFlt:              m.ProcStat.MinFlt,
			CMinFlt:             m.ProcStat.CMinFlt,
			MajFlt:              m.ProcStat.MajFlt,
			CMajFlt:             m.ProcStat.CMajFlt,
			UTime:               m.ProcStat.UTime,
			STime:               m.ProcStat.STime,
			CUTime:              m.ProcStat.CUTime,
			CSTime:              m.ProcStat.CSTime,
			Priority:            m.ProcStat.Priority,
			Nice:                m.ProcStat.Nice,
			NumThreads:          m.ProcStat.NumThreads,
			Starttime:           m.ProcStat.Starttime,
			VSize:               m.ProcStat.VSize,
			RSS:                 m.ProcStat.RSS,
			RSSLimit:            m.ProcStat.RSSLimit,
			Processor:           m.ProcStat.Processor,
			RTPriority:          m.ProcStat.RTPriority,
			Policy:              m.ProcStat.Policy,
			DelayAcctBlkIOTicks: m.ProcStat.DelayAcctBlkIOTicks,
			GuestTime:           m.ProcStat.GuestTime,
			CGuestTime:          m.ProcStat.CGuestTime,
		},
		ProcStatus: ProcStatus{
			PID:                      m.ProcStatus.PID,
			Name:                     m.ProcStatus.Name,
			TGID:                     m.ProcStatus.TGID,
			NSpids:                   m.ProcStatus.NSpids,
			VmPeak:                   m.ProcStatus.VmPeak,
			VmSize:                   m.ProcStatus.VmSize,
			VmLck:                    m.ProcStatus.VmLck,
			VmPin:                    m.ProcStatus.VmPin,
			VmHWM:                    m.ProcStatus.VmHWM,
			VmRSS:                    m.ProcStatus.VmRSS,
			RssAnon:                  m.ProcStatus.RssAnon,
			RssFile:                  m.ProcStatus.RssFile,
			RssShmem:                 m.ProcStatus.RssShmem,
			VmData:                   m.ProcStatus.VmData,
			VmStk:                    m.ProcStatus.VmStk,
			VmExe:                    m.ProcStatus.VmExe,
			VmLib:                    m.ProcStatus.VmLib,
			VmPTE:                    m.ProcStatus.VmPTE,
			VmPMD:                    m.ProcStatus.VmPMD,
			VmSwap:                   m.ProcStatus.VmSwap,
			HugetlbPages:             m.ProcStatus.HugetlbPages,
			VoluntaryCtxtSwitches:    m.ProcStatus.VoluntaryCtxtSwitches,
			NonVoluntaryCtxtSwitches: m.ProcStatus.NonVoluntaryCtxtSwitches,
			UIDs:                     m.ProcStatus.UIDs,
			GIDs:                     m.ProcStatus.GIDs,
			CpusAllowedList:          m.ProcStatus.CpusAllowedList,
		},
		ProcIO: ProcIO{
			RChar:               m.ProcIO.RChar,
			WChar:               m.ProcIO.WChar,
			SyscR:               m.ProcIO.SyscR,
			SyscW:               m.ProcIO.SyscW,
			ReadBytes:           m.ProcIO.ReadBytes,
			WriteBytes:          m.ProcIO.WriteBytes,
			CancelledWriteBytes: m.ProcIO.CancelledWriteBytes,
		},
		ProcSchedstat: ProcSchedstat{
			RunningNanoseconds: m.ProcSchedstat.RunningNanoseconds,
			WaitingNanoseconds: m.ProcSchedstat.WaitingNanoseconds,
			RunTimeslices:      m.ProcSchedstat.RunTimeslices,
		},
	}
	return &tmp
}
func trigger(c Comm) {
	//works
	for true {
		fmt.Println("triggering")

		time.Sleep(500 * time.Millisecond)
		c.measFS <- true
	}
}

func waitSig(done chan bool) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
}
