package flag

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/types"
)

type Buildpack struct {
	types.FilteredString
}

func (b *Buildpack) UnmarshalFlag(val string) error {
	b.ParseValue(val)
	return nil
}
