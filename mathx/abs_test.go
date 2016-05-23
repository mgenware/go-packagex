package mathx

import "testing"

func TestAbsInt(t *testing.T) {
	var x int
	x = -1
	if r := AbsInt(x); r != 1 {
		errorf(t, 1, r)
	}
	x = 12
	if r := AbsInt(x); r != 12 {
		errorf(t, 12, r)
	}
	x = 0
	if r := AbsInt(x); r != 0 {
		errorf(t, 0, r)
	}
}

func TestAbsInt64(t *testing.T) {
	var x int64
	x = -1
	if r := AbsInt64(x); r != 1 {
		errorf(t, 1, r)
	}
	x = 12
	if r := AbsInt64(x); r != 12 {
		errorf(t, 12, r)
	}
	x = 0
	if r := AbsInt64(x); r != 0 {
		errorf(t, 0, r)
	}
}
