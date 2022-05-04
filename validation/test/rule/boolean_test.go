package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_boolean_with_true(t *testing.T) {
	err := rule.Boolean{}.Verify(support.NewValue(true))
	require.Nil(t, err)
}

func Test_boolean_with_nil(t *testing.T) {
	err := rule.Boolean{}.Verify(support.NewValue(nil))
	require.EqualError(t, err, "the :attribute must be a boolean, nil given")
}

func Test_boolean_with_int_0(t *testing.T) {
	err := rule.Boolean{}.Verify(support.NewValue(0))
	require.Nil(t, err)
}

func Test_boolean_with_int_1(t *testing.T) {
	err := rule.Boolean{}.Verify(support.NewValue(1))
	require.Nil(t, err)
}

func Test_boolean_with_invalid_string(t *testing.T) {
	err := rule.Boolean{}.Verify(support.NewValue("qwerty"))
	require.EqualError(t, err, "the :attribute must be a boolean, qwerty given")
}

func Test_boolean_with_string_true(t *testing.T) {
	err := rule.Boolean{}.Verify(support.NewValue("true"))
	require.EqualError(t, err, "the :attribute must be a boolean, true given")
}

func Test_boolean_with_string_false(t *testing.T) {
	err := rule.Boolean{}.Verify(support.NewValue("false"))
	require.EqualError(t, err, "the :attribute must be a boolean, false given")
}

func Test_boolean_with_string_zero(t *testing.T) {
	err := rule.Boolean{}.Verify(support.NewValue("0"))
	require.Nil(t, err)
}

func Test_boolean_with_string_one(t *testing.T) {
	err := rule.Boolean{}.Verify(support.NewValue("1"))
	require.Nil(t, err)
}

func Test_boolean_with_camel_true(t *testing.T) {
	err := rule.Boolean{}.Verify(support.NewValue("True"))
	require.EqualError(t, err, "the :attribute must be a boolean, True given")
}
