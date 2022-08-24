package environmentvariablegroups_test

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/i18n"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/configuration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEnvironmentVariableGroups(t *testing.T) {
	config := configuration.NewRepositoryWithDefaults()
	i18n.T = i18n.Init(config)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Environment Variable Groups Suite")
}
