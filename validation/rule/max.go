package rule

import (
	"github.com/confetti-framework/framework/support"
	"reflect"
)

type Max struct {
	Len int
}

func (m Max) Verify(value support.Value) error {
	input := value.Raw()
	amount := getAmount(input)
	if amount > m.Len {
		return errorWithExpectInput(m.getError(input), m.Len, amount)
	}
	return nil
}

func (m Max) getError(input interface{}) error {
	switch support.Kind(input) {
	case reflect.Slice, reflect.Map, reflect.Array:
		return MayNotHaveMoreThanItemsError
	default:
		return MayNotBeGreaterThanError
	}
}
