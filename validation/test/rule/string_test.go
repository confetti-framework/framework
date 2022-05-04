package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_string_with_nil_value(t *testing.T) {
	value := support.NewValue(nil)
	err := rule.String{}.Verify(value)
	require.EqualError(t, err, "the :attribute must be a string")
}

func Test_string_with_string_value(t *testing.T) {
	value := support.NewValue("dog")
	err := rule.String{}.Verify(value)
	require.Nil(t, err)
}
