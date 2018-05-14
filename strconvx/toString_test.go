package strconvx

import (
	"fmt"
	"math"
	"testing"
)

func TestToString(t *testing.T) {
	if r := ToString(math.MaxFloat64); r != fmt.Sprintf("%v", math.MaxFloat64) {
		t.Errorf("Expected %v, got %v", fmt.Sprintf("%v", math.MaxFloat64), r)
	}
}
