package support

import (
	"github.com/confetti-framework/errors"
	"github.com/spf13/cast"
	"reflect"
)

type Map map[string]Value

func NewMap(itemsRange ...interface{}) Map {
	result, err := NewMapE(itemsRange...)
	if err != nil {
		panic(err)
	}
	return result
}

func NewMapE(itemsRange ...interface{}) (Map, error) {
	var err error
	result := Map{}

	for _, rawItems := range itemsRange {
		v := reflect.ValueOf(rawItems)
		if v.Kind() != reflect.Map {
			err = errors.WithStack(errors.Wrap(CanNotCreateMapError, "type %s", v.Kind().String()))
			continue
		}

		for _, key := range v.MapKeys() {
			value := v.MapIndex(key).Interface()
			key, err := cast.ToStringE(key.Interface())
			if err != nil {
				return nil, errors.WithStack(errors.Wrap(err, "invalid key in map"))
			}
			result[key] = NewValue(value)
		}
	}

	return result, err
}

func (m Map) Raw() interface{} {
	result := map[string]interface{}{}

	for key, value := range m {
		// Handle value
		result[key] = value.Raw()
	}

	return result
}

func (m Map) Get(key string) Value {
	result, err := m.GetE(key)
	if err != nil {
		panic(err)
	}
	return result
}

// GetE gets the first value associated with the given key.
// If there are no values associated with the key, GetE returns
// nil. To access multiple values, use GetCollection or Collection.
func (m Map) GetE(key string) (Value, error) {
	if key == "" {
		return NewValue(m), nil
	}

	currentKey, rest := splitKey(key)

	// when you request something with an asterisk, you always develop a collection
	if currentKey == "*" {
		collection := Collection{}
		for _, value := range m {
			nestedValueRaw, err := value.GetE(joinRest(rest))
			if err != nil {
				return nestedValueRaw, err
			}
			switch nestedValues := nestedValueRaw.source.(type) {
			case Collection:
				for _, nestedValue := range nestedValues {
					collection = collection.Push(nestedValue)
				}
			case Map:
				for _, nestedValue := range nestedValues {
					collection = collection.Push(nestedValue)
				}
			default:
				// If there are no keys to search further, the nested value is the final value
				collection = collection.Push(nestedValueRaw)
			}
		}

		return NewValue(collection), nil
	}

	value, found := m[key]
	if found {
		return value, nil
	}
	value, found = m[currentKey]
	if !found {
		return Value{}, errors.Wrap(CanNotFoundValueError, "key '%s'%s", currentKey, getKeyInfo(key, currentKey))
	}

	switch value.Source().(type) {
	case Collection:
		return value.Collection().GetE(joinRest(rest))
	case Map:
		return value.Map().GetE(joinRest(rest))
	default:
		return value.GetE(joinRest(rest))
	}
}

// SetE sets the key to value by dot notation
func (m Map) SetE(key string, input interface{}) (Map, error) {
	currentKey, rest := splitKey(key)
	value := NewValue(input)

	// If we have a dot notation we want to set the value deeper
	if key != currentKey {
		currentValue := m[currentKey]
		nestedValue, err := currentValue.SetE(joinRest(rest), value)
		if err != nil {
			return m, err
		}
		m[currentKey] = nestedValue
	} else {
		m[currentKey] = NewValue(input)
	}

	return m, nil
}

func (m Map) Only(keys ...string) Map {
	result, err := m.OnlyE(keys...)
	if err != nil {
		panic(err)
	}
	return result
}

func (m Map) OnlyE(originKeys ...string) (Map, error) {
	result := Map{}
	var err error

	keys := GetSearchableKeys(originKeys, NewValue(m))
	for _, key := range keys {
		item, err := m.GetE(key)
		if err == nil {
			result, err = result.SetE(key, item)
		}
	}

	return result, err
}

func (m Map) Except(keys ...string) Map {
	result := m.Copy()
	for _, key := range keys {
		delete(result, key)
	}

	return result
}

// Push adds the value to key. It appends to any existing values
// associated with key. If the value is in collection, push
// the value to the collection.
func (m Map) Push(key string, input interface{}) Map {
	if rawValue, found := m[key]; found {
		source := rawValue.Source()
		switch source.(type) {
		case Collection:
			collection := source.(Collection)
			m[key] = NewValue(collection.Push(input))
		default:
			m[key] = NewValue(input)
		}
	} else {
		m[key] = NewValue(input)
	}

	return m
}

// Delete deletes the values associated with key.
func (m Map) Delete(key string) {
	delete(m, key)
}

func (m Map) Collection() Collection {
	collection := Collection{}
	for _, value := range m {
		collection = collection.Push(value)
	}

	return collection
}

func (m Map) Merge(maps ...Map) Map {
	for _, bag := range maps {
		for key, item := range bag {
			m.Push(key, item)
		}
	}

	return m
}

// Copy generates a new struct with the same data as the old struct
func (m Map) Copy() Map {
	newMap := Map{}
	for key, value := range m {
		newMap[key] = value
	}

	return newMap
}

func (m Map) First() Value {
	return m.Collection().First()
}

func (m Map) Has(keys ...string) bool {
	for _, key := range keys {
		_, err := m.GetE(key)
		if err != nil {
			return false
		}
	}

	return true
}

func (m Map) HasAny(keys ...string) bool {
	if len(keys) == 0 {
		return true
	}

	for _, key := range keys {
		result, err := m.GetE(key)
		if err == nil && !result.Empty() {
			return true
		}
	}

	return false
}

func (m Map) Missing(keys ...string) bool {
	return !m.Has(keys...)
}

func (m Map) Filled(keys ...string) bool {
	for _, key := range keys {
		result, err := m.GetE(key)
		if err != nil || result.Empty() {
			return false
		}
	}

	return true
}

func (m Map) Empty() bool {
	return len(m) == 0
}

func getKeyInfo(key string, currentKey string) string {
	info := ""
	if currentKey != key {
		info = " ('" + key + "')"
	}
	return info
}
