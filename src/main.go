package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/procfs"
	"github.com/reiver/go-cast"
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
func main() {

	// parse("testout.log")
	// os.Exit(0)
	// Test()
	// os.Exit(0)
	defer fmt.Println("Defferd")
	comm := NewComm()
	measure := NewMeasure(comm)
	// l := NewLogger[ParquetMessage](".", 5000, 10)
	l := NewDumper(".", 5000)
	l.Start()
	defer l.Stop()
	// a, _ := reflector.Reflect(&message{})
	// fmt.Println(a)
	// os.Exit(0)
	go measure.Start()
	go trigger(comm)
	go Recieve(comm, l)
	done := make(chan bool, 1)
	waitSig(done)
	<-done
}
func serialize(m message) {
	enc := json.NewEncoder(os.Stdout)
	err := enc.Encode(m)
	if err != nil {
		log.Fatal("encode:", err)
	}

	fmt.Println("datagram>>", "<<datagram")
}

func Recieve(c Comm, l Dumper) {

	for true {
		a := <-c.datagram
		// fmt.Println(a.Type, a.Time)
		// fmt.Println(a.ProcStat)
		// fmt.Println("Recieved")
		l.AddLog(transform(a))
		// transform(a)
	}
}

type ParquetMessage struct {
	Type     int      `parquet:"name=type, type=INT32"`
	Time     int64    `parquet:"name=time, type=INT64"`
	ProcStat ProcStat `parquet:"name=procstat, type=STRUCT"`
}

type ProcStat struct {
	PID                 int    `parquet:"name=pid, type=INT32"`
	Comm                string `parquet:"name=comm, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	State               string `parquet:"name=state, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Session             int    `parquet:"name=session, type=INT32"`
	Flags               int    `parquet:"name=flags, type=INT32, convertedtype=UINT_32"`
	MinFlt              int    `parquet:"name=minflt, type=INT32, convertedtype=UINT_32"`
	CMinFlt             int    `parquet:"name=cminflt, type=INT32, convertedtype=UINT_32"`
	MajFlt              int    `parquet:"name=majflt, type=INT32, convertedtype=UINT_32"`
	CMajFlt             int    `parquet:"name=cmajflt, type=INT32, convertedtype=UINT_32"`
	UTime               int    `parquet:"name=utime, type=INT32, convertedtype=UINT_32"`
	STime               int    `parquet:"name=stime, type=INT32, convertedtype=UINT_32"`
	CUTime              int    `parquet:"name=cutime, type=INT32"`
	CSTime              int    `parquet:"name=cstime, type=INT32"`
	Priority            int    `parquet:"name=priority, type=INT32"`
	Nice                int    `parquet:"name=nice, type=INT32"`
	NumThreads          int    `parquet:"name=numthreads, type=INT32"`
	Starttime           int64  `parquet:"name=starttime, type=INT64, convertedtype=UINT_64"`
	VSize               int    `parquet:"name=vsize, type=INT32, convertedtype=UINT_32"`
	RSS                 int    `parquet:"name=rss, type=INT32"`
	RSSLimit            int64  `parquet:"name=rsslimit, type=INT64, convertedtype=UINT_64"`
	Processor           int    `parquet:"name=processor, type=INT32, convertedtype=UINT_32"`
	RTPriority          int    `parquet:"name=rtpriority, type=INT32, convertedtype=UINT_32"`
	Policy              int    `parquet:"name=policy, type=INT32, convertedtype=UINT_32"`
	DelayAcctBlkIOTicks int64  `parquet:"name=delayacctblkioticks, type=INT64, convertedtype=UINT_64"`
	GuestTime           int    `parquet:"name=guesttime, type=INT32"`
	CGuestTime          int    `parquet:"name=cguesttime, type=INT32"`
}

func transform(m message) *ParquetMessage {
	tmp := ParquetMessage{Type: m.Type, Time: m.Time,
		ProcStat: ProcStat{
			PID:                 m.ProcStat.PID,
			Comm:                m.ProcStat.Comm,
			State:               m.ProcStat.State,
			Session:             m.ProcStat.Session,
			Flags:               cast.MustInt(m.ProcStat.Flags),
			MinFlt:              cast.MustInt(m.ProcStat.MinFlt),
			CMinFlt:             cast.MustInt(m.ProcStat.CMinFlt),
			MajFlt:              cast.MustInt(m.ProcStat.MajFlt),
			CMajFlt:             cast.MustInt(m.ProcStat.CMajFlt),
			UTime:               cast.MustInt(m.ProcStat.UTime),
			STime:               cast.MustInt(m.ProcStat.STime),
			CUTime:              m.ProcStat.CUTime,
			CSTime:              m.ProcStat.CSTime,
			Priority:            m.ProcStat.Priority,
			Nice:                m.ProcStat.Nice,
			NumThreads:          m.ProcStat.NumThreads,
			Starttime:           cast.MustInt64(m.ProcStat.Starttime),
			VSize:               cast.MustInt(m.ProcStat.VSize),
			RSS:                 m.ProcStat.RSS,
			RSSLimit:            cast.MustInt64(m.ProcStat.RSSLimit),
			Processor:           cast.MustInt(m.ProcStat.Processor),
			RTPriority:          cast.MustInt(m.ProcStat.RTPriority),
			Policy:              cast.MustInt(m.ProcStat.Policy),
			DelayAcctBlkIOTicks: cast.MustInt64(m.ProcStat.DelayAcctBlkIOTicks),
			GuestTime:           m.ProcStat.GuestTime,
			CGuestTime:          m.ProcStat.CGuestTime,
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

func (s *Measure) Start() {
	for true {
		<-s.comm.measFS
		a, _ := s.FS.AllProcs()
		for _, l1 := range a {
			tmp, _ := l1.Stat()
			datagram := message{Type: 1,
				Time:     time.Now().UnixMilli(),
				ProcStat: tmp}
			s.comm.datagram <- datagram

		}
	}
}
