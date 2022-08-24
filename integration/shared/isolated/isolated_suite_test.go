package isolated

import (
	"testing"

	"github.com/LukasHeimann/cloudfoundrycli/v8/integration/helpers/commonisolated"
)

var (
	// Suite Level
	apiURL            string
	skipSSLValidation bool
	ReadOnlyOrg       string
	ReadOnlySpace     string

	// Per Test Level
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
