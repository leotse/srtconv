package main

import (
	"fmt"
	"io/ioutil"
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

	// read file
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Panic(err)
	}

	// parse file
	content := string(file)
	parsed, err := srtfix.ParseSrtFile(content)
	if err != nil {
		log.Panic(err)
	}

	// resolve time overlap
	resolved := srtfix.Resolve(parsed)
	for _, caption := range resolved {
		fmt.Println(caption)
		fmt.Println()
	}
}

func usage() {
	fmt.Println()
	fmt.Println("\tUsage: srtfix [input_srt_path]")
	fmt.Println()
}
