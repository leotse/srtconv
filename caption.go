package srtfix

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Time handles the serializing of srt time format
type Time time.Duration

// Caption is a piece of text that displays at the
// specify time range in a video
type Caption struct {
	ID    int
	Start Time
	End   Time
	Text  string

	// the following fields are only to improve perf
	StartText string
	EndText   string
}

// Merge the given caption's text with me, while keeping my
// ID, Start time and End time
func (c *Caption) Merge(with *Caption) {
	c.Text = c.Text + "\n" + with.Text
}

// String is the *Caption Stringer
func (c *Caption) String() string {
	parts := []string{
		strconv.Itoa(c.ID),
		fmt.Sprintf("%v --> %v", c.StartText, c.EndText),
		c.Text,
	}
	return strings.Join(parts, "\n")
}

// copy the given caption and return a new instance
func copy(c *Caption) *Caption {
	return &Caption{
		ID:        c.ID,
		Start:     c.Start,
		End:       c.End,
		Text:      c.Text,
		StartText: c.StartText,
		EndText:   c.EndText,
	}
}
