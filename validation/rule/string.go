package rule

import (
	"github.com/confetti-framework/framework/support"
	"reflect"
)

type String struct{}

func (s String) Verify(value support.Value) error {
	kind := support.Kind(value.Raw())
	if kind != reflect.String {
		return MustBeAStringError
	}
	return nil
}
