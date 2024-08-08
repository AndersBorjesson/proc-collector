package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/procfs"
	"github.com/wirelessr/avroschema"
	"github.com/wirelessr/avroschema/mongo"
)

func main() {

	parse("testout.log")
	os.Exit(0)
	defer fmt.Println("Defferd")
	comm := NewComm()
	measure := NewMeasure(comm)
	reflector := new(avroschema.Reflector)
	reflector.Mapper = mongo.MgmExtension

	// a, _ := reflector.Reflect(&message{})
	// fmt.Println(a)
	// os.Exit(0)
	go measure.Start()
	go trigger(comm)
	go Recieve(comm)
	done := make(chan bool, 1)
	waitSig(done)
	<-done
}
func serialize(m message) {
	b, _ := json.Marshal(m)
	fmt.Println("datagram>>", string(b), "<<datagram")
}

func Recieve(c Comm) {

	for true {
		a := <-c.datagram
		fmt.Println(a.Type, a.Time)
		fmt.Println(a.ProcStat)
		serialize(a)
	}
}

func trigger(c Comm) {
	//works
	for true {
		fmt.Println("triggering")

		time.Sleep(100 * time.Millisecond)
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
