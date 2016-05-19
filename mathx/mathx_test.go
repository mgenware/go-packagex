package mathx

import "testing"

func TestMaxInt(t *testing.T) {
	var x, y int
	x = -1
	y = 1
	if r := MaxInt(x, y); r != y {
		t.Errorf("Expected %v, got %v", y, r)
	}
	if r := MaxInt(y, x); r != y {
		t.Errorf("Expected %v, got %v", y, r)
	}
	if r := MaxInt(y, y); r != y {
		t.Errorf("Expected %v, got %v", y, r)
	}
}

func TestMaxInt64(t *testing.T) {
	var x, y int64
	x = -1
	y = 1
	if r := MaxInt64(x, y); r != y {
		t.Errorf("Expected %v, got %v", y, r)
	}
	if r := MaxInt64(y, x); r != y {
		t.Errorf("Expected %v, got %v", y, r)
	}
	if r := MaxInt64(y, y); r != y {
		t.Errorf("Expected %v, got %v", y, r)
	}
}

func TestMinInt(t *testing.T) {
	var x, y int
	x = -1
	y = 1
	if r := MinInt(x, y); r != x {
		t.Errorf("Expected %v, got %v", x, r)
	}
	if r := MinInt(y, x); r != x {
		t.Errorf("Expected %v, got %v", x, r)
	}
	if r := MinInt(y, y); r != y {
		t.Errorf("Expected %v, got %v", y, r)
	}
}

func TestMinInt64(t *testing.T) {
	var x, y int64
	x = -1
	y = 1
	if r := MinInt64(x, y); r != x {
		t.Errorf("Expected %v, got %v", x, r)
	}
	if r := MinInt64(y, x); r != x {
		t.Errorf("Expected %v, got %v", x, r)
	}
	if r := MinInt64(y, y); r != y {
		t.Errorf("Expected %v, got %v", y, r)
	}
}
