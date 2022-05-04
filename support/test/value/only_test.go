package value

import (
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_map_only_from_empty_value(t *testing.T) {
	value := support.NewValue(nil)
	require.Equal(t, nil, value.Only().Raw())
}

func Test_map_only_from_value(t *testing.T) {
	value := support.NewValue(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	require.Equal(t, map[string]interface{}{"username": "apple_pear"}, value.Only("username").Raw())
}

func Test_collection_only_from_value(t *testing.T) {
	value := support.NewValue([]string{"salamander", "koala"})
	require.Equal(t, []interface{}{"salamander", "koala"}, value.Only("*").Raw())
}

func Test_only_values_with_multiple_rules(t *testing.T) {
	data := support.NewValue(map[string]string{
		"title":       "Horse",
		"description": "Big animal",
	}).Only("title")
	require.Equal(t, map[string]interface{}{"title": "Horse"}, data.Raw())
}

func Test_only_nested_values(t *testing.T) {
	data := support.NewValue(map[string]map[string]string{
		"animal": {"title": "Horse"},
	}).Only("*.title")
	require.Equal(t, map[string]interface{}{"animal": map[string]interface{}{"title": "Horse"}}, data.Raw())
}

func Test_map_only_with_multiple_nested_values(t *testing.T) {
	data := support.NewValue(map[string]map[string]string{
		"book":  {"title": "Narnia"},
		"movie": {"title": "Lion King"},
	}).Only("*.title")
	require.Equal(t, map[string]interface{}{
		"book":  map[string]interface{}{"title": "Narnia"},
		"movie": map[string]interface{}{"title": "Lion King"},
	}, data.Raw())
}

func Test_only_collection_with_multiple_nested_values(t *testing.T) {
	data := support.NewValue([]map[string]string{
		{"title": "Narnia"},
		{"title": "Lion King"},
	}).Only("*.title")
	require.Equal(t, []interface{}{
		map[string]interface{}{"title": "Narnia"},
		map[string]interface{}{"title": "Lion King"},
	}, data.Raw())
}

func Test_only_map_with_one_invalid_keys(t *testing.T) {
	data := support.NewValue(map[string]string{"title": "Watch"}).Only("title", "description")
	require.Equal(t, map[string]interface{}{"title": "Watch"}, data.Raw())
}
