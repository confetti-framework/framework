package test

import (
	"github.com/confetti-framework/support/str"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_upper_first_with_empty_string(t *testing.T) {
	result := str.UpperFirst("")
	require.Equal(t, "", result)
}

func Test_upper_first_with_multiple_words(t *testing.T) {
	result := str.UpperFirst("a horse is happy")
	require.Equal(t, "A horse is happy", result)
}

func Test_in_slice_with_no_parameter(t *testing.T) {
	require.False(t, str.InSlice("phone"))
}

func Test_in_slice_with_one_non_existing_string(t *testing.T) {
	require.False(t, str.InSlice("phone", "bag"))
}

func Test_in_slice_with_one_existing_string(t *testing.T) {
	require.True(t, str.InSlice("phone", "phone"))
}

func Test_in_slice_with_multiple_one_matched_parameters(t *testing.T) {
	require.True(t, str.InSlice("phone", "TV", "phone", "tabel"))
}

func Test_in_slice_with_integer(t *testing.T) {
	require.True(t, str.InSlice(1, 0, 1))
}
