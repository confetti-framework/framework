package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_size_without_parameters_and_nil_value(t *testing.T) {
	value := support.NewValue(nil)
	err := rule.Size{}.Verify(value)
	require.Nil(t, err)
}

func Test_size_number_to_high(t *testing.T) {
	value := support.NewValue(6)
	err := rule.Size{Len: 5}.Verify(value)
	require.NotNil(t, err)
	require.Equal(t, "the :attribute must be 5, 6 given", err.Error())
}

func Test_size_number_lower(t *testing.T) {
	value := support.NewValue(4)
	err := rule.Size{Len: 5}.Verify(value)
	require.NotNil(t, err)
	require.Equal(t, "the :attribute must be 5, 4 given", err.Error())
}

func Test_size_number_equal(t *testing.T) {
	value := support.NewValue(6)
	err := rule.Size{Len: 6}.Verify(value)
	require.Nil(t, err)
}

func Test_size_slice_to_high(t *testing.T) {
	value := support.NewValue([]int{1, 2, 3, 4, 5, 6})
	err := rule.Size{Len: 2}.Verify(value)
	require.NotNil(t, err)
	require.Equal(t, "the :attribute must contain 2 items, 6 items given", err.Error())
}

func Test_size_slice_equal(t *testing.T) {
	value := support.NewValue([]int{1, 2, 3})
	err := rule.Size{Len: 3}.Verify(value)
	require.Nil(t, err)
}

func Test_size_map_to_high(t *testing.T) {
	value := support.NewValue(map[int]string{1: "_", 2: "_", 3: "_", 4: "_", 5: "_", 6: "_"})
	err := rule.Size{Len: 2}.Verify(value)
	require.NotNil(t, err)
	require.Equal(t, "the :attribute must contain 2 items, 6 items given", err.Error())
}

func Test_size_float_to_high(t *testing.T) {
	value := support.NewValue(6.7)
	err := rule.Size{Len: 2}.Verify(value)
	require.NotNil(t, err)
	require.Equal(t, "the :attribute must be 2, 6 given", err.Error())
}
