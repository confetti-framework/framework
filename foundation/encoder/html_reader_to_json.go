package encoder

import (
	"fmt"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/framework/contract/inter"
	"reflect"
)

type StringerToHtml struct{}

func (h StringerToHtml) IsAble(object interface{}) bool {
	_, ok := object.(fmt.Stringer)
	return ok
}

func (h StringerToHtml) EncodeThrough(_ inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	result, ok := object.(fmt.Stringer)
	if !ok {
		return "", errors.Wrap(EncodeError, "can not encode to html with an unsupported type "+reflect.TypeOf(object).String())
	}
	return result.String(), nil
}
