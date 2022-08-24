//go:build !windows
// +build !windows

package generic_test

import (
	"path/filepath"

	. "github.com/LukasHeimann/cloudfoundrycli/v8/util/generic"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExecutableFilename", func() {
	When("a filename which must be turned into an executable filename is input", func() {
		It("does nothing on unix", func() {
			myPath := filepath.Join("foo", "bar")
			Expect(ExecutableFilename(myPath)).To(Equal(myPath))
		})
	})
})
