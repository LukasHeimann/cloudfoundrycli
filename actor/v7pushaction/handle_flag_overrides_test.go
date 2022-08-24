package v7pushaction_test

import (
	. "github.com/LukasHeimann/cloudfoundrycli/v8/actor/v7pushaction"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/manifestparser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HandleFlagOverrides", func() {
	var (
		pushActor           *Actor
		baseManifest        manifestparser.Manifest
		flagOverrides       FlagOverrides
		transformedManifest manifestparser.Manifest
		executeErr          error

		testFuncCallCount int
	)

	testTransformManifestFunc := func(manifest manifestparser.Manifest, overrides FlagOverrides) (manifestparser.Manifest, error) {
		testFuncCallCount += 1
		return manifest, nil
	}

	BeforeEach(func() {
		baseManifest = manifestparser.Manifest{}
		flagOverrides = FlagOverrides{}
		testFuncCallCount = 0

		pushActor, _, _ = getTestPushActor()
		pushActor.TransformManifestSequence = []HandleFlagOverrideFunc{
			testTransformManifestFunc,
		}
	})

	JustBeforeEach(func() {
		transformedManifest, executeErr = pushActor.HandleFlagOverrides(baseManifest, flagOverrides)
	})

	It("calls each transform-manifest function", func() {
		Expect(testFuncCallCount).To(Equal(1))
		Expect(executeErr).NotTo(HaveOccurred())
		Expect(transformedManifest).To(Equal(baseManifest))
	})
})
