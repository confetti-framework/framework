package rule

import (
	"github.com/confetti-framework/framework/support"
	"reflect"
)

type Min struct {
	Len int
}

func (m Min) Verify(value support.Value) error {
	input := value.Raw()
	amount := getAmount(input)
	if amount < m.Len {
		return errorWithExpectInput(m.getError(input), m.Len, amount)
	}
	return nil
}

func (m Min) getError(input interface{}) error {
	switch support.Kind(input) {
	case reflect.Slice, reflect.Map, reflect.Array:
		return MustBeAtLeastThanItemsError
	default:
		return MustBeAtLeastThanError
	}
}
