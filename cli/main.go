package main

import (
	"fmt"
	"log"
	"os"

	"github.com/leotse/srtfix"
)

func main() {

	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	// file to read from arg
	filepath := os.Args[1]

	// init converter
	converter := srtfix.NewFileConverter(
		NewFileReader(filepath),
		NewFileWriter(filepath+".out"),
		NewParser(),
		NewResolver(),
	)

	err := converter.Convert()
	if err != nil {
		log.Panic(err)
	}
}

func usage() {
	fmt.Println()
	fmt.Println("\tUsage: srtfix [input_srt_path]")
	fmt.Println()
}
