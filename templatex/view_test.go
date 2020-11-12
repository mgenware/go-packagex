package templatex

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/mgenware/go-packagex/v6/iox"
	"github.com/mgenware/go-packagex/v6/test"
)

func newViewFile() string {
	t := iox.MustNewTempFileWithContent("view_test", []byte("1{{.}}"))
	return t.Path()
}

func TestView(t *testing.T) {
	file := newViewFile()
	v := MustParseView(file, false)
	got := v.MustExecuteToString("haha")
	test.Assert(t, got, "1haha")
	ioutil.WriteFile(file, []byte("2{{.}}"), 0644)
	got = v.MustExecuteToString("haha")
	test.Assert(t, got, "1haha")
}

func TestDevView(t *testing.T) {
	file := newViewFile()
	v := MustParseView(file, true)
	got := v.MustExecuteToString("haha")
	test.Assert(t, got, "1haha")
	time.Sleep(time.Second)
	ioutil.WriteFile(file, []byte("2{{.}}"), 0644)
	got = v.MustExecuteToString("haha")
	test.Assert(t, got, "2haha")
}
