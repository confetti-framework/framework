package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_map_with_nil_value(t *testing.T) {
	value := support.NewValue(nil)
	err := rule.Map{}.Verify(value)
	require.EqualError(t, err, "the :attribute must be a map")
}

func Test_map_with_map_value(t *testing.T) {
	value := support.NewValue(map[int]int{1: 1, 2: 1})
	err := rule.Map{}.Verify(value)
	require.Nil(t, err)
}
