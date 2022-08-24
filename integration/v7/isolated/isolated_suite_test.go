package isolated

import (
	"testing"

	"github.com/LukasHeimann/cloudfoundrycli/v8/integration/helpers/commonisolated"
)

const (
	RealIsolationSegment = commonisolated.RealIsolationSegment
	DockerImage          = commonisolated.DockerImage
)

var (
	// Suite Level
	apiURL            string
	skipSSLValidation bool
	ReadOnlyOrg       string
	ReadOnlySpace     string

	// Per test
	homeDir string
)

func TestIsolated(t *testing.T) {
	commonisolated.CommonTestIsolated(t)
}

var _ = commonisolated.CommonGinkgoSetup(
	"summary_isi.txt",
	&apiURL,
	&skipSSLValidation,
	&ReadOnlyOrg,
	&ReadOnlySpace,
	&homeDir,
)
