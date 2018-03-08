package main

import (
	"fmt"
	"io/ioutil"

	"github.com/leotse/srtfix"
)

func NewFileReader(path string) *FileReader {
	return &FileReader{path: path}
}

func NewStdoutWriter() *StdoutWriter {
	return &StdoutWriter{}
}

func NewFileWriter(path string) *FileWriter {
	return &FileWriter{path: path}
}

func NewParser() *SrtParser {
	return &SrtParser{}
}

func NewResolver() *CaptionResolver {
	return &CaptionResolver{}
}

type FileReader struct {
	path string
}

func (reader *FileReader) Read() (string, error) {
	content, err := ioutil.ReadFile(reader.path)
	return string(content), err
}

type StdoutWriter struct {
}

func (writer *StdoutWriter) Write(content string) error {
	fmt.Println(content)
	return nil
}

type FileWriter struct {
	path string
}

func (writer *FileWriter) Write(content string) error {
	return ioutil.WriteFile(writer.path, []byte(content), 0644)
}

type SrtParser struct {
}

func (parser *SrtParser) Parse(content string) ([]*srtfix.Caption, error) {
	return srtfix.ParseSrtFile(content)
}

type CaptionResolver struct {
}

func (resolver *CaptionResolver) Resolve(captions []*srtfix.Caption) []*srtfix.Caption {
	return srtfix.Resolve(captions)
}
