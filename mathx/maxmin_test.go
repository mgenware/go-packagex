package mathx

import "testing"

func errorf(t *testing.T, expected interface{}, got interface{}) {
	t.Errorf("Expected %v, got %v", expected, got)
}

func TestMaxMinInt(t *testing.T) {
	var x, y int
	x = -1
	y = 1
	if r := MaxInt(x, y); r != y {
		errorf(t, y, r)
	}
	if r := MinInt(x, y); r != x {
		errorf(t, x, r)
	}
	if r := MaxInt(x, x); r != x {
		errorf(t, x, r)
	}
	if r := MinInt(x, x); r != x {
		errorf(t, x, r)
	}
}

func TestMaxMinInt64(t *testing.T) {
	var x, y int64
	x = -1
	y = 1
	if r := MaxInt64(x, y); r != y {
		errorf(t, y, r)
	}
	if r := MinInt64(x, y); r != x {
		errorf(t, x, r)
	}
	if r := MaxInt64(x, x); r != x {
		errorf(t, x, r)
	}
	if r := MinInt64(x, x); r != x {
		errorf(t, x, r)
	}
}

func TestMaxMinUint(t *testing.T) {
	var x, y uint
	x = 9
	y = 7970
	if r := MaxUint(x, y); r != y {
		errorf(t, y, r)
	}
	if r := MinUint(x, y); r != x {
		errorf(t, x, r)
	}
	if r := MaxUint(x, x); r != x {
		errorf(t, x, r)
	}
	if r := MinUint(x, x); r != x {
		errorf(t, x, r)
	}
}

func TestMaxMinUint64(t *testing.T) {
	var x, y uint64
	x = 9
	y = 7970
	if r := MaxUint64(x, y); r != y {
		errorf(t, y, r)
	}
	if r := MinUint64(x, y); r != x {
		errorf(t, x, r)
	}
	if r := MaxUint64(x, x); r != x {
		errorf(t, x, r)
	}
	if r := MinUint64(x, x); r != x {
		errorf(t, x, r)
	}
}

func TestMaxMinByte(t *testing.T) {
	var x, y byte
	x = 9
	y = 255
	if r := MaxByte(x, y); r != y {
		errorf(t, y, r)
	}
	if r := MinByte(x, y); r != x {
		errorf(t, x, r)
	}
	if r := MaxByte(x, x); r != x {
		errorf(t, x, r)
	}
	if r := MinByte(x, x); r != x {
		errorf(t, x, r)
	}
}
