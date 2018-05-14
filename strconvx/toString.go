package strconvx

import "strconv"

// ToStringInt calls strconv.FormatInt(int64(i), 10).
func ToStringInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

// ToStringUint calls strconv.FormatUint(uint64(i), 10).
func ToStringUint(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}

// ToStringInt64 calls strconv.FormatInt(i, 10).
func ToStringInt64(i int64) string {
	return strconv.FormatInt(i, 10)
}

// ToStringUint64 calls strconv.FormatUint(i, 10).
func ToStringUint64(i uint64) string {
	return strconv.FormatUint(i, 10)
}

// ToStringFloat64 calls strconv.FormatFloat(f, 'f', -1, 64).
func ToStringFloat64(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// ToStringFloat32 calls return strconv.FormatFloat(f, 'f', -1, 32).
func ToStringFloat32(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 32)
}
