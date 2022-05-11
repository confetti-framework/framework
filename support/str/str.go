package str

import (
	"regexp"
	"strings"
	"unicode"
)

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

// Return the remainder of a string after the first occurrence of a given value.
func After(subject string, search string) string {
	if len(search) == 0 {
		return subject
	}
	results := strings.SplitN(subject, search, 2)
	return results[len(results)-1]
}

// Return the remainder of a string after the last occurrence of a given value.
func AfterLast(subject string, search string) string {
	if len(search) == 0 {
		return subject
	}
	position := strings.LastIndex(subject, search)

	if position == -1 {
		return subject
	}

	return subject[position+len(search):]
}

// Get the portion of a string before the first occurrence of a given value.
func Before(subject string, search string) string {
	if len(search) == 0 {
		return subject
	}
	position := strings.Index(subject, search)

	if position == -1 {
		return subject
	}

	return subject[:position]
}

func BeforeLast(subject string, search string) string {
	if len(search) == 0 {
		return subject
	}
	position := strings.LastIndex(subject, search)

	if position == -1 {
		return subject
	}

	return subject[:position]
}

func Between(subject string, from string, to string) string {
	if len(from) == 0 || len(to) == 0 {
		return subject
	}

	return BeforeLast(After(subject, from), to)
}

func Contains(haystack string, needle string) bool {
	if len(needle) == 0 {
		return false
	}

	return strings.Contains(haystack, needle)
}

func ContainsFromSlice(haystack string, needles []string) bool {
	if len(needles) == 0 {
		return false
	}

	for _, needle := range needles {
		if Contains(haystack, needle) {
			return true
		}
	}

	return false
}

func ContainsAllFromSlice(haystack string, needles []string) bool {
	if len(needles) == 0 {
		return false
	}

	for _, needle := range needles {
		if !Contains(haystack, needle) {
			return false
		}
	}

	return true
}

func EndsWith(haystack string, needle string) bool {
	if len(needle) == 0 {
		return false
	}

	return strings.HasSuffix(haystack, needle)
}

func StartsWith(haystack string, needle string) bool {
	if len(needle) == 0 {
		return false
	}

	return strings.HasPrefix(haystack, needle)
}

func Lower(value string) string {
	return strings.ToLower(value)
}

func Upper(value string) string {
	return strings.ToUpper(value)
}

func Finish(value string, cap string) string {
	quoted := regexp.QuoteMeta(cap)

	re := regexp.MustCompile("(?:" + quoted + ")+$")
	return re.ReplaceAllString(value, "") + cap
}

func Start(value string, prefix string) string {
	quoted := regexp.QuoteMeta(prefix)

	re := regexp.MustCompile("^(?:" + quoted + ")+")
	return prefix + re.ReplaceAllString(value, "")
}

//
//func Title(value string) string {
//  // TODO
//	return ""
//}
//
// func Kebab(vale string) string {
// 	// TODO
// 	return ""
// }
//
// func Length(value string) int {
// 	// TODO
// 	return 0
// }
//
// func LimitCharacters(value string, limit int, end string) string{
// 	// TODO
// 	return ""
// }
//
// func LimitWords(value string, limit int, end string) string{
// 	// TODO
// 	return ""
// }
//
// func PadBoth(value string, length int, pad string) string {
// 	// TODO
// 	return ""
// }
//
// func PadLeft(value string, length int, pad string) string {
// 	// TODO
// 	return ""
// }
//
// func PadRight(value string, length int, pad string) string {
// 	// TODO
// 	return ""
// }
//
// func ReplaceArray(search string, replace []string, subject string) string {
// 	// TODO
// 	return ""
// }
//
// func ReplaceFirst(search string, replace string, subject string) string {
// 	// TODO
// 	return ""
// }
//
// func ReplaceLast(search string, replace string, subject string) string {
// 	// TODO
// 	return ""
// }
//
// func Slug(value string) string {
// 	// TODO
// 	return ""
// }
//
// func SlugWithDelimiter(value string, delimiter string) string {
// 	// TODO
// 	return ""
// }
//
// func Snake(value string) string {
// 	// TODO
// 	return ""
// }
//
// func Studly(value string) string {
// 	// TODO
// 	return ""
// }
//
// func UcFirst(value string) string {
// 	// TODO
// 	return ""
// }
//
