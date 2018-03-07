package srtfix_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	. "github.com/leotse/srtfix"
)

var _ = Describe("Resolver", func() {
	Describe("Resolve()", func() {
		It("returns empty captinos for nil input", func() {
			resolved := Resolve(nil)
			Ω(resolved).Should(HaveLen(0))
		})
		It("returns empty captions for empty input", func() {
			resolved := Resolve(Empty)
			Ω(resolved).Should(HaveLen(0))
		})
		It("does not change a single caption", func() {
			resolved := Resolve(OneCaption)
			Ω(resolved).Should(HaveLen(1))
			Ω(resolved).Should(Equal(OneCaption))
		})
		It("does not change valid captions", func() {
			resolved := Resolve(NoOverlap)
			Ω(resolved).Should(HaveLen(2))
			Ω(resolved).Should(Equal(NoOverlap))
		})
		It("merges 2 overlapping captions", func() {
			resolved := Resolve(TwoOverlap)
			Ω(resolved).Should(HaveLen(1))
			Ω(resolved[0]).Should(Equal(&Caption{
				ID:    1,
				Start: TwoOverlap[0].Start,
				End:   TwoOverlap[0].End,
				Text:  TwoOverlap[0].Text + "\n" + TwoOverlap[1].Text,
			}))
		})
		It("merges 3 overlapping captions", func() {
			resolved := Resolve(ThreeOverlap)
			Ω(resolved).Should(HaveLen(2))
			Ω(resolved[0]).Should(Equal(&Caption{
				ID:    1,
				Start: ThreeOverlap[0].Start,
				End:   ThreeOverlap[0].End,
				Text:  ThreeOverlap[0].Text + "\n" + ThreeOverlap[1].Text,
			}))
			Ω(resolved[1]).Should(Equal(&Caption{
				ID:    2,
				Start: ThreeOverlap[2].Start,
				End:   ThreeOverlap[2].End,
				Text:  ThreeOverlap[2].Text,
			}))
		})
	})
})

///////////////
// Test Data //
///////////////

var Empty = []*Caption{}
var OneCaption = []*Caption{
	&Caption{
		ID:    1,
		Start: Time(50 * time.Millisecond),
		End:   Time(3 * time.Second),
		Text:  "hello",
	},
}
var NoOverlap = []*Caption{
	&Caption{
		ID:    1,
		Start: Time(10 * time.Millisecond),
		End:   Time(2 * time.Second),
		Text:  "abc",
	},
	&Caption{
		ID:    2,
		Start: Time(2 * time.Second),
		End:   Time(3 * time.Second),
		Text:  "def",
	},
}
var TwoOverlap = []*Caption{
	&Caption{
		ID:    2,
		Start: Time(50 * time.Millisecond),
		End:   Time(3 * time.Second),
		Text:  "hello 1",
	},
	&Caption{
		ID:    3,
		Start: Time(2 * time.Second),
		End:   Time(4 * time.Second),
		Text:  "hello 2",
	},
}
var ThreeOverlap = []*Caption{
	&Caption{
		ID:    1,
		Start: Time(30 * time.Millisecond),
		End:   Time(4*time.Second + 380*time.Millisecond),
		Text:  "hello 1",
	},
	&Caption{
		ID:    2,
		Start: Time(1*time.Second + 890*time.Millisecond),
		End:   Time(7*time.Second + 200*time.Millisecond),
		Text:  "hello 2",
	},
	&Caption{
		ID:    3,
		Start: Time(4*time.Second + 380*time.Millisecond),
		End:   Time(10*time.Second + 170*time.Millisecond),
		Text:  "hello 3",
	},
}
