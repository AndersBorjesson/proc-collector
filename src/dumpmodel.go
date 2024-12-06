package main

type PL struct {
	ProcStat       ProcStat_gen
	ProcStatus     ProcStatus_gen
	ProcIO         ProcIO_gen
	Time           []int64
	RefTime        []int64
	ProcSchedstat  ProcSchedstat_gen
	ConnectionData ConnectionData_gen
	ConnectionInfo ConnectionInfo_gen
	Type           []int
}
type ProcStat_gen struct {
	PPID                []int
	TTY                 []int
	CMinFlt             []uint
	Priority            []int
	Processor           []uint
	Policy              []uint
	UTime               []uint
	VSize               []uint
	RSS                 []int
	RTPriority          []uint
	Flags               []uint
	MajFlt              []uint
	Nice                []int
	DelayAcctBlkIOTicks []uint64
	Starttime           []uint64
	RSSLimit            []uint64
	CGuestTime          []int
	Comm                []string
	CMajFlt             []uint
	CSTime              []int
	MinFlt              []uint
	PGRP                []int
	Session             []int
	TPGID               []int
	STime               []uint
	GuestTime           []int
	PID                 []int
	State               []string
	CUTime              []int
	NumThreads          []int
}
type ProcStatus_gen struct {
	PID                      []int
	VmPeak                   []uint64
	VmHWM                    []uint64
	VmExe                    []uint64
	VmLib                    []uint64
	VmPTE                    []uint64
	VmPMD                    []uint64
	VmSize                   []uint64
	RssAnon                  []uint64
	VmData                   []uint64
	VmSwap                   []uint64
	VoluntaryCtxtSwitches    []uint64
	VmPin                    []uint64
	VmRSS                    []uint64
	RssFile                  []uint64
	Name                     []string
	TGID                     []int
	VmLck                    []uint64
	RssShmem                 []uint64
	VmStk                    []uint64
	HugetlbPages             []uint64
	NonVoluntaryCtxtSwitches []uint64
}
type ProcIO_gen struct {
	RChar               []uint64
	WChar               []uint64
	SyscR               []uint64
	SyscW               []uint64
	ReadBytes           []uint64
	WriteBytes          []uint64
	CancelledWriteBytes []int64
}
type ProcSchedstat_gen struct {
	RunningNanoseconds []uint64
	WaitingNanoseconds []uint64
	RunTimeslices      []uint64
}
type ConnectionData_gen struct {
	ProcessId       []int32
	DownloadBytes   []int
	UploadBytes     []int
	UploadPackets   []int
	DownloadPackets []int
	ProcessName     []string
	InterfaceName   []string
}
type ConnectionInfo_gen struct {
	TotalDownloadPackets []int
	TotalUploadPackets   []int
	TotalConnections     []int
	TotalDownloadBytes   []int
	TotalUploadBytes     []int
}

func NewPL(length int) PL {
	return PL{
		Type:    make([]int, 0, length),
		RefTime: make([]int64, 0, length),
		ProcSchedstat: ProcSchedstat_gen{
			RunningNanoseconds: make([]uint64, 0, length),
			WaitingNanoseconds: make([]uint64, 0, length),
			RunTimeslices:      make([]uint64, 0, length),
		},
		ConnectionData: ConnectionData_gen{
			ProcessName:     make([]string, 0, length),
			InterfaceName:   make([]string, 0, length),
			ProcessId:       make([]int32, 0, length),
			DownloadBytes:   make([]int, 0, length),
			UploadBytes:     make([]int, 0, length),
			UploadPackets:   make([]int, 0, length),
			DownloadPackets: make([]int, 0, length),
		},
		ConnectionInfo: ConnectionInfo_gen{
			TotalDownloadPackets: make([]int, 0, length),
			TotalUploadPackets:   make([]int, 0, length),
			TotalConnections:     make([]int, 0, length),
			TotalDownloadBytes:   make([]int, 0, length),
			TotalUploadBytes:     make([]int, 0, length),
		},
		Time: make([]int64, 0, length),
		ProcStat: ProcStat_gen{
			UTime:               make([]uint, 0, length),
			VSize:               make([]uint, 0, length),
			RSS:                 make([]int, 0, length),
			RTPriority:          make([]uint, 0, length),
			Flags:               make([]uint, 0, length),
			MajFlt:              make([]uint, 0, length),
			Nice:                make([]int, 0, length),
			DelayAcctBlkIOTicks: make([]uint64, 0, length),
			Starttime:           make([]uint64, 0, length),
			RSSLimit:            make([]uint64, 0, length),
			CGuestTime:          make([]int, 0, length),
			Comm:                make([]string, 0, length),
			CMajFlt:             make([]uint, 0, length),
			CSTime:              make([]int, 0, length),
			MinFlt:              make([]uint, 0, length),
			PGRP:                make([]int, 0, length),
			Session:             make([]int, 0, length),
			TPGID:               make([]int, 0, length),
			STime:               make([]uint, 0, length),
			GuestTime:           make([]int, 0, length),
			PID:                 make([]int, 0, length),
			State:               make([]string, 0, length),
			CUTime:              make([]int, 0, length),
			NumThreads:          make([]int, 0, length),
			Policy:              make([]uint, 0, length),
			PPID:                make([]int, 0, length),
			TTY:                 make([]int, 0, length),
			CMinFlt:             make([]uint, 0, length),
			Priority:            make([]int, 0, length),
			Processor:           make([]uint, 0, length),
		},
		ProcStatus: ProcStatus_gen{
			VmHWM:                    make([]uint64, 0, length),
			VmExe:                    make([]uint64, 0, length),
			VmLib:                    make([]uint64, 0, length),
			VmPTE:                    make([]uint64, 0, length),
			VmPMD:                    make([]uint64, 0, length),
			PID:                      make([]int, 0, length),
			VmPeak:                   make([]uint64, 0, length),
			VmData:                   make([]uint64, 0, length),
			VmSwap:                   make([]uint64, 0, length),
			VoluntaryCtxtSwitches:    make([]uint64, 0, length),
			VmSize:                   make([]uint64, 0, length),
			RssAnon:                  make([]uint64, 0, length),
			RssFile:                  make([]uint64, 0, length),
			VmPin:                    make([]uint64, 0, length),
			VmRSS:                    make([]uint64, 0, length),
			VmLck:                    make([]uint64, 0, length),
			RssShmem:                 make([]uint64, 0, length),
			VmStk:                    make([]uint64, 0, length),
			HugetlbPages:             make([]uint64, 0, length),
			NonVoluntaryCtxtSwitches: make([]uint64, 0, length),
			Name:                     make([]string, 0, length),
			TGID:                     make([]int, 0, length),
		},
		ProcIO: ProcIO_gen{
			WriteBytes:          make([]uint64, 0, length),
			CancelledWriteBytes: make([]int64, 0, length),
			RChar:               make([]uint64, 0, length),
			WChar:               make([]uint64, 0, length),
			SyscR:               make([]uint64, 0, length),
			SyscW:               make([]uint64, 0, length),
			ReadBytes:           make([]uint64, 0, length),
		},
	}
}
func AddData(data message, ref *PL) {
	(*ref).RefTime = append((*ref).RefTime, data.RefTime)
	(*ref).ProcSchedstat.RunningNanoseconds = append((*ref).ProcSchedstat.RunningNanoseconds, data.ProcSchedstat.RunningNanoseconds)
	(*ref).ProcSchedstat.WaitingNanoseconds = append((*ref).ProcSchedstat.WaitingNanoseconds, data.ProcSchedstat.WaitingNanoseconds)
	(*ref).ProcSchedstat.RunTimeslices = append((*ref).ProcSchedstat.RunTimeslices, data.ProcSchedstat.RunTimeslices)
	(*ref).ConnectionData.DownloadBytes = append((*ref).ConnectionData.DownloadBytes, data.ConnectionData.DownloadBytes)
	(*ref).ConnectionData.UploadBytes = append((*ref).ConnectionData.UploadBytes, data.ConnectionData.UploadBytes)
	(*ref).ConnectionData.UploadPackets = append((*ref).ConnectionData.UploadPackets, data.ConnectionData.UploadPackets)
	(*ref).ConnectionData.DownloadPackets = append((*ref).ConnectionData.DownloadPackets, data.ConnectionData.DownloadPackets)
	(*ref).ConnectionData.ProcessName = append((*ref).ConnectionData.ProcessName, data.ConnectionData.ProcessName)
	(*ref).ConnectionData.InterfaceName = append((*ref).ConnectionData.InterfaceName, data.ConnectionData.InterfaceName)
	(*ref).ConnectionData.ProcessId = append((*ref).ConnectionData.ProcessId, data.ConnectionData.ProcessId)
	(*ref).ConnectionInfo.TotalDownloadBytes = append((*ref).ConnectionInfo.TotalDownloadBytes, data.ConnectionInfo.TotalDownloadBytes)
	(*ref).ConnectionInfo.TotalUploadBytes = append((*ref).ConnectionInfo.TotalUploadBytes, data.ConnectionInfo.TotalUploadBytes)
	(*ref).ConnectionInfo.TotalDownloadPackets = append((*ref).ConnectionInfo.TotalDownloadPackets, data.ConnectionInfo.TotalDownloadPackets)
	(*ref).ConnectionInfo.TotalUploadPackets = append((*ref).ConnectionInfo.TotalUploadPackets, data.ConnectionInfo.TotalUploadPackets)
	(*ref).ConnectionInfo.TotalConnections = append((*ref).ConnectionInfo.TotalConnections, data.ConnectionInfo.TotalConnections)
	(*ref).Type = append((*ref).Type, data.Type)
	(*ref).ProcStat.PGRP = append((*ref).ProcStat.PGRP, data.ProcStat.PGRP)
	(*ref).ProcStat.Session = append((*ref).ProcStat.Session, data.ProcStat.Session)
	(*ref).ProcStat.TPGID = append((*ref).ProcStat.TPGID, data.ProcStat.TPGID)
	(*ref).ProcStat.STime = append((*ref).ProcStat.STime, data.ProcStat.STime)
	(*ref).ProcStat.GuestTime = append((*ref).ProcStat.GuestTime, data.ProcStat.GuestTime)
	(*ref).ProcStat.PID = append((*ref).ProcStat.PID, data.ProcStat.PID)
	(*ref).ProcStat.State = append((*ref).ProcStat.State, data.ProcStat.State)
	(*ref).ProcStat.CUTime = append((*ref).ProcStat.CUTime, data.ProcStat.CUTime)
	(*ref).ProcStat.NumThreads = append((*ref).ProcStat.NumThreads, data.ProcStat.NumThreads)
	(*ref).ProcStat.Policy = append((*ref).ProcStat.Policy, data.ProcStat.Policy)
	(*ref).ProcStat.PPID = append((*ref).ProcStat.PPID, data.ProcStat.PPID)
	(*ref).ProcStat.TTY = append((*ref).ProcStat.TTY, data.ProcStat.TTY)
	(*ref).ProcStat.CMinFlt = append((*ref).ProcStat.CMinFlt, data.ProcStat.CMinFlt)
	(*ref).ProcStat.Priority = append((*ref).ProcStat.Priority, data.ProcStat.Priority)
	(*ref).ProcStat.Processor = append((*ref).ProcStat.Processor, data.ProcStat.Processor)
	(*ref).ProcStat.UTime = append((*ref).ProcStat.UTime, data.ProcStat.UTime)
	(*ref).ProcStat.VSize = append((*ref).ProcStat.VSize, data.ProcStat.VSize)
	(*ref).ProcStat.RSS = append((*ref).ProcStat.RSS, data.ProcStat.RSS)
	(*ref).ProcStat.RTPriority = append((*ref).ProcStat.RTPriority, data.ProcStat.RTPriority)
	(*ref).ProcStat.Flags = append((*ref).ProcStat.Flags, data.ProcStat.Flags)
	(*ref).ProcStat.MajFlt = append((*ref).ProcStat.MajFlt, data.ProcStat.MajFlt)
	(*ref).ProcStat.Nice = append((*ref).ProcStat.Nice, data.ProcStat.Nice)
	(*ref).ProcStat.DelayAcctBlkIOTicks = append((*ref).ProcStat.DelayAcctBlkIOTicks, data.ProcStat.DelayAcctBlkIOTicks)
	(*ref).ProcStat.Starttime = append((*ref).ProcStat.Starttime, data.ProcStat.Starttime)
	(*ref).ProcStat.RSSLimit = append((*ref).ProcStat.RSSLimit, data.ProcStat.RSSLimit)
	(*ref).ProcStat.CGuestTime = append((*ref).ProcStat.CGuestTime, data.ProcStat.CGuestTime)
	(*ref).ProcStat.Comm = append((*ref).ProcStat.Comm, data.ProcStat.Comm)
	(*ref).ProcStat.CMajFlt = append((*ref).ProcStat.CMajFlt, data.ProcStat.CMajFlt)
	(*ref).ProcStat.CSTime = append((*ref).ProcStat.CSTime, data.ProcStat.CSTime)
	(*ref).ProcStat.MinFlt = append((*ref).ProcStat.MinFlt, data.ProcStat.MinFlt)
	(*ref).ProcStatus.HugetlbPages = append((*ref).ProcStatus.HugetlbPages, data.ProcStatus.HugetlbPages)
	(*ref).ProcStatus.NonVoluntaryCtxtSwitches = append((*ref).ProcStatus.NonVoluntaryCtxtSwitches, data.ProcStatus.NonVoluntaryCtxtSwitches)
	(*ref).ProcStatus.Name = append((*ref).ProcStatus.Name, data.ProcStatus.Name)
	(*ref).ProcStatus.TGID = append((*ref).ProcStatus.TGID, data.ProcStatus.TGID)
	(*ref).ProcStatus.VmLck = append((*ref).ProcStatus.VmLck, data.ProcStatus.VmLck)
	(*ref).ProcStatus.RssShmem = append((*ref).ProcStatus.RssShmem, data.ProcStatus.RssShmem)
	(*ref).ProcStatus.VmStk = append((*ref).ProcStatus.VmStk, data.ProcStatus.VmStk)
	(*ref).ProcStatus.VmPTE = append((*ref).ProcStatus.VmPTE, data.ProcStatus.VmPTE)
	(*ref).ProcStatus.VmPMD = append((*ref).ProcStatus.VmPMD, data.ProcStatus.VmPMD)
	(*ref).ProcStatus.PID = append((*ref).ProcStatus.PID, data.ProcStatus.PID)
	(*ref).ProcStatus.VmPeak = append((*ref).ProcStatus.VmPeak, data.ProcStatus.VmPeak)
	(*ref).ProcStatus.VmHWM = append((*ref).ProcStatus.VmHWM, data.ProcStatus.VmHWM)
	(*ref).ProcStatus.VmExe = append((*ref).ProcStatus.VmExe, data.ProcStatus.VmExe)
	(*ref).ProcStatus.VmLib = append((*ref).ProcStatus.VmLib, data.ProcStatus.VmLib)
	(*ref).ProcStatus.VmSize = append((*ref).ProcStatus.VmSize, data.ProcStatus.VmSize)
	(*ref).ProcStatus.RssAnon = append((*ref).ProcStatus.RssAnon, data.ProcStatus.RssAnon)
	(*ref).ProcStatus.VmData = append((*ref).ProcStatus.VmData, data.ProcStatus.VmData)
	(*ref).ProcStatus.VmSwap = append((*ref).ProcStatus.VmSwap, data.ProcStatus.VmSwap)
	(*ref).ProcStatus.VoluntaryCtxtSwitches = append((*ref).ProcStatus.VoluntaryCtxtSwitches, data.ProcStatus.VoluntaryCtxtSwitches)
	(*ref).ProcStatus.VmPin = append((*ref).ProcStatus.VmPin, data.ProcStatus.VmPin)
	(*ref).ProcStatus.VmRSS = append((*ref).ProcStatus.VmRSS, data.ProcStatus.VmRSS)
	(*ref).ProcStatus.RssFile = append((*ref).ProcStatus.RssFile, data.ProcStatus.RssFile)
	(*ref).ProcIO.ReadBytes = append((*ref).ProcIO.ReadBytes, data.ProcIO.ReadBytes)
	(*ref).ProcIO.WriteBytes = append((*ref).ProcIO.WriteBytes, data.ProcIO.WriteBytes)
	(*ref).ProcIO.CancelledWriteBytes = append((*ref).ProcIO.CancelledWriteBytes, data.ProcIO.CancelledWriteBytes)
	(*ref).ProcIO.RChar = append((*ref).ProcIO.RChar, data.ProcIO.RChar)
	(*ref).ProcIO.WChar = append((*ref).ProcIO.WChar, data.ProcIO.WChar)
	(*ref).ProcIO.SyscR = append((*ref).ProcIO.SyscR, data.ProcIO.SyscR)
	(*ref).ProcIO.SyscW = append((*ref).ProcIO.SyscW, data.ProcIO.SyscW)
	(*ref).Time = append((*ref).Time, data.Time)
}
