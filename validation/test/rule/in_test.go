package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_in_with_no_options(t *testing.T) {
	err := rule.In{}.Verify(support.NewValue(nil))
	require.EqualError(t, err, "option With is required")
}

func Test_with_one_option(t *testing.T) {
	value := support.NewValue("tiger")
	err := rule.In{}.With("tiger").Verify(value)
	require.Nil(t, err)
}

func Test_with_two_options(t *testing.T) {
	value := support.NewValue("tiger")
	err := rule.In{}.With("tiger", "owl").Verify(value)
	require.Nil(t, err)
}

func Test_with_an_integer(t *testing.T) {
	value := support.NewValue(44)
	err := rule.In{}.With(44).Verify(value)
	require.Nil(t, err)
}

func Test_with_invalid_value(t *testing.T) {
	value := support.NewValue("salamander")
	err := rule.In{}.With("bison").Verify(value)
	require.EqualError(t, err, "the selected :attribute is invalid")
}
