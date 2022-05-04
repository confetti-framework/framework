package support

import (
	"github.com/confetti-framework/errors"
	"reflect"
	"strconv"
)

type Collection []Value

func NewCollection(items ...interface{}) Collection {
	collection := Collection{}

	for _, item := range items {
		if inputCollection, ok := item.(Collection); ok {
			collection = append(collection, inputCollection...)
			continue
		}

		switch Kind(item) {
		case reflect.Array, reflect.Slice:
			s := reflect.ValueOf(item)
			for i := 0; i < s.Len(); i++ {
				value := s.Index(i).Interface()
				collection = append(collection, NewValue(value))
			}
		default:
			collection = append(collection, NewValue(item))
		}
	}

	return collection
}

func (c Collection) Raw() interface{} {
	var result []interface{}
	var raw interface{}

	for _, value := range c {
		raw = value.Raw()
		result = append(result, raw)
	}

	return result
}

func (c Collection) Get(key string) Value {
	result, err := c.GetE(key)
	if err != nil {
		panic(err)
	}
	return result
}

func (c Collection) GetE(key string) (Value, error) {
	if key == "" {
		return NewValue(c), nil
	}

	currentKey, rest := splitKey(key)

	// when you request something with an Asterisk, you always develop a collection
	if currentKey == "*" {

		flattenCollection := Collection{}
		flattenMap := Map{}

		for _, value := range c {
			switch Kind(value.Source()) {
			case reflect.Slice, reflect.Array:
				flattenCollection = append(flattenCollection, value.Source().(Collection)...)
			case reflect.Map:
				flattenMap = value.Source().(Map).Merge(flattenMap)
			default:
				return NewValue(c), nil
			}
		}

		if len(flattenMap) > 0 {
			return flattenMap.GetE(joinRest(rest))
		}
		return flattenCollection.GetE(joinRest(rest))
	}

	index, err := strconv.Atoi(currentKey)
	if err != nil {
		return Value{}, errors.Wrap(InvalidCollectionKeyError, "'%s' can only be a number or *", key)
	}

	if len(c) < (index + 1) {
		return Value{}, errors.Wrap(CanNotFoundValueError, "'%s' not found", key)
	}

	return c[index].GetE(joinRest(rest))
}

func (c Collection) First() Value {
	if len(c) == 0 {
		return Value{}
	}

	return c[0]
}

func (c Collection) Push(item interface{}) Collection {
	return append(c, NewValue(item))
}

func (c Collection) SetE(key string, value interface{}) (Collection, error) {
	if key == "" {
		return c.Push(value), nil
	}

	currentKey, rest := splitKey(key)
	_, err := strconv.Atoi(currentKey)
	if currentKey != "*" && err != nil {
		return c, errors.Wrap(InvalidCollectionKeyError, "key '%s' can only begin with an asterisk or number", key)
	}

	if len(rest) == 0 {
		return c.Push(value), nil
	}
	nestedValue, err := NewValue(nil).SetE(joinRest(rest), value)
	if err != nil {
		return c, err
	}
	return c.Push(nestedValue), nil
}

func (c Collection) Reverse() Collection {
	items := c
	for left, right := 0, len(items)-1; left < right; left, right = left+1, right-1 {
		items[left], items[right] = items[right], items[left]
	}

	return items
}

// Determine if an item exists in the collection by a string
func (c Collection) Contains(search interface{}) bool {
	for _, item := range c {
		if item.Source() == search {
			return true
		}
	}

	return false
}

func (c Collection) Only(keys ...string) Collection {
	result, err := c.OnlyE(keys...)
	if err != nil {
		panic(err)
	}
	return result
}

func (c Collection) OnlyE(expectedKeys ...string) (Collection, error) {
	result := Collection{}
	var err error

	keys := GetSearchableKeys(expectedKeys, NewValue(c))
	for _, key := range keys {
		nestedValue, err := c.GetE(key)
		if errors.Is(err, InvalidCollectionKeyError) {
			return result, errors.Wrap(err, "invalid only key on collection")
		}
		if err == nil {
			result, err = result.SetE(key, nestedValue)
		}
	}

	return result, err
}

// The len method returns the length of the collection
func (c Collection) Len() int {
	return len(c)
}

func (c Collection) Empty() bool {
	return len(c) == 0
}
