package stringsx

import "unicode/utf8"

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
