package flag_test

import (
	. "github.com/LukasHeimann/cloudfoundrycli/v8/command/flag"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TrimmedString", func() {
	var trimmedString TrimmedString

	Describe("UnmarshalFlag", func() {
		BeforeEach(func() {
			err := trimmedString.UnmarshalFlag("   some string   ")
			Expect(err).NotTo(HaveOccurred())
		})

		It("has the right value", func() {
			Expect(trimmedString).To(BeEquivalentTo("some string"))
		})
	})
})
