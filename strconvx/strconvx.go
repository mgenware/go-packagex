package strconvx

import "strconv"

// ParseFloat returns strconv.ParseFloat(s, 64)
func ParseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// ParseFloatOrPanic panics if ParseFloat returns an error
func ParseFloatOrPanic(s string) float64 {
	r, err := ParseFloat(s)
	if err != nil {
		panic(err)
	}
	return r
}

// ParseInt returns strconv.Atoi(s)
func ParseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ParseIntOrPanic panics if ParseInt returns an error
func ParseIntOrPanic(s string) (int, error) {
	r, err := ParseInt(s)
	if err != nil {
		panic(err)
	}
	return r
}

// ParseUint returns strconv.ParseUint(s, 10, 0)
func ParseUint(s string) (uint, error) {
	return strconv.ParseUint(s, 10, 0)
}

// ParseUintOrPanic panics if ParseUint returns an error
func ParseUintOrPanic(s string) uint {
	r, err := ParseUint(s)
	if err != nil {
		panic(err)
	}
	return r
}
