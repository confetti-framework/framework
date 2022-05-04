package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ends_with_without_comparison(t *testing.T) {
	value := support.NewValue(nil)
	err := rule.Ends{}.Verify(value)
	require.EqualError(t, err, "option With is required")
}

func Test_ends_with_one_comparison(t *testing.T) {
	value := support.NewValue("tiger")
	err := rule.Ends{}.With("tiger").Verify(value)
	require.Nil(t, err)
}

func Test_ends_with_invalid_comparison(t *testing.T) {
	value := support.NewValue("tiger")
	err := rule.Ends{}.With("dog").Verify(value)
	require.NotNil(t, err)
	require.Equal(t, "the :attribute must end with dog, tiger given", err.Error())
}

func Test_ends_with_with_multiple_comparisons_and_one_matched(t *testing.T) {
	value := support.NewValue("tiger")
	err := rule.Ends{}.With("dog", "tiger").Verify(value)
	require.Nil(t, err)
}

func Test_ends_with_with_multiple_comparisons_and_non_matched(t *testing.T) {
	value := support.NewValue("dog")
	err := rule.Ends{}.With("snake", "tiger").Verify(value)
	require.NotNil(t, err)
	require.Equal(t, "the :attribute must end with snake or tiger, dog given", err.Error())
}

func Test_ends_with_one_suffix_comparison(t *testing.T) {
	value := support.NewValue("snowy owl")
	err := rule.Ends{}.With("owl").Verify(value)
	require.Nil(t, err)
}
