package value

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var theErr = errors.New("the error")

func Test_error_to_string(t *testing.T) {
	value, err := support.NewValue("valid_string", theErr).StringE()
	assert.Equal(t, "", value)
	assert.Equal(t, theErr, err)
}

func Test_error_to_int(t *testing.T) {
	value, err := support.NewValue(1, theErr).IntE()
	assert.Equal(t, 0, value)
	assert.ErrorIs(t, err, theErr)
}

func Test_error_to_float(t *testing.T) {
	value, err := support.NewValue(1., theErr).FloatE()
	assert.Equal(t, 0., value)
	assert.ErrorIs(t, err, theErr)
}

func Test_error_to_map(t *testing.T) {
	value, err := support.NewValue(1., theErr).MapE()
	assert.Equal(t, support.Map{}, value)
	assert.ErrorIs(t, err, theErr)
}

func Test_error_to_collection(t *testing.T) {
	require.Panics(t, func() {
		support.NewValue([]string{}, theErr).Collection()
	})
}
