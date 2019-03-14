package stringsx

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

// SubString returns a subset of a string with given startIndex and endIndex.
// Note that endIndex is not included in returned substring.
func SubString(str string, startIndex, endIndex int) string {
	return string([]rune(str)[startIndex:endIndex])
}

// SubStringFromStart returns a subset of a string with given startIndex.
func SubStringFromStart(str string, startIndex int) string {
	endIndex := utf8.RuneCountInString(str)
	return SubString(str, startIndex, endIndex)
}

// SubStringToEnd returns a subset of a string with given endIndex.
// Note that endIndex is not included in returned substring.
func SubStringToEnd(str string, endIndex int) string {
	return SubString(str, 0, endIndex)
}

// Truncate truncates a string with given length.
func Truncate(str string, length int) string {
	if utf8.RuneCountInString(str) <= length {
		return str
	}
	return SubStringToEnd(str, length)
}

// Reverse reverses the given string.
func Reverse(str string) string {
	chars := []rune(str)
	for i, j := 0, len(chars)-1; i < j; {
		chars[i], chars[j] = chars[j], chars[i]
		i++
		j--
	}
	return string(chars)
}

// JoinAll concatenates the elements of a to create a single string with given separator string.
func JoinAll(a []interface{}, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return fmt.Sprint(a[0])
	}

	buffer := &bytes.Buffer{}
	buffer.WriteString(fmt.Sprint(a[0]))
	for i := 1; i < len(a); i++ {
		buffer.WriteString(sep)
		buffer.WriteString(fmt.Sprint(a[i]))
	}
	return buffer.String()
}
