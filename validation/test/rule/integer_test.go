package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_integer_with_nil_value(t *testing.T) {
	value := support.NewValue(nil)
	err := rule.Integer{}.Verify(value)
	require.EqualError(t, err, "the :attribute must be an integer")
}

func Test_integer_with_bool_value(t *testing.T) {
	value := support.NewValue(true)
	err := rule.Integer{}.Verify(value)
	require.EqualError(t, err, "the :attribute must be an integer")
}

func Test_integer_with_float_value(t *testing.T) {
	value := support.NewValue(3.0)
	err := rule.Integer{}.Verify(value)
	require.EqualError(t, err, "the :attribute must be an integer")
}

func Test_integer_with_valid_integer_value(t *testing.T) {
	value := support.NewValue(2)
	err := rule.Integer{}.Verify(value)
	require.Nil(t, err)
}

func Test_integer_with_invalid_value(t *testing.T) {
	value := support.NewValue("two")
	err := rule.Integer{}.Verify(value)
	require.EqualError(t, err, "the :attribute must be an integer")
}

func Test_integer_with_valid_string_value(t *testing.T) {
	value := support.NewValue("2")
	err := rule.Integer{}.Verify(value)
	require.Nil(t, err)
}
