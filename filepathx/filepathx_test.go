package filepathx

import (
	"path/filepath"
	"strings"
	"testing"
)

func errorf(t *testing.T, expected interface{}, got interface{}) {
	t.Errorf("Expected %v, got %v", expected, got)
}

func TestTrimExt(t *testing.T) {
	var got string
	var exp string

	got = TrimExt("a.json")
	exp = "a"
	if got != exp {
		errorf(t, exp, got)
	}

	got = TrimExt("abc")
	exp = "abc"
	if got != exp {
		errorf(t, exp, got)
	}

	got = TrimExt("a.json.x.yy")
	exp = "a.json.x"
	if got != exp {
		errorf(t, exp, got)
	}

	got = TrimExt(".js")
	exp = ""
	if got != exp {
		errorf(t, exp, got)
	}
}

func TestTempFilePath(t *testing.T) {
	var got string

	got = TempFilePath("", "")
	if !strings.Contains(got, "tmp-") {
		errorf(t, "contains tmp-", got)
	}

	got = TempFilePath(".jpeg", "")
	if filepath.Ext(got) != ".jpeg" {
		errorf(t, "has a jpeg ext", got)
	}

	got = TempFilePath("", "haha")
	if !strings.Contains(got, "haha-") {
		errorf(t, "contains haha-", got)
	}
}
