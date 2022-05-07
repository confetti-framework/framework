package encode

import (
	"github.com/confetti-framework/framework/foundation/encoder"
	"github.com/confetti-framework/framework/foundation/test/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_json_response_from_bool_true(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, true, mock.JsonEncoders)

	require.Equal(t, "true", result)
}

func Test_json_response_from_bool_false(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, false, mock.JsonEncoders)

	require.Equal(t, "false", result)
}
