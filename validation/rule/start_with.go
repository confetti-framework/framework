package rule

import (
	"github.com/confetti-framework/framework/support"
	"strings"
)

type Start struct {
	with []string
}

func (e Start) With(with ...string) Start {
	e.with = with
	return e
}

func (e Start) Verify(value support.Value) error {
	if len(e.with) == 0 {
		return OptionWithIsRequiredError
	}
	for _, expect := range e.with {
		if strings.HasPrefix(value.String(), expect) {
			return nil
		}
	}
	return errorWithExpectInput(MuseStartWithError, strings.Join(e.with, " or "), value.String())
}
