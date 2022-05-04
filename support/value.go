package support

import (
	"github.com/confetti-framework/errors"
	"github.com/spf13/cast"
	"reflect"
	"strconv"
	"strings"
)

type Value struct {
	source interface{}
}

type nonValueError struct {
	error
}

func (w nonValueError) Unwrap() error {
	return w.error
}

func NewValue(val interface{}, preErr ...error) Value {
	if len(preErr) > 0 {
		val = nonValueError{preErr[0]}
	}

	switch val.(type) {
	case []byte:
		return Value{val}
	case Collection, Map:
		return Value{val}
	case Value:
		return val.(Value)
	}

	switch Kind(val) {
	case reflect.Slice, reflect.Array:
		result := NewCollection(val)
		return Value{result}
	case reflect.Map:
		result, err := NewMapE(val)
		if err != nil {
			val = nonValueError{err}
		}
		val = result
	}

	return Value{val}
}

func (v Value) Source() interface{} {
	return v.source
}

func (v Value) Raw() interface{} {
	if result, ok := v.source.(Value); ok {
		return result.Raw()
	}
	if result, ok := v.source.(Collection); ok {
		return result.Raw()
	}
	if result, ok := v.source.(Map); ok {
		return result.Raw()
	}

	return v.source
}

func (v Value) Get(key string) Value {
	result, err := v.GetE(key)
	if err != nil {
		panic(err)
	}
	return result
}

func (v Value) GetE(key string) (Value, error) {
	if key == "" {
		return v, nil
	}

	currentKey, rest := splitKey(key)
	nextKey := joinRest(rest)

	// when you request something with an Asterisk, you always develop a collection
	if currentKey == "*" {

		switch v.source.(type) {
		case Collection:
			return v.source.(Collection).GetE(nextKey)
		case Map:
			return v.source.(Map).GetE(nextKey)
		default:
			return Value{}, errors.New("*: is not a Collection or Map")
		}

	}

	switch source := v.source.(type) {
	case Collection:
		keyInt, err := strconv.Atoi(currentKey)
		if err != nil {
			return Value{}, err
		}
		collection := v.source.(Collection)
		if len(collection) < (keyInt + 1) {
			return Value{}, errors.Wrap(CanNotFoundValueError, "key '%s'%s", currentKey, getKeyInfo(key, currentKey))
		}
		return collection[keyInt].GetE(nextKey)
	case Map:
		value, ok := v.source.(Map)[currentKey]
		if !ok {
			return value, errors.Wrap(CanNotFoundValueError, "key '%s'%s", currentKey, getKeyInfo(key, currentKey))
		}
		return value.GetE(nextKey)
	default:
		switch Kind(source) {
		case reflect.Struct:
			val := reflect.ValueOf(source).FieldByName(currentKey)
			if val.IsValid() {
				return NewValue(val.Interface()).GetE(nextKey)
			} else {
				return Value{}, errors.New(currentKey + ": can't find value")
			}

		}
		return Value{}, errors.New(currentKey + ": is not a struct, Collection or Map")
	}
}

// A value can contain a collection.
func (v Value) Collection() Collection {
	if  err, isPre := v.source.(nonValueError); isPre {
		panic(err)
	}

	switch v.source.(type) {
	case Collection:
		return v.source.(Collection)
	case Map:
		return v.source.(Map).Collection()
	default:
		return NewCollection(v.source)
	}
}

func (v Value) Map() Map {
	result, err := v.MapE()
	if err != nil {
		panic(err)
	}
	return result
}

func (v Value) MapE() (Map, error) {
	source := v.source
	if  err, isPre := source.(nonValueError); isPre {
		return Map{}, errors.Unwrap(err)
	}

	switch valueType := source.(type) {
	case Map:
		return source.(Map), nil
	default:
		return nil, errors.New("can't create map from reflect.Kind " + strconv.Itoa(int(Kind(valueType))))
	}
}

func (v Value) String() string {
	result, err := v.StringE()
	if err != nil {
		panic(err)
	}

	return result
}

func (v Value) StringE() (string, error) {
	var result string
	var err error

	source := v.source
	if  err, isPre := source.(nonValueError); isPre {
		return "", errors.Unwrap(err)
	}

	switch source := source.(type) {
	case Collection:
		result, err = source.First().StringE()
	case Map:
		result, err = source.First().StringE()
	default:
		result, err = cast.ToStringE(source)
	}

	return result, err
}

func (v Value) Int() int {
	values, err := v.IntE()
	if err != nil {
		panic(err)
	}

	return values
}

func (v Value) IntE() (int, error) {
	var result int
	var err error

	source := v.source
	if  err, isPre := source.(nonValueError); isPre {
		return 0, errors.Unwrap(err)
	}

	switch source := source.(type) {
	case Collection:
		result, err = source.First().IntE()
	case Map:
		result, err = source.First().IntE()
	case []byte:
		result, err = cast.ToIntE(string(source))
	default:
		result, err = cast.ToIntE(source)
	}

	return result, err
}

func (v Value) Float() float64 {
	values, err := v.FloatE()
	if err != nil {
		panic(err)
	}

	return values
}

func (v Value) FloatE() (float64, error) {
	var result float64
	var err error

	source := v.source
	if  err, isPre := source.(nonValueError); isPre {
		return 0, errors.Unwrap(err)
	}

	switch source := source.(type) {
	case Collection:
		result, err = source.First().FloatE()
	case Map:
		result, err = source.First().FloatE()
	case []byte:
		result, err = cast.ToFloat64E(string(source))
	default:
		result, err = cast.ToFloat64E(source)
	}

	return result, err
}

func (v Value) Bool() bool {
	source := v.source
	switch v := source.(type) {
	case []byte:
		source = string(v)
	}

	switch source {
	case true, 1, "1", "true", "on", "yes":
		return true
	default:
		return false
	}
}

func (v Value) Filled() bool {
	return !(v.source == nil || len(v.Collection()) == 0 || v.source == "")
}

func (v Value) Empty() bool {
	return v.source == nil || v.source == ""
}

// Split slices Value into all substrings separated by separator and returns a slice of
// the strings between those separators.
//
// If Value does not contain separator and separator is not empty, Split returns a
// slice of length 1 whose only element is Value.
func (v Value) Split(separator string) Collection {
	rawStrings := strings.Split(v.String(), separator)
	var result Collection
	for _, rawString := range rawStrings {
		result = append(result, NewValue(rawString))
	}

	return NewCollection(result)
}

func (v Value) Set(key string, input interface{}) Value {
	result, err := v.SetE(key, input)
	if err != nil {
		panic(err)
	}
	return result
}

func (v Value) SetE(key string, input interface{}) (Value, error) {
	currentKey, _ := splitKey(key)
	if currentKey == "" {
		v.source = input
	}
	// if key is an asterisk, create collection if necessary
	if currentKey == "*" && v.source == nil {
		v.source = NewCollection()
	}
	if _, isCollection := v.source.(Collection); !isCollection && currentKey == "*" {
		return v, errors.Wrap(CanNotAppendValueError, "can not append value on '%s'", Kind(v.source))
	}
	// if value is nil, create a map to set the value
	if v.source == nil {
		v.source = NewMap()
	}

	switch source := v.source.(type) {
	case Map:
		nestedMap, err := source.SetE(key, input)
		if err != nil {
			return v, err
		}
		return NewValue(nestedMap), nil
	case Collection:
		collection, err := source.SetE(key, input)
		v.source = collection
		return v, err
	}
	return v, nil
}

func (v Value) Only(keys ...string) Value {
	result, err := v.OnlyE(keys...)
	if err != nil {
		panic(err)
	}
	return result
}

func (v Value) OnlyE(keys ...string) (Value, error) {
	switch source := v.source.(type) {
	case Map:
		result, err := source.OnlyE(keys...)
		return NewValue(result), err
	case Collection:
		result, err := source.OnlyE(keys...)
		return NewValue(result), err
	}
	return v, nil
}

func (v Value) Error() error {
	if  err, isPre := v.source.(nonValueError); isPre {
		return errors.Unwrap(err)
	}
	return nil
}

// convert keys with an asterisk to usable keys
func GetSearchableKeys(originKeys []string, value Value) []Key {
	var result []Key

	for _, originKey := range originKeys {
		keys := GetSearchableKeysByOneKey(originKey, value)
		result = append(result, keys...)
	}
	return result
}

// convert key with an asterisk to usable keys
func GetSearchableKeysByOneKey(originKey string, input Value) []Key {
	var keys []Key
	if !strings.Contains(originKey, "*") {
		return append(keys, originKey)
	}

	switch source := input.source.(type) {
	case Map:
		keys = getKeysByMap(keys, source, originKey)
	case Collection:
		keys = getKeysByCollection(keys, source, originKey)
	}

	return keys
}

func getKeysByCollection(keys []Key, source Collection, originKey string) []Key {
	_, rest := splitKey(originKey)
	for realKey, nestedValue := range source {
		nestedKeys := GetSearchableKeysByOneKey(joinRest(rest), nestedValue)
		keys = appendNestedKeys(keys, nestedKeys, strconv.Itoa(realKey))
	}
	return keys
}

func getKeysByMap(keys []Key, source Map, originKey string) []Key {
	current, rest := splitKey(originKey)
	for realKey, nestedValue := range source {
		if current == realKey || current == "*" {
			nestedKeys := GetSearchableKeysByOneKey(joinRest(rest), nestedValue)
			keys = appendNestedKeys(keys, nestedKeys, realKey)
		}
	}
	return keys
}

func appendNestedKeys(keys, nestedKeys []Key, realKey string) []Key {
	for _, nestedKey := range nestedKeys {
		keys = append(keys, PrefixKey(realKey, nestedKey))
	}
	return keys
}
