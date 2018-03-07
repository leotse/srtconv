package srtfix

import (
	"bytes"
)

// NewFileConverter creates a FileConverter instance
func NewFileConverter(
	reader FileReader,
	writer FileWriter,
	parser Parser,
	resolver Resolver,
) *FileConverter {
	return &FileConverter{
		reader:   reader,
		writer:   writer,
		parser:   parser,
		resolver: resolver,
	}
}

// FileConverter converts an input srt file into a normalized output
type FileConverter struct {
	reader   FileReader
	writer   FileWriter
	parser   Parser
	resolver Resolver
}

// FileReader can read a file
type FileReader interface {
	Read() (string, error)
}

// FileWriter can write content to a file
type FileWriter interface {
	Write(content string) error
}

// Parser can parse a srt file
type Parser interface {
	Parse(content string) ([]*Caption, error)
}

// Resolver can resolve overlapping times in captions
type Resolver interface {
	Resolve(captions []*Caption) []*Caption
}

// Convert read srt file from reader, and output normalized srt to writer
func (c *FileConverter) Convert() error {

	content, err := c.reader.Read()
	if err != nil {
		return err
	}

	captions, err := c.parser.Parse(content)
	if err != nil {
		return err
	}

	resolved := c.resolver.Resolve(captions)

	// build output srt content
	var buffer bytes.Buffer
	for _, caption := range resolved {
		buffer.WriteString(caption.String())
		buffer.WriteString("\n\n")
	}

	// and finally write to file!
	c.writer.Write(buffer.String())

	return nil
}
