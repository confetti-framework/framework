package val_errors

import (
	"fmt"
	"github.com/confetti-framework/errors"
	"strings"
)

func FindError(errs []error, attribute string) error {
	for _, err := range errs {
		errAttribute, _ := FindByAttribute(err)
		if attribute == errAttribute {
			return err
		}
	}
	return nil
}

func FindByAttribute(err error) (string, bool) {
	var fieldHolder *withAttribute

	ok := errors.As(err, &fieldHolder)
	if !ok {
		return "", false
	}

	message, ok := fieldHolder.attributes["attribute"]
	return message, ok
}

func WithAttribute(err error, key string, attribute string) error {
	if err == nil {
		return nil
	}
	return &withAttribute{err, map[string]string{key: attribute}}
}

func WithAttributes(err error, attributes map[string]string) error {
	if err == nil {
		return nil
	}
	return &withAttribute{err, attributes}
}

type withAttribute struct {
	cause      error
	attributes map[string]string
}

func (w *withAttribute) Error() string {
	err := w.cause.Error()
	for placeholder, attribute := range w.attributes {
		placeholder = ":" + strings.TrimLeft(placeholder, ":")
		err = strings.Replace(err, placeholder, attribute, 1)
	}
	return err
}

func (w *withAttribute) Format(st fmt.State, verb rune) {
	errors.Format(st, verb, w.cause)
}

func (w *withAttribute) Unwrap() error {
	return w.cause
}
