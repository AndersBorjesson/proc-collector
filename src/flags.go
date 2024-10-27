package main

import (
	"flag"
	"fmt"
	"os"
)

func flagParse() {
	convertCmd := flag.NewFlagSet("convert", flag.ExitOnError)
	convertPath := convertCmd.String("path", "./", "path where the memdump files are located")

	collectCmd := flag.NewFlagSet("collect", flag.ExitOnError)
	collectProcFsFreq := collectCmd.Int64("procfs", 1000, "delay between procfs reads in milliseconds")
	if len(os.Args) < 2 {
		fmt.Println("expected subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "convert":
		convertCmd.Parse(os.Args[2:])

		path := *convertPath
		ConvertMemdump2Json(path)
	case "collect":
		collectCmd.Parse(os.Args[2:])
		Collect(*collectProcFsFreq)

	}
}
