package srtfix_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/leotse/srtfix"
)

const OneLiner = `
1
00:00:00,030 --> 00:00:04,380
hello world
`

const MultiLiner = `


1
00:00:00,030 --> 00:00:04,380
hello 111

2
00:00:01,890 --> 00:00:07,200
hello 222

3
00:00:04,380 --> 00:00:10,170
hello 333


`

var _ = Describe("Parser", func() {
	Describe("ParseSrtFile()", func() {
		It("returns SrtFormatErr for invalid srt", func() {
			_, err := srtfix.ParseSrtFile("i am an invalid srt")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(srtfix.SrtFormatErr))
		})
		It("returns empty captions for empty srt", func() {
			captions, err := srtfix.ParseSrtFile("")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(captions).Should(HaveLen(0))
		})
		It("returns parsed captions", func() {
			captions, err := srtfix.ParseSrtFile(MultiLiner)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(captions).Should(HaveLen(3))
			Ω(captions[0]).Should(Equal(&srtfix.Caption{
				ID:    1,
				Start: 30 * time.Millisecond,
				End:   4*time.Second + 380*time.Millisecond,
				Text:  "hello 111",
			}))
			Ω(captions[1]).Should(Equal(&srtfix.Caption{
				ID:    2,
				Start: 1*time.Second + 890*time.Millisecond,
				End:   7*time.Second + 200*time.Millisecond,
				Text:  "hello 222",
			}))
			Ω(captions[2]).Should(Equal(&srtfix.Caption{
				ID:    3,
				Start: 4*time.Second + 380*time.Millisecond,
				End:   10*time.Second + 170*time.Millisecond,
				Text:  "hello 333",
			}))
		})
	})
	Describe("ParseCaption()", func() {
		It("returns CaptionFormatErr for invalid caption", func() {
			_, err := srtfix.ParseCaption("invalid caption")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(srtfix.CaptionFormatErr))
		})
		It("returns caption", func() {
			caption, err := srtfix.ParseCaption(OneLiner)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(caption).Should(Equal(&srtfix.Caption{
				ID:    1,
				Start: 30 * time.Millisecond,
				End:   4*time.Second + 380*time.Millisecond,
				Text:  "hello world",
			}))
		})
	})
	Describe("ParseTime()", func() {
		It("returns TimeFormatErr for invalid time", func() {
			_, err := srtfix.ParseTime("invalid time")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(srtfix.TimeFormatErr))
		})
		It("returns TimeFormatErr for empty time", func() {
			_, err := srtfix.ParseTime("")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(srtfix.TimeFormatErr))
		})
		It("returns TimeFormatErr for invalid time 00:00,390", func() {
			_, err := srtfix.ParseTime("00:00,390")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(srtfix.TimeFormatErr))
		})
		It("returns TimeFormatErr for invalid time 00:00:ab,090", func() {
			_, err := srtfix.ParseTime("00:00:ab,090")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(srtfix.TimeFormatErr))
		})
		It("returns duration for a valid time", func() {
			duration, err := srtfix.ParseTime("00:00:04,380")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(duration).Should(Equal(time.Duration(time.Second*4 + time.Millisecond*380)))
		})
	})
})
