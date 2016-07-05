package stringsx

import "unicode/utf8"

func SubString(str string, startIndex, endIndex int) string {
	return string([]rune(str)[startIndex:endIndex])
}

func SubStringFromStart(str string, startIndex int) string {
	endIndex := utf8.RuneCountInString(str)
	return SubString(str, startIndex, endIndex)
}

func SubStringToEnd(str string, endIndex int) string {
	return SubString(str, 0, endIndex)
}

func Truncate(str string, length int) string {
	return SubStringToEnd(str, length)
}
