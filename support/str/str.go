package str

import "unicode"

func UpperFirst(input string) string {
	if len(input) == 0 {
		return ""
	}
	tmp := []rune(input)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}

func InSlice(input interface{}, expects ...interface{}) bool {
	for _, expect := range expects {
		if input == expect {
			return true
		}
	}
	return false
}
