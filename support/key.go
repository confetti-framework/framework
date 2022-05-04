package support

type Key = string

func PrefixKey(searchablePrefix, key string) Key {
	if key != "" {
		searchablePrefix = searchablePrefix + "."
	}
	return searchablePrefix + key
}
