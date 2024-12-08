package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func listFiles(folder string) []string {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	memdumps := make([]string, 0)
	for _, file := range files {
		if !file.IsDir() {
			if strings.Contains(file.Name(), ".memdump") {
				memdump := folder + "/" + file.Name()
				memdumps = append(memdumps, memdump)
			}
		}

	}
	return memdumps
}

func ConvertMemdump2Json(path string) {
	memdumps := listFiles(path)
	for _, memdump := range memdumps {
		Dumper := NewDumperFromFile(memdump)
		defer Dumper.Close()
		decoded1 := NewPL(50000)
		decoded := &decoded1
		Dumper.GetFromFile(&decoded)
		rfile, err := os.Create(memdump + ".json")
		if err != nil {
			log.Fatal(err)
		}
		defer rfile.Close()
		enc := json.NewEncoder(rfile)
		err = enc.Encode(*decoded)
		if err != nil {
			log.Fatal(err)
		}
	}
}
