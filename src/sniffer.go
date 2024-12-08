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
		refTime := time.Now().UnixMilli()
		datagram := message{Type: 3,
			Time:    time.Now().UnixMilli(),
			RefTime: refTime,
			ConnectionInfo: ConnectionInfo{
				TotalConnections:     b.TotalConnections,
				TotalDownloadBytes:   b.TotalDownloadBytes,
				TotalUploadBytes:     b.TotalUploadBytes,
				TotalDownloadPackets: b.TotalDownloadPackets,
				TotalUploadPackets:   b.TotalUploadPackets,
			}}
		s.comm.datagram <- datagram
		for _, i := range b.Connections {
			datagram := message{Type: 2,
				Time:           time.Now().UnixMilli(),
				RefTime:        refTime,
				ConnectionData: *i}

			s.comm.datagram <- datagram
		}

	}

}
