package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/confetti-framework/framework/validation/val"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_filled_and_present(t *testing.T) {
	errors := val.Validate(nil,
		support.NewValue(map[string]string{"age": "60"}),
		val.Verify("age", rule.Filled{}),
	)
	require.Empty(t, errors)
}

func Test_filled_and_not_present(t *testing.T) {
	errors := val.Validate(nil,
		support.NewValue(nil),
		val.Verify("name", rule.Filled{}),
	)
	require.Empty(t, errors)
}

func Test_filled_and_present_but_empty(t *testing.T) {
	errors := val.Validate(nil,
		support.NewValue(map[string]string{"age": ""}),
		val.Verify("age", rule.Filled{}),
	)
	require.EqualError(t, errors[0], "the age field must have a value")
}

func Test_filled_and_present_but_empty_slice(t *testing.T) {
	errors := val.Validate(nil,
		support.NewValue(map[string]interface{}{"names": []string{}}),
		val.Verify("names", rule.Filled{}),
	)
	require.EqualError(t, errors[0], "the names field must have a value")
}
