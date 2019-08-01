package templatex

import (
	"io/ioutil"
	"testing"

	"github.com/mgenware/go-packagex/v5/iox"
	"github.com/mgenware/go-packagex/v5/test"
)

func newViewFile() string {
	t := iox.MustNewTempFileWithContent("view_test", []byte("1{{.}}"))
	return t.Path()
}

func TestView(t *testing.T) {
	file := newViewFile()
	v := MustParseView(file, false)
	got := v.MustExecuteToString("haha")
	test.Compare(t, "1haha", got)
	ioutil.WriteFile(file, []byte("2{{.}}"), 0644)
	got = v.MustExecuteToString("haha")
	test.Compare(t, "1haha", got)
}

func TestDevView(t *testing.T) {
	file := newViewFile()
	v := MustParseView(file, true)
	got := v.MustExecuteToString("haha")
	test.Compare(t, "1haha", got)
	ioutil.WriteFile(file, []byte("2{{.}}"), 0644)
	got = v.MustExecuteToString("haha")
	test.Compare(t, "2haha", got)
}
