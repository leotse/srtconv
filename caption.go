package srtfix

import (
	"fmt"
	"time"
)

// Caption is a piece of text that displays at the
// specify time range in a video
type Caption struct {
	ID    int
	Start time.Duration
	End   time.Duration
	Text  string
}

// Merge the given caption's text with me, while keeping my
// ID, Start time and End time
func (c *Caption) Merge(with *Caption) {
	c.Text = c.Text + "\n" + with.Text
}

// String is the *Caption Stringer
func (c *Caption) String() string {
	return fmt.Sprintf("%v %v %v %v\n", c.ID, c.Start, c.End, c.Text)
}

// copy the given caption and return a new instance
func copy(c *Caption) *Caption {
	return &Caption{
		ID:    c.ID,
		Start: c.Start,
		End:   c.End,
		Text:  c.Text,
	}
}
