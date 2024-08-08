package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
	derivemessages(rawdata)

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
	fmt.Println(len(data))
	return data
}
