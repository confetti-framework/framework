package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/spf13/cast"
	"reflect"
)

type Size struct {
	Len int
}

func (s Size) Verify(value support.Value) error {
	input := value.Raw()
	amount := getAmount(input)
	if amount != s.Len {
		return errorWithExpectInput(s.getError(input), s.Len, amount)
	}
	return nil
}

func getAmount(input interface{}) int {
	switch support.Kind(input) {
	case reflect.Slice, reflect.Map, reflect.Array:
		return reflect.ValueOf(input).Len()
	default:
		return cast.ToInt(input)
	}
}

func (s Size) getError(input interface{}) error {
	switch support.Kind(input) {
	case reflect.Slice, reflect.Map, reflect.Array:
		return MustBeContainItemsError
	default:
		return MustBeError
	}
}
