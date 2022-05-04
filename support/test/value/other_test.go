package value

import (
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockEmptyStruct struct{}
type mockStruct struct {
	Field string
}

var mockFunc = func() {}

func Test_new_invalid_value(t *testing.T) {
	require.Panics(t, func() {
		support.NewValue(map[interface{}]string{
			mockFunc: "val",
		})
	})
}

func Test_get_collection_from_value_with_asterisks(t *testing.T) {
	value := support.NewValue(support.NewCollection([]string{"the_value"}))
	require.Equal(t, []interface{}{"the_value"}, value.Get("*").Raw())
}

func Test_get_map_from_value_with_asterisks(t *testing.T) {
	value := support.NewValue(support.NewMap(map[string]interface{}{"key": "the_value"}))
	require.Equal(t, map[string]interface{}{"key": "the_value"}, value.Get("*").Raw())
}

func Test_get_from_invalid_value_with_asterisks(t *testing.T) {
	value := support.NewValue("non collection/value")
	require.Panics(t, func() {
		value.Get("*").Raw()
	})
}

func Test_get_from_empty_struct(t *testing.T) {
	value := support.NewValue(mockEmptyStruct{})
	require.Panics(t, func() {
		value.Get("field").Raw()
	})
}

func Test_get_from_struct(t *testing.T) {
	value := support.NewValue(mockStruct{"fieldvalue"})
	require.Equal(t, "fieldvalue", value.Get("Field").Raw())
}

func Test_get_from_int(t *testing.T) {
	value := support.NewValue(12)
	require.Panics(t, func() {
		value.Get("field").Raw()
	})
}

func Test_get_collection_from_map_in_value(t *testing.T) {
	value := support.NewValue(map[string]string{"1": "12"})
	require.Equal(t, "12", value.Collection().First().Raw())
}

func Test_get_collection_from_string_in_value(t *testing.T) {
	value := support.NewValue("12")
	require.Equal(t, []interface{}{"12"}, value.Collection().Raw())
}

func Test_get_map_from_string_in_value(t *testing.T) {
	value := support.NewValue("12")
	require.Panics(t, func() { value.Map() })
}

func Test_get_string_from_collection_in_value(t *testing.T) {
	value := support.NewValue([]string{"12"})
	require.Equal(t, "12", value.String())
}

func Test_get_string_from_map_in_value(t *testing.T) {
	value := support.NewValue(map[string]string{"1": "12"})
	require.Equal(t, "12", value.String())
}

func Test_get_valid_int_from_value(t *testing.T) {
	value := support.NewValue(12)
	require.Equal(t, 12, value.Int())
}

func Test_get_invalid_int_from_value(t *testing.T) {
	value := support.NewValue("invalid_int")
	require.Panics(t, func() { value.Int() })
}

func Test_get_int_from_collection(t *testing.T) {
	value := support.NewValue([]int{12})
	require.Equal(t, 12, value.Int())
}

func Test_get_int_from_map(t *testing.T) {
	value := support.NewValue(map[string]int{"first": 12})
	require.Equal(t, 12, value.Int())
}

func Test_filled_nil(t *testing.T) {
	value := support.NewValue(nil)
	require.False(t, value.Filled())
}

func Test_filled_empty_string(t *testing.T) {
	value := support.NewValue("")
	require.False(t, value.Filled())
}

func Test_filled_empty_collection(t *testing.T) {
	value := support.NewValue([]string{})
	require.False(t, value.Filled())
}

func Test_slit_string(t *testing.T) {
	value := support.NewValue("val1,val2")
	require.Equal(t, []interface{}{"val1", "val2"}, value.Split(",").Raw())
}

func Test_set_string(t *testing.T) {
	value := support.NewValue("val1")
	require.Equal(t, "val1", value.Set("no_map", "v").Raw())
}

func Test_byte_slice_to_string(t *testing.T) {
	value := support.NewValue([]byte("dog"))
	require.Nil(t, value.Error())
	require.Equal(t, []byte("dog"), value.Raw())
	require.Equal(t, "dog", value.String())
}

func Test_byte_slice_to_int(t *testing.T) {
	value := support.NewValue([]byte("5"))
	require.Nil(t, value.Error())
	require.Equal(t, []byte("5"), value.Raw())
	require.Equal(t, 5, value.Int())
}

func Test_byte_slice_to_float(t *testing.T) {
	value := support.NewValue([]byte("5"))
	require.Nil(t, value.Error())
	require.Equal(t, []byte("5"), value.Raw())
	require.Equal(t, 5., value.Float())
}

func Test_byte_slice_to_bool(t *testing.T) {
	value := support.NewValue([]byte("true"))
	require.Nil(t, value.Error())
	require.Equal(t, []byte("true"), value.Raw())
	require.True(t, value.Bool())
}
