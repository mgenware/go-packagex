package strconvx

import "strconv"

// ParseFloat returns strconv.ParseFloat(s, 64).
func ParseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// ParseInt returns strconv.Atoi(s).
func ParseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ParseUint returns strconv.ParseUint(s, 10, 0).
func ParseUint(s string) (uint, error) {
	r, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(r), nil
}

// ParseInt64 returns strconv.ParseInt(s, 10, 64).
func ParseInt64(s string) (int64, error) {
	r, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(r), nil
}

// ParseUint64 returns strconv.ParseUint(s, 10, 64).
func ParseUint64(s string) (uint64, error) {
	r, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(r), nil
}
