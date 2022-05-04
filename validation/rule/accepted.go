package rule

import (
	"github.com/confetti-framework/framework/contract/inter"
	"github.com/confetti-framework/framework/support"
)

type Accepted struct{}

func (a Accepted) Requirements() []inter.Rule {
	return []inter.Rule{
		Present{},
	}
}

func (a Accepted) Verify(value support.Value) error {
	if !value.Bool() {
		return MustBeAcceptedError
	}
	return nil
}
