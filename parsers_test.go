package srtconv_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/leotse/srtconv"
)

const OneLiner = `
1
00:00:00,030 --> 00:00:04,380
hello world
`

var _ = Describe("Parser", func() {
	Describe("ParseSrtFile()", func() {
		It("returns SrtFormatErr for invalid srt", func() {
			_, err := srtconv.ParseSrtFile("i am an invalid srt")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(srtconv.SrtFormatErr))
		})
		It("returns empty captions for empty srt", func() {
			captions, err := srtconv.ParseSrtFile("")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(captions).Should(HaveLen(0))
		})
		It("returns one line of caption for one-line srt ", func() {
			captions, err := srtconv.ParseSrtFile(OneLiner)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(captions).Should(HaveLen(1))

			caption := captions[0]
			Ω(caption.ID).Should(Equal(1))
		})
	})
	Describe("ParseCaption()", func() {
		It("returns CaptionFormatErr for invalid caption", func() {
			_, err := srtconv.ParseCaption("invalid caption")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(srtconv.CaptionFormatErr))
		})
		It("returns caption", func() {
			caption, err := srtconv.ParseCaption(OneLiner)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(caption).Should(Equal(&srtconv.Caption{
				ID:    1,
				Start: 30 * time.Millisecond,
				End:   4*time.Second + 380*time.Millisecond,
				Text:  "hello world",
			}))
		})
	})
	Describe("ParseTime()", func() {
		It("returns TimeFormatErr for invalid time", func() {
			_, err := srtconv.ParseTime("invalid time")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(srtconv.TimeFormatErr))
		})
		It("returns TimeFormatErr for empty time", func() {
			_, err := srtconv.ParseTime("")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(srtconv.TimeFormatErr))
		})
		It("returns TimeFormatErr for invalid time 00:00,390", func() {
			_, err := srtconv.ParseTime("00:00,390")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(srtconv.TimeFormatErr))
		})
		It("returns TimeFormatErr for invalid time 00:00:ab,090", func() {
			_, err := srtconv.ParseTime("00:00:ab,090")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(srtconv.TimeFormatErr))
		})
		It("returns duration for a valid time", func() {
			duration, err := srtconv.ParseTime("00:00:04,380")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(duration).Should(Equal(time.Duration(time.Second*4 + time.Millisecond*380)))
		})
	})
})
