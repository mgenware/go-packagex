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
