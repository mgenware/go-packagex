package strconvx

import "strconv"

// ParseFloat64 calls strconv.ParseFloat(s, 64).
func ParseFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// ParseFloat32 calls strconv.ParseFloat(s, 32).
func ParseFloat32(s string) (float32, error) {
	value, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return float32(0), err
	}
	return float32(value), nil
}

// ParseInt calls strconv.Atoi(s).
func ParseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ParseUint calls strconv.ParseUint(s, 10, 0).
func ParseUint(s string) (uint, error) {
	r, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(r), nil
}

// ParseInt64 calls strconv.ParseInt(s, 10, 64).
func ParseInt64(s string) (int64, error) {
	r, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(r), nil
}

// ParseUint64 calls strconv.ParseUint(s, 10, 64).
func ParseUint64(s string) (uint64, error) {
	r, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(r), nil
}
