package srtfix

import "time"

// Caption is a piece of text that displays at the
// specify time range in a video
type Caption struct {
	ID    int
	Start time.Duration
	End   time.Duration
	Text  string
}

// Merge the given caption with me, keepin my current ID and Start time.
func (c *Caption) Merge(with *Caption) {
	c.End = with.End
	c.Text = c.Text + "\n" + with.Text
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
