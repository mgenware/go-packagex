package mathx

// AbsInt returns the absolute value of an int number.
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// AbsInt64 returns the absolute value of an int number.
func AbsInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
