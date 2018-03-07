package srtfix_test

import (
	"github.com/golang/mock/gomock"
	"time"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"

	. "github.com/leotse/srtfix"
	. "github.com/leotse/srtfix/mocks"
)

var _ = Describe("Converter", func() {

	var ctrl *gomock.Controller
	var reader *MockFileReader
	var writer *MockFileWriter
	var parser *MockParser
	var resolver *MockResolver
	var converter *FileConverter

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		reader = NewMockFileReader(ctrl)
		writer = NewMockFileWriter(ctrl)
		parser = NewMockParser(ctrl)
		resolver = NewMockResolver(ctrl)
		converter = NewFileConverter(reader, writer, parser, resolver)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Convert()", func() {
		It("performs reader.Read() -> resolver.Resolve() -> writer.Write()", func() {
			reader.EXPECT().Read().Return(SrtContent, nil).Times(1)
			parser.EXPECT().Parse(SrtContent).Return(ParsedCaptions, nil).Times(1)
			resolver.EXPECT().Resolve(ParsedCaptions).Return(ParsedCaptions).Times(1)
			writer.EXPECT().Write(SrtResult).Times(1)
			converter.Convert()
		})
	})
})

const SrtContent = `1
00:00:00,030 --> 00:00:04,380
hello world
`

const SrtResult = `1
00:00:00,030 --> 00:00:04,380
hello world
`

var ParsedCaptions = []*Caption{
	&Caption{
		ID:        1,
		Start:     Time(30 * time.Millisecond),
		End:       Time(4*time.Second + 380*time.Millisecond),
		Text:      "hello world",
		StartText: "00:00:00,030",
		EndText:   "00:00:04,380",
	},
}
