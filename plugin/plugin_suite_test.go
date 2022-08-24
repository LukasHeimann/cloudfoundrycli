package plugin_test

import (
	"path/filepath"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/pluginbuilder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	pluginbuilder.BuildTestBinary(filepath.Join("..", "fixtures", "plugins"), "test_1")
	RunSpecs(t, "Plugin Suite")
}
