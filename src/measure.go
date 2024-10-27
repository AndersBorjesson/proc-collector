package main

import (
	"fmt"
	"log"
	"time"

	"github.com/prometheus/procfs"
)

type Measure struct {
	FS   procfs.FS
	comm Comm
}

func NewMeasure(c Comm) Measure {
	FS, err := procfs.NewDefaultFS()
	if err != nil {
		fmt.Println(err)
	}
	return Measure{FS: FS,
		comm: c}
}

// tmp, err := l1.IO()
// tmp, err := l1.Schedstat()
// tmp, err := l1.NewStatus()
func PIO(pid procfs.Proc) procfs.ProcIO {
	tmp, err := pid.IO()
	if err != nil {
		log.Println(err)
	}
	return tmp
}

func PS(pid procfs.Proc) procfs.ProcStatus {
	tmp, err := pid.NewStatus()
	if err != nil {
		log.Println(err)
	}
	return tmp
}

func PSS(pid procfs.Proc) procfs.ProcSchedstat {
	tmp, err := pid.Schedstat()
	if err != nil {
		log.Println(err)
	}
	return tmp
}

func PStat(pid procfs.Proc) procfs.ProcStat {
	tmp, err := pid.Stat()
	if err != nil {
		log.Println(err)
	}
	return tmp
}
func (s *Measure) Start() {
	for true {
		<-s.comm.measFS
		a, _ := s.FS.AllProcs()
		for _, l1 := range a {

			datagram := message{Type: 1,
				Time:          time.Now().UnixMilli(),
				ProcStat:      PStat(l1),
				ProcStatus:    PS(l1),
				ProcIO:        PIO(l1),
				ProcSchedstat: PSS(l1)}
			s.comm.datagram <- datagram

		}
	}
}
