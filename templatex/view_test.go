package templatex

import (
	"os"
	"testing"
	"time"

	"github.com/mgenware/go-packagex/v5/iox"
	"github.com/mgenware/go-packagex/v5/test"
)

func newViewFile() string {
	t, err := iox.NewTempFileWithContent("", "view_test", []byte("1{{.}}"))
	if err != nil {
		panic(err)
	}
	return t.Path()
}

func TestView(t *testing.T) {
	file := newViewFile()
	v := MustParseView(file, false)
	got := v.MustExecuteToString("haha")
	test.Assert(t, got, "1haha")
	os.WriteFile(file, []byte("2{{.}}"), 0644)
	got = v.MustExecuteToString("haha")
	test.Assert(t, got, "1haha")
}

func TestDevView(t *testing.T) {
	file := newViewFile()
	v := MustParseView(file, true)
	got := v.MustExecuteToString("haha")
	test.Assert(t, got, "1haha")
	time.Sleep(time.Second)
	os.WriteFile(file, []byte("2{{.}}"), 0644)
	got = v.MustExecuteToString("haha")
	test.Assert(t, got, "2haha")
}
