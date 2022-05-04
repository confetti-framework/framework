package value

import (
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_bool_true_from_true(t *testing.T) {
	require.True(t, support.NewValue(true).Bool())
}

func Test_bool_true_from_false(t *testing.T) {
	require.False(t, support.NewValue(false).Bool())
}

func Test_bool_true_from_int_one(t *testing.T) {
	require.True(t, support.NewValue(1).Bool())
}

func Test_bool_true_from_int_two(t *testing.T) {
	require.False(t, support.NewValue(2).Bool())
}

func Test_bool_true_from_int_zero(t *testing.T) {
	require.False(t, support.NewValue(0).Bool())
}

func Test_bool_true_from_string_one(t *testing.T) {
	require.True(t, support.NewValue("1").Bool())
}

func Test_bool_true_from_string_true(t *testing.T) {
	require.True(t, support.NewValue("true").Bool())
}

func Test_bool_true_from_string_on(t *testing.T) {
	require.True(t, support.NewValue("on").Bool())
}

func Test_bool_true_from_string_yes(t *testing.T) {
	require.True(t, support.NewValue("yes").Bool())
}

func Test_bool_with_false(t *testing.T) {
	value := support.NewValue(false).Bool()

	require.False(t, value)
}

func Test_bool_with_true(t *testing.T) {
	value := support.NewValue(true).Bool()

	require.True(t, value)
}
