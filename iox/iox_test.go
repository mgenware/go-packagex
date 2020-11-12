package iox

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/mgenware/go-packagex/v6/test"
)

func nonExistingPath() string {
	return filepath.Join(os.TempDir(), "___")
}

func TestPathExists(t *testing.T) {
	r, err := PathExists(os.TempDir())
	test.PanicIfErr(err)
	test.Assert(t, r, true)
	r, err = PathExists(nonExistingPath())
	test.PanicIfErr(err)
	test.Assert(t, r, false)
}

func TestFileExists(t *testing.T) {
	f, err := ioutil.TempFile("", "go-packagex.iox")
	defer os.Remove(f.Name())
	test.PanicIfErr(err)
	r, err := FileExists(f.Name())
	test.PanicIfErr(err)
	test.Assert(t, r, true)

	r, err = FileExists(nonExistingPath())
	test.PanicIfErr(err)
	test.Assert(t, r, false)
}

func TestDirectoryExists(t *testing.T) {
	r, err := DirectoryExists(os.TempDir())
	test.PanicIfErr(err)
	test.Assert(t, r, true)

	r, err = DirectoryExists(nonExistingPath())
	test.PanicIfErr(err)
	test.Assert(t, r, false)
}
