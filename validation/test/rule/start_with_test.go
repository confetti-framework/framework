package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_start_with_without_comparison(t *testing.T) {
	value := support.NewValue(nil)
	err := rule.Start{}.Verify(value)
	require.EqualError(t, err, "option With is required")
}

func Test_start_with_one_comparison(t *testing.T) {
	value := support.NewValue("tiger")
	err := rule.Start{}.With("tiger").Verify(value)
	require.Nil(t, err)
}

func Test_start_with_invalid_comparison(t *testing.T) {
	value := support.NewValue("tiger")
	err := rule.Start{}.With("dog").Verify(value)
	require.NotNil(t, err)
	require.Equal(t, "the :attribute must start with dog, tiger given", err.Error())
}

func Test_start_with_with_multiple_comparisons_and_one_matched(t *testing.T) {
	value := support.NewValue("tiger")
	err := rule.Start{}.With("dog", "tiger").Verify(value)
	require.Nil(t, err)
}

func Test_start_with_with_multiple_comparisons_and_non_matched(t *testing.T) {
	value := support.NewValue("dog")
	err := rule.Start{}.With("snake", "tiger").Verify(value)
	require.NotNil(t, err)
	require.Equal(t, "the :attribute must start with snake or tiger, dog given", err.Error())
}

func Test_start_with_one_suffix_comparison(t *testing.T) {
	value := support.NewValue("owl snowy")
	err := rule.Start{}.With("owl").Verify(value)
	require.Nil(t, err)
}
