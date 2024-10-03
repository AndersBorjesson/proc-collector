package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)

func readfile(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	d, err := io.ReadAll(f)
	return d
}

func parse(filename string) {
	rawdata := string(readfile(filename))
	messages := derivemessages(rawdata)
	analyticalData := reformat(messages)
	toParquet(analyticalData)
}

type AnalyticalForm struct {
	ProcStat Procstat
}

type Procstat struct {
	Type []int
	Time []int64
}

func reformat(messages []message) AnalyticalForm {
	data := AnalyticalForm{ProcStat: Procstat{Type: make([]int, 0, len(messages)),
		Time: make([]int64, 0, len(messages))}}
	for _, l1 := range messages {
		if l1.Type == 1 {
			data.ProcStat.Time = append(data.ProcStat.Time, l1.Time)
		}
	}
	return data
}

func toParquet(Data AnalyticalForm) {
	// Create a new Parquet file writer with schema definition

	fw, err := local.NewLocalFileWriter("example.parquet")
	if err != nil {
		log.Fatalln("Can't create local file", err)
	}
	defer fw.Close()
	pw, err := writer.NewParquetWriter(fw, new(AnalyticalForm), 4)
	if err != nil {
		log.Fatalln("Can't create parquet writer", err)
	}
	defer pw.WriteStop()
	fmt.Println(Data)
	if err = pw.Write(Data); err != nil {
		log.Fatalln("Write error : ", err)
	}
}
func derivemessages(rawdata string) []message {
	tmp := strings.Split(rawdata, "datagram>>")
	data := make([]message, 0, len(tmp))
	for _, v := range tmp {
		if strings.Contains(v, "<<datagram") {
			tmp2 := strings.Split(v, "<<datagram")
			datagram := message{}
			if err := json.Unmarshal([]byte(tmp2[0]), &datagram); err != nil {
				log.Println(err)
			} else {
				data = append(data, datagram)
			}
		}
	}
	return data
}
