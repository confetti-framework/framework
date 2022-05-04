package rule

import (
	"github.com/confetti-framework/framework/support"
	"reflect"
)

type Map struct{}

func (m Map) Verify(value support.Value) error {
	kind := support.Kind(value.Raw())
	if kind != reflect.Map {
		return MustBeAMapError
	}
	return nil
}
