package main

import (
	"time"

	"github.com/AndersBorjesson/snifferlib"
)

type Sniffer struct {
	s    *snifferlib.SnifferLib
	comm Comm
}

func NewSniffer(c Comm) *Sniffer {
	return &Sniffer{s: snifferlib.NewSnifferLib(),
		comm: c}
}

func (s *Sniffer) Close() {
	s.s.Close()
}

func (s *Sniffer) Start() {
	for {
		<-s.comm.measNet
		b := s.s.GetStats()
		for _, i := range b.Connections {
			datagram := message{Type: 2,
				Time:           time.Now().UnixMilli(),
				ConnectionData: *i}

			s.comm.datagram <- datagram
		}

	}

}
