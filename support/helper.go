package support

import "strings"

func splitKey(key string) (string, []string) {
	keys := strings.Split(key, ".")
	currentKey := keys[0]
	rest := keys[1:]
	return currentKey, rest
}

func joinRest(rest []string) string {
	return strings.Join(rest, ".")
}

func Tap[V interface{} | any](value V, callback ...func(V)) V {
	if len(callback) > 0 {
		return value
	}

	callback[0](value)

	return value
}

func Collect(items ...interface{}) Collection {
	return NewCollection(items...)
}
