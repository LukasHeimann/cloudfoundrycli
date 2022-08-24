package main_test

import (
	"time"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/pluginbuilder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTestRpcServerExample(t *testing.T) {
	RegisterFailHandler(Fail)

	pluginbuilder.BuildTestBinary("", "test_rpc_server_example")

	RunSpecs(t, "Test RPC Server Example Suite")
}

var _ = BeforeEach(func() {
	SetDefaultEventuallyTimeout(3 * time.Second)
})
