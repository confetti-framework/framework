package rule

import (
	"github.com/confetti-framework/framework/support"
)

type IntegerAble struct{}

func (i IntegerAble) Verify(value support.Value) error {
	switch value.Raw().(type) {
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		_, err := value.IntE()
		if err == nil {
			return nil
		}
	}

	return MustBeAnIntegerError
}
