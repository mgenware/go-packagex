package mathx

// MaxInt returns the larger of two int numbers.
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// MaxInt64 returns the larger of two int64 numbers.
func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

// MaxUint returns the larger of two uint numbers.
func MaxUint(x, y uint) uint {
	if x > y {
		return x
	}
	return y
}

// MaxUint64 returns the larger of two uint64 numbers.
func MaxUint64(x, y uint64) uint64 {
	if x > y {
		return x
	}
	return y
}

// MaxByte returns the larger of two byte numbers.
func MaxByte(x, y byte) byte {
	if x > y {
		return x
	}
	return y
}

// MinInt returns the smaller of two int numbers.
func MinInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// MinInt64 returns the smaller of two int64 numbers.
func MinInt64(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}

// MinUint returns the smaller of two uint numbers.
func MinUint(x, y uint) uint {
	if x > y {
		return y
	}
	return x
}

// MinUint64 returns the smaller of two uint64 numbers.
func MinUint64(x, y uint64) uint64 {
	if x > y {
		return y
	}
	return x
}

// MinByte returns the smaller of two byte numbers.
func MinByte(x, y byte) byte {
	if x > y {
		return y
	}
	return x
}
