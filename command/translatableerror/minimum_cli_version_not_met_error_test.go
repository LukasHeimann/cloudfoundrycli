package translatableerror_test

import (
	. "github.com/LukasHeimann/cloudfoundrycli/v8/command/translatableerror"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MinimumCLIVersionNotMetError", func() {
	Describe("Error", func() {
		When("BinaryVersion is not empty", func() {
			It("returns a template with the value of the BinaryVersion field", func() {
				err := MinimumCLIVersionNotMetError{
					BinaryVersion: "1.2.3",
				}

				Expect(err.Error()).To(Equal("Cloud Foundry API version {{.APIVersion}} requires CLI version {{.MinCLIVersion}}. You are currently on version {{.BinaryVersion}}. To upgrade your CLI, please visit: https://github.com/cloudfoundry/cli#downloads"))
			})
		})
	})
})
