package filepathx

import (
	"path/filepath"
	"testing"

	"github.com/mgenware/go-packagex/v6/test"
)

func errorf(t *testing.T, expected interface{}, got interface{}) {
	t.Errorf("Expected %v, got %v", expected, got)
}

func mustTempFilePath(dir, pattern, ext string) string {
	s, err := TempFilePath(dir, pattern, ext)
	if err != nil {
		panic(err)
	}
	return s
}

func TestTrimExt(t *testing.T) {
	var got string
	var exp string

	got = TrimExt("a.json")
	exp = "a"
	test.Assert(t, got, exp)

	got = TrimExt("abc")
	exp = "abc"
	test.Assert(t, got, exp)

	got = TrimExt("a.json.x.yy")
	exp = "a.json.x"
	test.Assert(t, got, exp)

	got = TrimExt(".js")
	exp = ""
	test.Assert(t, got, exp)
}

func TestTempFilePath(t *testing.T) {
	var got string

	got = mustTempFilePath("", "", ".jpeg")
	if filepath.Ext(got) != ".jpeg" {
		errorf(t, "has a jpeg ext", got)
	}
}
