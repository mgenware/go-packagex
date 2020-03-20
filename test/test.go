package test

import (
	"runtime/debug"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Compare compares the given two values, and make testing fail if they are not equivalent.
func Compare(t *testing.T, want, got interface{}) {
	if diff := cmp.Diff(want, got); diff != "" {
		debug.PrintStack()
		t.Fatalf("Mismatch (-want +got):\n%s", diff)
	}
}

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
