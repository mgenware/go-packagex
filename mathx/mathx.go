package mathx

// MaxInt returns the larger of two int numbers.
func MaxInt(x, y) int {
	if x > y {
		return x
	}
	return y
}

// MaxInt64 returns the larger of two int64 numbers.
func MaxInt64(x, y) int64 {
	if x > y {
		return x
	}
	return y
}

// MinInt returns the smaller of two int numbers.
func MinInt(x, y) int {
	if x > y {
		return y
	}
	return x
}

// MinInt64 returns the smaller of two int64 numbers.
func MinInt64(x, y) int64 {
	if x > y {
		return y
	}
	return x
}
