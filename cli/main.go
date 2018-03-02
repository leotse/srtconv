package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/leotse/srtfix"
)

func main() {

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
	log.Println(parsed)

	// resolve time overlap
	resolved := srtfix.Resolve(parsed)
	log.Println(resolved)
}
