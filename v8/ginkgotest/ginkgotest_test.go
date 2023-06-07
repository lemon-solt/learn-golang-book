package ginkgotest_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ginkgotest", func() {
	Context("ginkgotest context", func ()  {
		It("数値をテストする", func(){
			Expect(0).To(Equal(0))
		})
	})
})
