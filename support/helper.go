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
