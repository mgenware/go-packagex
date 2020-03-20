package mathx

import (
	"testing"

	"github.com/mgenware/go-packagex/v5/test"
)

func TestAbsInt(t *testing.T) {
	test.Compare(t, 1, AbsInt(-1))
	test.Compare(t, 12, AbsInt(12))
	test.Compare(t, 0, AbsInt(0))
}

func TestAbsInt64(t *testing.T) {
	test.Compare(t, int64(1), AbsInt64(-1))
	test.Compare(t, int64(12), AbsInt64(12))
	test.Compare(t, int64(0), AbsInt64(0))
}
