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
