package main

import (
	"github.com/Kalimaha/ginkgo/reporter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestLambda(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithCustomReporters(t, "Test Lambda", []Reporter{reporter.New()})
}

var _ = Describe("Calculator", func() {
	It("sums two numbers", func() {
		Expect(Sum(2, 3)).To(Equal(5))
	})
})
