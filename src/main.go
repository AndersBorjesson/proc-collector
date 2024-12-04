package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Collect(collectProcFsFreq int64, collectNetFreq int64) {

	comm := NewComm()
	measure := NewMeasure(comm)
	go measure.Start()

	sniffer := NewSniffer(comm)
	go sniffer.Start()
	defer sniffer.Close()

	l := NewDumper(".", 50000)
	l.Start()
	defer l.Stop()

	go trigger(comm.measFS, collectProcFsFreq, false, "procfs")
	go trigger(comm.measNet, collectNetFreq, false, "network")
	go Recieve(comm, &l)
	done := make(chan bool, 1)
	waitSig(done)
	<-done
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

		(*l).AddLog(transform(a))
	}
}

func transform(m message) *ParquetMessage {
	switch m.Type {
	case 1:
		tmp := ParquetMessage{Type: m.Type, Time: m.Time, RefTime: m.RefTime,
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
	case 2:
		tmp := ParquetMessage{Type: m.Type, Time: m.Time, RefTime: m.RefTime,
			ConnectionData: m.ConnectionData}
		return &tmp

	case 3:
		tmp := ParquetMessage{Type: m.Type, Time: m.Time, RefTime: m.RefTime,
			ConnectionInfo: ConnectionInfo{
				TotalConnections:     m.ConnectionInfo.TotalConnections,
				TotalDownloadBytes:   m.ConnectionInfo.TotalDownloadBytes,
				TotalUploadBytes:     m.ConnectionInfo.TotalUploadBytes,
				TotalDownloadPackets: m.ConnectionInfo.TotalDownloadPackets,
				TotalUploadPackets:   m.ConnectionInfo.TotalUploadPackets,
			}}
		return &tmp
	}
	return nil
}

func trigger(c chan bool, Freq int64, verbose bool, name string) {
	//works
	for true {
		time.Sleep(time.Duration(Freq) * time.Millisecond)
		c <- true
		if verbose {
			log.Println("Triggered sampling of ", name)
		}
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
