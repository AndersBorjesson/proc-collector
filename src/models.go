package main

import "github.com/prometheus/procfs"

type message struct {
	Type          int
	Time          int64
	ProcStat      procfs.ProcStat
	Procs         procfs.Procs
	ProcStatus    procfs.ProcStatus
	ProcIO        procfs.ProcIO
	ProcSchedstat procfs.ProcSchedstat
}
type Comm struct {
	measFS   chan bool
	datagram chan message
}

func NewComm() Comm {
	return Comm{measFS: make(chan bool),
		datagram: make(chan message, 1000)}
}

type ParquetMessage struct {
	Type          int           `parquet:"name=type, type=INT32"`
	Time          int64         `parquet:"name=time, type=INT64"`
	ProcStat      ProcStat      `parquet:"name=procstat, type=STRUCT"`
	ProcIO        ProcIO        `parquet:"name=procio, type=STRUCT"`
	ProcStatus    ProcStatus    `parquet:"name=procstatus, type=STRUCT"`
	ProcSchedstat ProcSchedstat `parquet:"name=procschedstat, type=STRUCT"`
}

type ProcStat struct {
	PID                 int    `parquet:"name=pid, type=INT32"`
	Comm                string `parquet:"name=comm, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	State               string `parquet:"name=state, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Session             int    `parquet:"name=session, type=INT32"`
	Flags               uint   `parquet:"name=flags, type=INT32, convertedtype=UINT_32"`
	MinFlt              uint   `parquet:"name=minflt, type=INT32, convertedtype=UINT_32"`
	CMinFlt             uint   `parquet:"name=cminflt, type=INT32, convertedtype=UINT_32"`
	MajFlt              uint   `parquet:"name=majflt, type=INT32, convertedtype=UINT_32"`
	CMajFlt             uint   `parquet:"name=cmajflt, type=INT32, convertedtype=UINT_32"`
	UTime               uint   `parquet:"name=utime, type=INT32, convertedtype=UINT_32"`
	STime               uint   `parquet:"name=stime, type=INT32, convertedtype=UINT_32"`
	CUTime              int    `parquet:"name=cutime, type=INT32"`
	CSTime              int    `parquet:"name=cstime, type=INT32"`
	Priority            int    `parquet:"name=priority, type=INT32"`
	Nice                int    `parquet:"name=nice, type=INT32"`
	NumThreads          int    `parquet:"name=numthreads, type=INT32"`
	Starttime           uint64 `parquet:"name=starttime, type=INT64, convertedtype=UINT_64"`
	VSize               uint   `parquet:"name=vsize, type=INT32, convertedtype=UINT_32"`
	RSS                 int    `parquet:"name=rss, type=INT32"`
	RSSLimit            uint64 `parquet:"name=rsslimit, type=INT64, convertedtype=UINT_64"`
	Processor           uint   `parquet:"name=processor, type=INT32, convertedtype=UINT_32"`
	RTPriority          uint   `parquet:"name=rtpriority, type=INT32, convertedtype=UINT_32"`
	Policy              uint   `parquet:"name=policy, type=INT32, convertedtype=UINT_32"`
	DelayAcctBlkIOTicks uint64 `parquet:"name=delayacctblkioticks, type=INT64, convertedtype=UINT_64"`
	GuestTime           int    `parquet:"name=guesttime, type=INT32"`
	CGuestTime          int    `parquet:"name=cguesttime, type=INT32"`
}

type ProcStatus struct {
	PID                      int    `parquet:"name=pid, type=INT32"`
	Name                     string `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	TGID                     int    `parquet:"name=tgid, type=INT32"`
	NSpids                   []uint64
	VmPeak                   uint64 // nolint:revive
	VmSize                   uint64 // nolint:revive
	VmLck                    uint64 // nolint:revive
	VmPin                    uint64 // nolint:revive
	VmHWM                    uint64 // nolint:revive
	VmRSS                    uint64 // nolint:revive
	RssAnon                  uint64 // nolint:revive
	RssFile                  uint64 // nolint:revive
	RssShmem                 uint64 // nolint:revive
	VmData                   uint64 // nolint:revive
	VmStk                    uint64 // nolint:revive
	VmExe                    uint64 // nolint:revive
	VmLib                    uint64 // nolint:revive
	VmPTE                    uint64 // nolint:revive
	VmPMD                    uint64 // nolint:revive
	VmSwap                   uint64 // nolint:revive
	HugetlbPages             uint64
	VoluntaryCtxtSwitches    uint64
	NonVoluntaryCtxtSwitches uint64
	UIDs                     [4]uint64
	GIDs                     [4]uint64
	CpusAllowedList          []uint64
}
type ProcIO struct {
	// Chars read.
	RChar uint64
	// Chars written.
	WChar uint64
	// Read syscalls.
	SyscR uint64
	// Write syscalls.
	SyscW uint64
	// Bytes read.
	ReadBytes uint64
	// Bytes written.
	WriteBytes uint64
	// Bytes written, but taking into account truncation. See
	// Documentation/filesystems/proc.txt in the kernel sources for
	// detailed explanation.
	CancelledWriteBytes int64
}
type ProcSchedstat struct {
	RunningNanoseconds uint64
	WaitingNanoseconds uint64
	RunTimeslices      uint64
}
