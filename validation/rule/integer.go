package rule

import (
	"github.com/confetti-framework/framework/support"
)

type Integer struct{}

func (i Integer) Verify(value support.Value) error {
	switch value.Raw().(type) {
	case nil, bool, float32, float64:
		return MustBeAnIntegerError
	}

	_, err := value.IntE()
	if err != nil {
		return MustBeAnIntegerError
	}

	return nil
}
