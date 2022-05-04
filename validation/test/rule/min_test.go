package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_min_without_parameters_and_nil_value(t *testing.T) {
	value := support.NewValue(nil)
	err := rule.Min{}.Verify(value)
	require.Nil(t, err)
}

func Test_min_number_to_high(t *testing.T) {
	value := support.NewValue(6)
	err := rule.Min{Len: 5}.Verify(value)
	require.Nil(t, err)
}

func Test_min_number_lower(t *testing.T) {
	value := support.NewValue(4)
	err := rule.Min{Len: 5}.Verify(value)
	require.NotNil(t, err)
	require.Equal(t, "the :attribute must be at least 5, 4 given", err.Error())
}

func Test_min_number_equal(t *testing.T) {
	value := support.NewValue(6)
	err := rule.Min{Len: 6}.Verify(value)
	require.Nil(t, err)
}

func Test_min_slice_to_low(t *testing.T) {
	value := support.NewValue([]int{1})
	err := rule.Min{Len: 2}.Verify(value)
	require.NotNil(t, err)
	require.Equal(t, "the :attribute must be at least 2 items, 1 items given", err.Error())
}

func Test_min_slice_equal(t *testing.T) {
	value := support.NewValue([]int{1, 2, 3})
	err := rule.Min{Len: 3}.Verify(value)
	require.Nil(t, err)
}

func Test_min_map_to_low(t *testing.T) {
	value := support.NewValue(map[int]string{1: "_"})
	err := rule.Min{Len: 2}.Verify(value)
	require.NotNil(t, err)
	require.Equal(t, "the :attribute must be at least 2 items, 1 items given", err.Error())
}

func Test_min_float_to_low(t *testing.T) {
	value := support.NewValue(1.3)
	err := rule.Min{Len: 2}.Verify(value)
	require.NotNil(t, err)
	require.Equal(t, "the :attribute must be at least 2, 1 given", err.Error())
}
