package strconvx

import "testing"

func TestParseFloat(t *testing.T) {
	if r, _ := ParseFloat("32"); r != 32 {
		t.Errorf("Expected %v, got %v", float64(32), r)
	}
	if r, _ := ParseFloat("-32.003"); r != -32.003 {
		t.Errorf("Expected %v, got %v", float64(-32.003), r)
	}
	if _, err := ParseFloat("aaa"); err == nil {
		t.Error("Error expected")
	}
}

func TestParseInt(t *testing.T) {
	if r, _ := ParseInt("32"); r != 32 {
		t.Errorf("Expected %v, got %v", int(32), r)
	}
	if r, _ := ParseInt("-3"); r != -3 {
		t.Errorf("Expected %v, got %v", int(-3), r)
	}
	if _, err := ParseInt("aaa"); err == nil {
		t.Error("Error expected")
	}
}

func TestParseUint(t *testing.T) {
	if r, _ := ParseUint("32"); r != 32 {
		t.Errorf("Expected %v, got %v", 32, r)
	}
	if _, err := ParseUint("-32"); err == nil {
		t.Error("Error expected")
	}
	if _, err := ParseUint("aaa"); err == nil {
		t.Error("Error expected")
	}
}

func TestParseInt64(t *testing.T) {
	if r, _ := ParseInt64("-9223372036854775808"); r != int64(-9223372036854775808) {
		t.Errorf("Expected %v, got %v", int64(-9223372036854775808), r)
	}
	if r, _ := ParseInt64("-9223372036854775808"); r != int64(-9223372036854775808) {
		t.Errorf("Expected %v, got %v", int64(-9223372036854775808), r)
	}
	if _, err := ParseInt64("18446744073709551615"); err == nil {
		t.Error("Error expected")
	}
}

func TestParseUint64(t *testing.T) {
	if r, _ := ParseUint64("18446744073709551615"); r != uint64(18446744073709551615) {
		t.Errorf("Expected %v, got %v", uint64(18446744073709551615), r)
	}
	if _, err := ParseUint64("-32"); err == nil {
		t.Error("Error expected")
	}
}
