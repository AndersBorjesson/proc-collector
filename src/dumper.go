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
	file      *os.File
	rfile     *os.File
	rreader   *bufio.Reader
	filepath  string
	data      []ParquetMessage
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
		filepath:  filepath,
		data:      make([]ParquetMessage, 0),
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
func (d *Dumper) AddLog(data *ParquetMessage) {
	// memdump.Encode(&d.buffer, data)
	d.data = append(d.data, *data)
	if len(d.data) > d.bufferlen {
		d.Flush()
	}
}

func (d *Dumper) GetLog(template interface{}) {
	memdump.Decode(&d.buffer, template)

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
	s.data = make([]ParquetMessage, 0)
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
