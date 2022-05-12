package test

import (
	"github.com/confetti-framework/framework/support/str"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_upper_first_with_empty_string(t *testing.T) {
	result := str.UpperFirst("")
	require.Equal(t, "", result)
}

func Test_upper_first_with_multiple_words(t *testing.T) {
	result := str.UpperFirst("a horse is happy")
	require.Equal(t, "A horse is happy", result)
}

func Test_in_slice_with_no_parameter(t *testing.T) {
	require.False(t, str.InSlice("phone"))
}

func Test_in_slice_with_one_non_existing_string(t *testing.T) {
	require.False(t, str.InSlice("phone", "bag"))
}

func Test_in_slice_with_one_existing_string(t *testing.T) {
	require.True(t, str.InSlice("phone", "phone"))
}

func Test_in_slice_with_multiple_one_matched_parameters(t *testing.T) {
	require.True(t, str.InSlice("phone", "TV", "phone", "tabel"))
}

func Test_in_slice_with_integer(t *testing.T) {
	require.True(t, str.InSlice(1, 0, 1))
}

func Test_After(t *testing.T) {
	// TODO: What if nothing is found?
	require.Equal(t, "", str.After("", ""))
	require.Equal(t, "", str.After("", "han"))
	require.Equal(t, "hannah", str.After("hannah", ""))
	require.Equal(t, "nah", str.After("hannah", "han"))
	require.Equal(t, "nah", str.After("hannah", "n"))
	require.Equal(t, "nah", str.After("eee hannah", "han"))
	require.Equal(t, "nah", str.After("ééé hannah", "han"))
	require.Equal(t, "hannah", str.After("hannah", "xxxx"))
	require.Equal(t, "nah", str.After("han0nah", "0"))
	require.Equal(t, "nah", str.After("han2nah", "2"))
}

func Test_AfterLast(t *testing.T) {
	// TODO: What if nothing is found?
	require.Equal(t, "", str.After("", ""))
	require.Equal(t, "", str.After("", "han"))
	require.Equal(t, "hannah", str.After("hannah", ""))
	require.Equal(t, "tte", str.AfterLast("yvette", "yve"))
	require.Equal(t, "e", str.AfterLast("yvette", "t"))
	require.Equal(t, "e", str.AfterLast("ééé yvette", "t"))
	require.Equal(t, "", str.AfterLast("yvette", "tte"))
	require.Equal(t, "yvette", str.AfterLast("yvette", "xxxx"))
	require.Equal(t, "te", str.AfterLast("yv0et0te", "0"))
	require.Equal(t, "te", str.AfterLast("yv0et0te", "0"))
	require.Equal(t, "te", str.AfterLast("yv2et2te", "2"))
	require.Equal(t, "foo", str.AfterLast("----foo", "---"))
}

func Test_Before(t *testing.T) {
	require.Equal(t, "hannah", str.Before("hannah", ""))
	require.Equal(t, "han", str.Before("hannah", "nah"))
	require.Equal(t, "ha", str.Before("hannah", "n"))
	require.Equal(t, "ééé ", str.Before("ééé hannah", "han"))
	require.Equal(t, "hannah", str.Before("hannah", "xxxx"))
	require.Equal(t, "han", str.Before("han0nah", "0"))
	require.Equal(t, "han", str.Before("han0nah", "0"))
	require.Equal(t, "han", str.Before("han2nah", "2"))
}

func Test_BeforeLast(t *testing.T) {
	require.Equal(t, "yve", str.BeforeLast("yvette", "tte"))
	require.Equal(t, "yvet", str.BeforeLast("yvette", "t"))
	require.Equal(t, "ééé ", str.BeforeLast("ééé yvette", "yve"))
	require.Equal(t, "", str.BeforeLast("yvette", "yve"))
	require.Equal(t, "yvette", str.BeforeLast("yvette", "xxxx"))
	require.Equal(t, "yvette", str.BeforeLast("yvette", ""))
	require.Equal(t, "yv0et", str.BeforeLast("yv0et0te", "0"))
	require.Equal(t, "yv0et", str.BeforeLast("yv0et0te", "0"))
	require.Equal(t, "yv2et", str.BeforeLast("yv2et2te", "2"))
}

func Test_Between(t *testing.T) {
	require.Equal(t, "abc", str.Between("abc", "", "c"))
	require.Equal(t, "abc", str.Between("abc", "a", ""))
	require.Equal(t, "abc", str.Between("abc", "", ""))
	require.Equal(t, "b", str.Between("abc", "a", "c"))
	require.Equal(t, "b", str.Between("dddabc", "a", "c"))
	require.Equal(t, "b", str.Between("abcddd", "a", "c"))
	require.Equal(t, "b", str.Between("dddabcddd", "a", "c"))
	require.Equal(t, "nn", str.Between("hannah", "ha", "ah"))
	require.Equal(t, "a]ab[b", str.Between("[a]ab[b]", "[", "]"))
	require.Equal(t, "foo", str.Between("foofoobar", "foo", "bar"))
	require.Equal(t, "bar", str.Between("foobarbar", "foo", "bar"))
}

func Test_Contains(t *testing.T) {
	require.True(t, str.Contains("taylor", "ylo"))
	require.True(t, str.Contains("taylor", "taylor"))
	require.False(t, str.Contains("taylor", "xxx"))
	require.False(t, str.Contains("taylor", ""))
	require.False(t, str.Contains("", ""))
}

func Test_ContainsFromSlice(t *testing.T) {
	require.True(t, str.ContainsFromSlice("taylor", []string{"ylo"}))
	require.True(t, str.ContainsFromSlice("taylor", []string{"xxx", "ylo"}))
	require.False(t, str.ContainsFromSlice("taylor", []string{"xxx"}))
	require.False(t, str.ContainsFromSlice("taylor", []string{}))
	require.False(t, str.ContainsFromSlice("taylor", []string{""}))
}

func Test_ContainsAllFromSlice(t *testing.T) {
	require.True(t, str.ContainsAllFromSlice("This is my name", []string{"This", "is"}))
	require.True(t, str.ContainsAllFromSlice("This is my name", []string{"my", "ame"}))
	require.True(t, str.ContainsAllFromSlice("taylor", []string{"tay", "ylo"}))
	require.False(t, str.ContainsAllFromSlice("taylor", []string{"xxx", "ylo"}))
	require.False(t, str.ContainsAllFromSlice("taylor", []string{"xxx", "tay"}))
	require.False(t, str.ContainsAllFromSlice("This is my name", []string{"are", "name"}))
	require.False(t, str.ContainsAllFromSlice("taylor", []string{}))
	require.False(t, str.ContainsAllFromSlice("taylor", []string{"", ""}))
}

func Test_EndsWith(t *testing.T) {
	require.True(t, str.EndsWith("This is my name", "name"))
	require.True(t, str.EndsWith("This is my name", "e"))
	require.True(t, str.EndsWith("jason", "on"))
	require.True(t, str.EndsWith("7", "7"))
	require.True(t, str.EndsWith("a7", "7"))
	require.False(t, str.EndsWith("jason", "no"))
	require.False(t, str.EndsWith("jason", ""))
	require.False(t, str.EndsWith("", ""))
	// Test for multibyte string support
	require.True(t, str.EndsWith("Jönköping", "öping"))
	require.True(t, str.EndsWith("Malmö", "mö"))
	require.True(t, str.EndsWith("Malmö", "mö"))
	require.False(t, str.EndsWith("Jönköping", "oping"))
	require.False(t, str.EndsWith("Malmö", "mo"))
	require.True(t, str.EndsWith("你好", "好"))
	require.False(t, str.EndsWith("你好", "你"))
	require.False(t, str.EndsWith("你好", "a"))
}

func Test_StartsWith(t *testing.T) {
	require.True(t, str.StartsWith("jason", "jas"))
	require.True(t, str.StartsWith("jason", "jason"))
	require.True(t, str.StartsWith("7a", "7"))
	require.True(t, str.StartsWith("7", "7"))
	require.False(t, str.StartsWith("jason", "J"))
	require.False(t, str.StartsWith("jason", ""))
	require.False(t, str.StartsWith("", ""))
	// Test for multibyte string support
	require.True(t, str.StartsWith("Jönköping", "Jö"))
	require.True(t, str.StartsWith("Malmö", "Malmö"))
	require.True(t, str.StartsWith("你好", "你"))
	require.False(t, str.StartsWith("Jönköping", "Jonko"))
	require.False(t, str.StartsWith("Malmö", "Malmo"))
	require.False(t, str.StartsWith("你好", "好"))
	require.False(t, str.StartsWith("你好", "a"))
}

func Test_Lower(t *testing.T) {
	require.Equal(t, "foo bar baz", str.Lower("FOO BAR BAZ"))
	require.Equal(t, "foo bar baz", str.Lower("fOo Bar bAz"))
}

func Test_Upper(t *testing.T) {
	require.Equal(t, "FOO BAR BAZ", str.Upper("foo bar baz"))
	require.Equal(t, "FOO BAR BAZ", str.Upper("fOo Bar bAZ"))
}

func Test_Finish(t *testing.T) {
	require.Equal(t, "abbc", str.Finish("ab", "bc"))
	require.Equal(t, "abbc", str.Finish("abbcbc", "bc"))
	require.Equal(t, "abcbbc", str.Finish("abcbbcbc", "bc"))
	require.Equal(t, "this/string/", str.Finish("this/string", "/"))
	require.Equal(t, "this/string/", str.Finish("this/string/", "/"))
}

func Test_Start(t *testing.T) {
	require.Equal(t, "/test/string", str.Start("test/string", "/"))
	require.Equal(t, "/test/string", str.Start("/test/string", "/"))
	require.Equal(t, "/test/string", str.Start("//test/string", "/"))
}

func Test_Length(t *testing.T) {
	require.Equal(t, 11, str.Length("foo bar baz"))
	require.Equal(t, 0, str.Length(""))
}

func Test_Title(t *testing.T) {
	// TODO: add more tests
	require.Equal(t, "New York", str.Title("new york"))
	require.Equal(t, "Νεα Υορκη", str.Title("νΕα υΟρΚη"))
}
