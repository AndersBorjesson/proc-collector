package main

import "github.com/prometheus/procfs"

type message struct {
	Type     int
	Time     int64
	ProcStat procfs.ProcStat
	Procs    procfs.Procs
}
type Comm struct {
	measFS   chan bool
	datagram chan message
}

func NewComm() Comm {
	return Comm{measFS: make(chan bool),
		datagram: make(chan message)}
}
