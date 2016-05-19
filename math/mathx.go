package mathx

// MaxInt returns the larger of 2 int numbers.
func MaxInt(x, y) int {
	if x > y {
		return x
	}
	return y
}

// MaxInt64 returns the larger of 2 int64 numbers.
func MaxInt64(x, y) int64 {
	if x > y {
		return x
	}
	return y
}

// MinInt returns the smaller of 2 int numbers.
func MinInt(x, y) int {
	if x > y {
		return y
	}
	return x
}

// MinInt returns the smaller of 2 int64 numbers.
func MinInt64(x, y) int64 {
	if x > y {
		return y
	}
	return x
}
