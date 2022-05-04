package value

import (
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_float_from_empty_string(t *testing.T) {
	value := support.NewValue("")

	result, err := value.FloatE()

	require.Equal(t, 0.0, result)
	require.Error(t, err, "unable to cast \"\" of type string to float64")
}

func Test_float_from_words(t *testing.T) {
	value := support.NewValue("four")

	result, err := value.FloatE()

	require.Equal(t, 0.0, result)
	require.EqualError(t, err, "unable to cast \"four\" of type string to float64")
}

func Test_float_from_long_number(t *testing.T) {
	value := support.NewValue(
		"12345678912345367891234523456567896123475123456789123453678912345234565678" +
			"9612347567912345678912345367891234567912344567891253456789612347567912" +
			"3456789123453678912345679123445678912534567896123475679123456789123453" +
			"6789123456791234456789125345678961234756791234567891234536789123456791" +
			"23445678912534567896123475679",
	)

	result, err := value.FloatE()

	require.Equal(t, 0.0, result)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "unable to cast \"1234")
	require.Contains(t, err.Error(), "475679\" of type string to float64")
}

func Test_float_from_string(t *testing.T) {
	value := support.NewValue("1.5123")

	result, err := value.FloatE()

	require.Equal(t, 1.5123, result)
	require.NoError(t, err)
}

func Test_float_from_different_int_types(t *testing.T) {
	var result float64

	result, _ = support.NewValue(312).FloatE()
	require.Equal(t, float64(312), result)

	result, _ = support.NewValue(int8(2)).FloatE()
	require.Equal(t, float64(2), result)

	result, _ = support.NewValue(int16(2)).FloatE()
	require.Equal(t, float64(2), result)
}

func Test_first_float_from_collection(t *testing.T) {
	result := support.NewValue(support.NewCollection(12.12)).Float()

	require.Equal(t, 12.12, result)
}

func Test_first_float_from_map(t *testing.T) {
	input := support.NewMap(map[string]interface{}{"total": 12.12})
	result := support.NewValue(input).Float()

	require.Equal(t, 12.12, result)
}

func Test_float_not_panic_without_error_receiver(t *testing.T) {
	require.NotPanics(t, func() {
		support.NewValue(123).Float()
	})
	require.Equal(t, float64(123), support.NewValue(123).Float())
}
