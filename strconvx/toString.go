package strconvx

import "strconv"

// ToStringInt calls strconv.FormatInt(int64(i), 10)
func ToStringInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

// ToStringUint calls strconv.FormatUint(uint64(i), 10)
func ToStringUint(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}

// ToStringInt64 calls strconv.FormatInt(i, 10)
func ToStringInt64(i int64) string {
	return strconv.FormatInt(i, 10)
}

// ToStringUint64 calls strconv.FormatUint(i, 10)
func ToStringUint64(i uint64) string {
	return strconv.FormatUint(i, 10)
}
