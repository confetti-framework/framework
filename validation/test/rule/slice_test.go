package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_slice_with_nil_value(t *testing.T) {
	value := support.NewValue(nil)
	err := rule.Slice{}.Verify(value)
	require.EqualError(t, err, "the :attribute must be a slice")
}

func Test_slice_with_slice_value(t *testing.T) {
	value := support.NewValue([]int{1, 2})
	err := rule.Slice{}.Verify(value)
	require.Nil(t, err)
}
