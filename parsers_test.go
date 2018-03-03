package srtfix_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	. "github.com/leotse/srtfix"
)

var _ = Describe("Parser", func() {
	Describe("ParseSrtFile()", func() {
		It("returns SrtFormatErr for invalid srt", func() {
			_, err := ParseSrtFile("i am an invalid srt")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(SrtFormatErr))
		})
		It("returns empty captions for empty srt", func() {
			captions, err := ParseSrtFile("")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(captions).Should(HaveLen(0))
		})
		It("returns parsed captions", func() {
			captions, err := ParseSrtFile(MultiLiner)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(captions).Should(HaveLen(3))
			Ω(captions[0]).Should(Equal(&Caption{
				ID:        1,
				Start:     Time(30 * time.Millisecond),
				End:       Time(4*time.Second + 380*time.Millisecond),
				Text:      "hello 111",
				StartText: "00:00:00,030",
				EndText:   "00:00:04,380",
			}))
			Ω(captions[1]).Should(Equal(&Caption{
				ID:        2,
				Start:     Time(1*time.Second + 890*time.Millisecond),
				End:       Time(7*time.Second + 200*time.Millisecond),
				Text:      "hello 222",
				StartText: "00:00:01,890",
				EndText:   "00:00:07,200",
			}))
			Ω(captions[2]).Should(Equal(&Caption{
				ID:        3,
				Start:     Time(4*time.Second + 380*time.Millisecond),
				End:       Time(10*time.Second + 170*time.Millisecond),
				Text:      "hello 333",
				StartText: "00:00:04,380",
				EndText:   "00:00:10,170",
			}))
		})
	})
	Describe("ParseCaption()", func() {
		It("returns CaptionFormatErr for invalid caption", func() {
			_, err := ParseCaption("invalid caption")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(CaptionFormatErr))
		})
		It("returns caption", func() {
			caption, err := ParseCaption(OneLiner)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(caption).Should(Equal(&Caption{
				ID:        1,
				Start:     Time(30 * time.Millisecond),
				End:       Time(4*time.Second + 380*time.Millisecond),
				Text:      "hello world",
				StartText: "00:00:00,030",
				EndText:   "00:00:04,380",
			}))
		})
	})
	Describe("ParseTime()", func() {
		It("returns TimeFormatErr for invalid time", func() {
			_, err := ParseTime("invalid time")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(TimeFormatErr))
		})
		It("returns TimeFormatErr for empty time", func() {
			_, err := ParseTime("")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(TimeFormatErr))
		})
		It("returns TimeFormatErr for invalid time 00:00,390", func() {
			_, err := ParseTime("00:00,390")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(TimeFormatErr))
		})
		It("returns TimeFormatErr for invalid time 00:00:ab,090", func() {
			_, err := ParseTime("00:00:ab,090")
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(MatchError(TimeFormatErr))
		})
		It("returns duration for a valid time", func() {
			duration, err := ParseTime("00:00:04,380")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(duration).Should(Equal(Time(time.Second*4 + time.Millisecond*380)))
		})
	})
})

///////////////
// Test Data //
///////////////

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
