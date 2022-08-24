package v7action_test

import (
	. "github.com/LukasHeimann/cloudfoundrycli/v8/actor/v7action"
	"github.com/LukasHeimann/cloudfoundrycli/v8/actor/v7action/v7actionfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Version Check Actions", func() {
	var (
		actor         *Actor
		fakeUAAClient *v7actionfakes.FakeUAAClient
	)

	BeforeEach(func() {
		actor, _, _, _, fakeUAAClient, _, _ = NewTestActor()
	})

	Describe("GetUAAAPIVersion", func() {
		It("returns the UAA API version", func() {
			expectedVersion := "1.96.0"
			fakeUAAClient.GetAPIVersionReturns(expectedVersion, nil)
			retrievedVersion, err := actor.GetUAAAPIVersion()
			Expect(retrievedVersion).To(Equal(expectedVersion))
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
