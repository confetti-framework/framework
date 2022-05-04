package test

import (
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type testInterface interface{}
type testStruct struct{}

func Test_name_from_nil(t *testing.T) {
	name := support.Name((*testInterface)(nil))
	require.Equal(t, "test.testInterface", name)
}

func Test_name_from_struct(t *testing.T) {
	name := support.Name(testStruct{})
	require.Equal(t, "test.testStruct", name)
}

func Test_name_from_string(t *testing.T) {
	name := support.Name("InterfaceByString")
	require.Equal(t, "InterfaceByString", name)
}

func Test_name_from_empty_string(t *testing.T) {
	name := support.Name("")
	require.Equal(t, "string", name)
}

func Test_name_from_int(t *testing.T) {
	name := support.Name(1)
	require.Equal(t, "1", name)
}

func Test_name_from_empty_int(t *testing.T) {
	name := support.Name(0)
	require.Equal(t, "int", name)
}

func Test_name_from_float_64(t *testing.T) {
	name := support.Name(1.0)
	require.Equal(t, "1", name)
}

func Test_name_from_empty_float_64(t *testing.T) {
	name := support.Name(0.0)
	require.Equal(t, "float64", name)
}

func Test_name_from_float_32(t *testing.T) {
	var element float32 = 1.0
	name := support.Name(element)
	require.Equal(t, "1", name)
}

func Test_name_from_empty_float_32(t *testing.T) {
	var element float32 = 0
	name := support.Name(element)
	require.Equal(t, "float32", name)
}

func Test_name_from_bool(t *testing.T) {
	name := support.Name(false)
	require.Equal(t, "bool", name)
}

func Test_type_from_interface(t *testing.T) {
	reflectType := support.Kind((*testInterface)(nil))
	require.Equal(t, reflect.Ptr, reflectType)
}

func Test_type_from_collection(t *testing.T) {
	reflectType := support.Kind(support.NewCollection("v1"))
	require.Equal(t, reflect.Slice, reflectType)
}

func Test_type_from_string(t *testing.T) {
	reflectType := support.Kind("string")
	require.Equal(t, reflect.String, reflectType)
}

func Test_package_from_string(t *testing.T) {
	require.Panics(t, func() {
		support.Package("string")
	})
}

func Test_package_from_nil(t *testing.T) {
	require.NotPanics(t, func() {
		support.Package(nil)
	})
}
