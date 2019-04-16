package test

import (
	"runtime/debug"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Compare compares the given two values, and make testing fail if they are not equivalent.
func Compare(t *testing.T, want, got interface{}) {
	debug.PrintStack()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("Mismatch (-want +got):\n%s", diff)
	}
}
