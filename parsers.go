package srtfix

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// ParseSrtFile parses the given content into a slice of catpion
func ParseSrtFile(content string) ([]*Caption, error) {

	// empty content?
	if len(content) == 0 {
		return []*Caption{}, nil
	}

	// split srt content into parts
	sections := strings.Split(strings.Trim(content, " \n"), "\n\n")

	// parse each caption sections
	captions := make([]*Caption, len(sections))
	var err error
	for i, section := range sections {
		captions[i], err = ParseCaption(section)
		if err != nil {
			return nil, errors.New(SrtFormatErr)
		}
	}

	return captions, nil
}

// ParseCaption parses a single line of caption
func ParseCaption(c string) (*Caption, error) {

	// split caption section into 3 lines
	parts := strings.Split(strings.Trim(c, " \n"), "\n")
	if len(parts) != 3 {
		return nil, errors.New(CaptionFormatErr)
	}

	var err error
	caption := Caption{}

	// 1st line is the ID
	caption.ID, err = strconv.Atoi(parts[0])
	if err != nil {
		return nil, errors.New(CaptionFormatErr)
	}

	// 2nd line is the time
	timeParts := strings.Split(parts[1], " ")

	caption.Start, err = ParseTime(timeParts[0])
	if err != nil {
		return nil, errors.New(CaptionFormatErr)
	}

	caption.End, err = ParseTime(timeParts[len(timeParts)-1])
	if err != nil {
		return nil, errors.New(CaptionFormatErr)
	}

	// 3rd line is the text
	caption.Text = strings.Trim(parts[2], " ")

	return &caption, nil
}

// ParseTime parses a srt time into a Duration
// Expects input to be in the format: `00:00:04,380`
func ParseTime(t string) (time.Duration, error) {

	// split the time into 4 parts
	parts := strings.FieldsFunc(t, timeDelimiters)
	if len(parts) != 4 {
		return time.Duration(0), errors.New(TimeFormatErr)
	}

	// convert parts into duration
	// each part should be a valid number
	partDurations := []time.Duration{time.Hour, time.Minute, time.Second, time.Millisecond}
	total := time.Duration(0)
	for i, part := range parts {
		numPart, err := strconv.Atoi(part)
		if err != nil {
			return time.Duration(0), errors.New(TimeFormatErr)
		}
		total = total + partDurations[i]*time.Duration(numPart)
	}

	return time.Duration(total), nil
}

func timeDelimiters(c rune) bool {
	return c == rune(':') || c == rune(',')
}
