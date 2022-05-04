package test

import (
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_raw_from_empty_string(t *testing.T) {
	require.Equal(t, "", support.NewValue("").Raw())
}

func Test_raw_from_nil_string(t *testing.T) {
	require.Equal(t, nil, support.NewValue(nil).Raw())
}

func Test_raw_from_string(t *testing.T) {
	require.Equal(t, "flour", support.NewValue("flour").Raw())
}

func Test_raw_from_bool(t *testing.T) {
	require.Equal(t, true, support.NewValue(true).Raw())
	require.Equal(t, false, support.NewValue(false).Raw())
}

func Test_raw_from_number(t *testing.T) {
	require.Equal(t, 100, support.NewValue(100).Raw())
	require.Equal(t, -100, support.NewValue(-100).Raw())
}

func Test_raw_from_float(t *testing.T) {
	require.Equal(t, 0.1, support.NewValue(0.1).Raw())
}

func Test_raw_from_collection_with_one_string(t *testing.T) {
	require.Equal(t, []interface{}{"door"}, support.NewCollection("door").Raw())
}

func Test_raw_from_collection_with_tho_strings(t *testing.T) {
	require.Equal(t, []interface{}{"foo", "bar"}, support.NewCollection("foo", "bar").Raw())
}

func Test_raw_from_collection_with_tho_numbers(t *testing.T) {
	require.Equal(t, []interface{}{12, 14}, support.NewCollection(12, 14).Raw())
}

func Test_raw_from_collection_with_tho_float(t *testing.T) {
	require.Equal(t, []interface{}{1.5, 0.4}, support.NewCollection(1.5, 0.4).Raw())
}

func Test_raw_from_value_with_collection(t *testing.T) {
	actual := support.NewValue(support.NewCollection("door")).Raw()
	require.Equal(t, []interface{}{"door"}, actual)
}

func Test_raw_from_map_with_strings(t *testing.T) {
	actual := support.NewMap(map[string]string{
		"chair": "blue",
		"table": "green",
	}).Raw()

	require.Equal(t, map[string]interface{}{"chair": "blue", "table": "green"}, actual)
}

func Test_raw_from_map_with_number_as_key(t *testing.T) {
	actual := support.NewMap(map[int]string{
		1: "blue",
		2: "green",
	}).Raw()

	require.Equal(t, map[string]interface{}{"1": "blue", "2": "green"}, actual)
}

func Test_raw_from_map_with_unknown_value_as_key(t *testing.T) {
	_, err := support.NewMapE(map[interface{}]string{
		testStruct{}: "blue",
	})

	require.EqualError(t, err, "invalid key in map: unable to cast test.testStruct{} of type test.testStruct to string")
}

func Test_raw_from_value_with_collection_and_map(t *testing.T) {
	actual := support.NewValue(
		support.NewCollection(
			support.NewMap(map[string]string{"key": "door"}),
		),
	).Raw()

	require.Equal(t, []interface{}{map[string]interface{}{"key": "door"}}, actual)
}

func Test_raw_from_value(t *testing.T) {
	raw := support.NewValue(100).Raw()
	require.Equal(t, 100, raw)
}

func Test_raw_collection_with_int_in_slice(t *testing.T) {
	value := support.NewCollection([]int{2})
	require.Equal(t, []interface{}{2}, value.Raw())
}
