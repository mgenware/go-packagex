package mathx

import (
	"testing"

	"github.com/mgenware/go-packagex/v6/test"
)

func TestAbsInt(t *testing.T) {
	test.Assert(t, AbsInt(-1), 1)
	test.Assert(t, AbsInt(12), 12)
	test.Assert(t, AbsInt(0), 0)
}

func TestAbsInt64(t *testing.T) {
	test.Assert(t, AbsInt64(-1), int64(1))
	test.Assert(t, AbsInt64(12), int64(12))
	test.Assert(t, AbsInt64(0), int64(0))
}
