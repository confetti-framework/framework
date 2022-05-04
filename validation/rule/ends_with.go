package rule

import (
	"github.com/confetti-framework/framework/support"
	"strings"
)

type Ends struct {
	with []string
}

func (e Ends) With(with ...string) Ends {
	e.with = with
	return e
}

func (e Ends) Verify(value support.Value) error {
	if len(e.with) == 0 {
		return OptionWithIsRequiredError
	}
	for _, expect := range e.with {
		if strings.HasSuffix(value.String(), expect) {
			return nil
		}
	}
	return errorWithExpectInput(MuseEndWithError, strings.Join(e.with, " or "), value.String())
}
