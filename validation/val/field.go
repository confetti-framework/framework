package val

import (
	"github.com/confetti-framework/contract/inter"
)

type Verification struct {
	Field string
	Rules []inter.Rule
	app   inter.AppReader
}

func Verify(field string, rules ...inter.Rule) Verification {
	return Verification{Field: field, Rules: rules}
}
