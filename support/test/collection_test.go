package test

import (
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_get_all_from_collection(t *testing.T) {
	values := support.NewCollection([]string{
		"Go",
		"David",
		"Sammy",
	})

	value := values.Get("*").Collection()
	require.Len(
		t,
		value,
		3,
	)
}

func Test_get_collection_by_key(t *testing.T) {
	values := support.NewMap(map[string][]string{
		"language": {"Go"},
		"names":    {"David", "Jona"},
	})

	languages := values.Get("language.*")
	require.Equal(t, "Go", languages.Collection().First().String())

	name := values.Get("names.*").Collection()[0].Raw()
	require.Equal(
		t,
		"David",
		name,
	)
}

func Test_get_collection_by_unknown_key(t *testing.T) {
	values := support.NewValue([]string{"house", "door"})

	result, err := values.GetE("2")
	require.EqualError(t, err, "key '2': can not found value")
	require.Equal(t, support.NewValue(emptyInterface), result)
}

func Test_get_collection_by_known_key(t *testing.T) {
	values := support.NewValue([]string{"house", "door"})

	result, err := values.GetE("1")
	require.Nil(t, err)
	require.Equal(t, "door", result.String())
}

func Test_collection_get_by_string(t *testing.T) {
	data := support.NewCollection([]string{})
	result, err := data.GetE("username")
	require.EqualError(t, err, "'username' can only be a number or *")
	require.Equal(t, support.NewValue(emptyInterface), result)
}

func Test_collection_push_value(t *testing.T) {
	data := support.NewCollection([]string{})
	data = data.Push(support.NewValue("apple_pear"))
	require.Equal(t, "apple_pear", data.Get("0").String())
}

func Test_collection_set_by_invalid_key(t *testing.T) {
	data := support.NewCollection([]string{})
	data, err := data.SetE("invalid_key", support.NewValue("apple_pear"))
	require.EqualError(t, err, "key 'invalid_key' can only begin with an asterisk or number")
	require.Equal(t, support.NewCollection([]string{}), data)
}

func Test_collection_set_asterisk(t *testing.T) {
	data := support.NewCollection([]string{})
	data, err := data.SetE("*", support.NewValue("apple_pear"))
	require.Nil(t, err)
	require.Equal(t, []interface{}{"apple_pear"}, data.Raw())
}

func Test_collection_set_nested_collection(t *testing.T) {
	data := support.NewCollection([]string{})
	data, err := data.SetE("*.*", support.NewValue("apple_pear"))
	require.Nil(t, err)
	require.Equal(t, []interface{}{[]interface{}{"apple_pear"}}, data.Raw())
}

func Test_collection_set_nested_collection_with_existing_data(t *testing.T) {
	data := support.NewCollection([]string{"berry"})
	data, err := data.SetE("*.*", support.NewValue("apple_pear"))
	require.Nil(t, err)
	require.Equal(t, []interface{}{"berry", []interface{}{"apple_pear"}}, data.Raw())
}

func Test_collection_set_collection(t *testing.T) {
	data := support.NewCollection()
	data, err := data.SetE("*", support.NewCollection(support.NewValue("apple_pear")))
	require.Nil(t, err)
	require.Equal(t, []interface{}{[]interface{}{"apple_pear"}}, data.Raw())
}

func Test_set_collection_on_empty_value(t *testing.T) {
	data := support.NewValue(nil)
	result, err := data.SetE("*", "water")
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{"water"}, result.Raw())
}

func Test_set_deep_collection_on_empty_value(t *testing.T) {
	data := support.NewValue(nil)
	result, err := data.SetE("*.*", "water")
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{[]interface{}{"water"}}, result.Raw())
}

func Test_set_collection_on_string_value(t *testing.T) {
	data := support.NewValue("rain")
	result, err := data.SetE("*", "water")
	assert.EqualError(t, err, "can not append value on 'string'")
	assert.Equal(t, support.NewValue("rain"), result)
}

func Test_set_on_empty_key(t *testing.T) {
	data := support.NewValue(nil)
	result, err := data.SetE("", "water")
	assert.Nil(t, err)
	assert.Equal(t, "water", result.Raw())
}

func Test_set_map_on_collection(t *testing.T) {
	var err error
	data := support.NewValue(nil)
	data, err = data.SetE("*.name", "Jaap")
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{map[string]interface{}{"name": "Jaap"}}, data.Raw())
}

func Test_collection_only_empty_collection(t *testing.T) {
	data := support.NewCollection()
	require.Equal(t, []interface{}(nil), data.Only().Raw())
}

func Test_collection_only_one_string(t *testing.T) {
	data := support.NewCollection("wolf")
	_, err := data.OnlyE("wolf")
	require.EqualError(t, err, "invalid only key on collection: 'wolf' can only be a number or *")
}

func Test_collection_only_nothing(t *testing.T) {
	data := support.NewCollection("wolf")
	require.Equal(t, []interface{}(nil), data.Only().Raw())
}

func Test_collection_only_with_asterisk(t *testing.T) {
	data := support.NewCollection("wolf")
	require.Equal(t, []interface{}{"wolf"}, data.Only("*").Raw())
}

func Test_collection_only_with_multiple_values(t *testing.T) {
	data := support.NewCollection("wolf", "lamb")
	require.Equal(t, []interface{}{"wolf", "lamb"}, data.Only("*").Raw())
}

func Test_collection_only_with_map(t *testing.T) {
	data := support.NewCollection(map[string]string{"lamp": "wool"})
	require.Equal(t, []interface{}{map[string]interface{}{"lamp": "wool"}}, data.Only("*").Raw())
}

func Test_collection_only_with_map_and_key(t *testing.T) {
	data := support.NewCollection(map[string]string{"lamp": "wool", "fish": "water"})
	require.Equal(t, []interface{}{map[string]interface{}{"lamp": "wool"}}, data.Only("*.lamp").Raw())
}

func Test_collection_reverse(t *testing.T) {
	data := support.NewCollection(map[string]string{"lamp": "wool", "fish": "water"})
	require.Equal(t, []interface{}{map[string]interface{}{"fish": "water", "lamp": "wool"}}, data.Reverse().Raw())
}

func Test_collection_not_contains(t *testing.T) {
	data := support.NewCollection(map[string]string{"lamp": "wool", "fish": "water"})
	require.False(t, data.Contains("ear"))
}

func Test_collection_contains(t *testing.T) {
	data := support.NewCollection("lamp", "water")
	require.True(t, data.Contains("lamp"))
}

func Test_collection_len(t *testing.T) {
	data := support.NewCollection("wool", "water")
	require.Equal(t, 2, data.Len())
}

func Test_collection_empty(t *testing.T) {
	data := support.NewCollection()
	require.True(t, data.Empty())
}

func Test_collection_not_empty(t *testing.T) {
	data := support.NewCollection("wool", "water")
	require.False(t, data.Empty())
}

var emptyInterface interface{}
