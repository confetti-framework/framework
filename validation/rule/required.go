package rule

import (
	"github.com/confetti-framework/framework/inter"
	"github.com/confetti-framework/framework/support"
)

type Required struct{}

func (r Required) Requirements() []inter.Rule {
	return []inter.Rule{
		Present{},
	}
}

func (r Required) Verify(value support.Value) error {
	if !value.Filled() {
		return IsRequiredError
	}
	return nil
}
