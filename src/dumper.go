package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/alexflint/go-memdump"
)

type Dumper struct {
	buffer    bytes.Buffer
	logDir    string
	bufferlen int
	currlen   int
	file      *os.File
	rfile     *os.File
	rreader   *bufio.Reader
	filepath  string
	data      PL
}

func NewDumper(logdir string, bufferlen int) Dumper {
	e, err := exists(logdir)
	if err != nil {
		log.Fatalln(err)
	}
	if !e {
		log.Fatalln("Log directory does not exist")
	}
	filepath := path.Join(logdir, GenerateParquetFileName("memdump"))
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return Dumper{buffer: bytes.Buffer{},
		logDir:    logdir,
		bufferlen: bufferlen,
		file:      file,
		currlen:   0,
		filepath:  filepath,
		data:      NewPL(bufferlen),
	}

}

func NewDumperFromFile(logfile string) Dumper {
	rfile, err := os.Open(logfile)
	if err != nil {
		log.Fatal(err)
	}
	buffer := bufio.NewReader(rfile)
	Dumper := Dumper{rfile: rfile,
		rreader: buffer}
	return Dumper
}

func (d *Dumper) Size() {
}
func (d *Dumper) AddLog(data *message) {
	// memdump.Encode(&d.buffer, data)
	AddData(*data, &d.data)
	d.currlen++
	if (d.currlen)%1000 == 0 {
		log.Println("Buffer length is ", d.currlen)
	}
	if d.currlen > d.bufferlen {
		d.Flush()
	}
}

func (d *Dumper) GetLog(template interface{}) {
	err := memdump.Decode(&d.buffer, template)
	if err != nil {
		log.Fatalln("GetLog failing :", err)
	}

}

func (d *Dumper) GetFromFile(template interface{}) {
	err := memdump.Decode(d.rfile, template)
	if err != nil {
		log.Fatalln("GetFromFile failing :", err)
	}
}

func (s *Dumper) Flush() {
	fmt.Println("Flushing")
	memdump.Encode(&s.buffer, &s.data)
	s.file.Write(s.buffer.Bytes())
	s.file.Close()
	filepath := path.Join(s.logDir, GenerateParquetFileName("memdump"))
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	s.file = file
	s.buffer.Reset()
	s.data = NewPL(s.bufferlen)
	s.currlen = 0
}
func (s *Dumper) Stop() {
	fmt.Println("stopping")
	memdump.Encode(&s.buffer, &s.data)
	s.file.Write(s.buffer.Bytes())
	s.file.Close()
}

func (s *Dumper) Close() {
	s.rfile.Close()
}
func (s *Dumper) Start() {
}
