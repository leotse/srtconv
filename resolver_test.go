package srtfix_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	. "github.com/leotse/srtfix"
)

var _ = Describe("Fixer", func() {
	Describe("Resolve()", func() {
		It("returns empty captinos for nil input", func() {
			resolved, err := Resolve(nil)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resolved).Should(HaveLen(0))
		})
		It("returns empty captions for empty input", func() {
			resolved, err := Resolve(Empty)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resolved).Should(HaveLen(0))
		})
		It("does not change a single caption", func() {
			resolved, err := Resolve(OneCaption)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resolved).Should(HaveLen(1))
			Ω(resolved).Should(Equal(OneCaption))
		})
		It("does not change valid captions", func() {
			resolved, err := Resolve(NoOverlap)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resolved).Should(HaveLen(2))
			Ω(resolved).Should(Equal(NoOverlap))
		})
		It("merges 2 overlapping captions", func() {
			resolved, err := Resolve(TwoOverlap)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resolved).Should(HaveLen(1))
			Ω(resolved[0]).Should(Equal(&Caption{
				ID:    1,
				Start: TwoOverlap[0].Start,
				End:   TwoOverlap[1].End,
				Text:  TwoOverlap[0].Text + "\n" + TwoOverlap[1].Text,
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
		Start: 50 * time.Millisecond,
		End:   3 * time.Second,
		Text:  "hello",
	},
}
var NoOverlap = []*Caption{
	&Caption{
		ID:    1,
		Start: 10 * time.Millisecond,
		End:   2 * time.Second,
		Text:  "abc",
	},
	&Caption{
		ID:    2,
		Start: 2 * time.Second,
		End:   3 * time.Second,
		Text:  "def",
	},
}
var TwoOverlap = []*Caption{
	&Caption{
		ID:    2,
		Start: 50 * time.Millisecond,
		End:   3 * time.Second,
		Text:  "hello 1",
	},
	&Caption{
		ID:    3,
		Start: 2 * time.Second,
		End:   4 * time.Second,
		Text:  "hello 2",
	},
}
