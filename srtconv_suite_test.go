package srtconv_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSrtconv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Srtconv Suite")
}
