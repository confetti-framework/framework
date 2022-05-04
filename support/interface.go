package support

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

func Name(element interface{}) string {
	switch Kind(element) {
	case reflect.String:
		if element == "" {
			return "string"
		}
		return cast.ToString(element)
	case reflect.Int:
		if element == 0 {
			return "int"
		}
		return cast.ToString(element)
	case reflect.Bool:
		return "bool"
	case reflect.Float64:
		if element == 0. {
			return "float64"
		}
		return cast.ToString(element)
	case reflect.Float32:
		var emptyFloat float32 = 0
		if element == emptyFloat {
			return "float32"
		}
		return cast.ToString(element)
	case reflect.Struct:
		return reflect.TypeOf(element).String()
	}

	if Kind(element) == reflect.Ptr && element == nil {
		panic("Nil value found. To bind an interface, use the following syntax: (*INTERFACE)(nil)")
	}

	// Get name of element and ignore pointer
	return strings.TrimLeft(fmt.Sprintf("%T", element), "*")
}

func Package(element interface{}) string {
	if element == nil {
		return reflect.TypeOf(&element).Elem().PkgPath()
	}

	return reflect.TypeOf(element).Elem().PkgPath()
}

func Kind(element interface{}) reflect.Kind {
	if element == nil {
		return reflect.TypeOf(&element).Kind()
	}

	switch element.(type) {
	case Collection:
		return reflect.Slice
	case Map:
		return reflect.Map
	}

	return reflect.TypeOf(element).Kind()
}

func Dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
