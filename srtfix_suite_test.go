package srtfix_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSrtfix(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Srtfix Suite")
}
