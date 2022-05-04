package rule

import (
	"github.com/confetti-framework/framework/support"
)

type In struct {
	with []interface{}
}

func (i In) Verify(value support.Value) error {
	if len(i.with) == 0 {
		return OptionWithIsRequiredError
	}
	for _, compare := range i.with {
		if value.Raw() == compare {
			return nil
		}
	}
	return SelectedIsInvalidError
}

func (i In) With(with ...interface{}) In {
	i.with = with
	return i
}
