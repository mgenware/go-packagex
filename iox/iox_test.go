package iox

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/mgenware/go-packagex/v5/test"
)

func nonExistingPath() string {
	return filepath.Join(os.TempDir(), "___")
}

func TestPathExists(t *testing.T) {
	r, err := PathExists(os.TempDir())
	test.PanicIfErr(err)
	test.Compare(t, true, r)
	r, err = PathExists(nonExistingPath())
	test.PanicIfErr(err)
	test.Compare(t, false, r)
}

func TestFileExists(t *testing.T) {
	f, err := ioutil.TempFile("", "go-packagex.iox")
	defer os.Remove(f.Name())
	test.PanicIfErr(err)
	r, err := FileExists(f.Name())
	test.PanicIfErr(err)
	test.Compare(t, true, r)

	r, err = FileExists(nonExistingPath())
	test.PanicIfErr(err)
	test.Compare(t, false, r)
}

func TestDirectoryExists(t *testing.T) {
	r, err := DirectoryExists(os.TempDir())
	test.PanicIfErr(err)
	test.Compare(t, true, r)

	r, err = DirectoryExists(nonExistingPath())
	test.PanicIfErr(err)
	test.Compare(t, false, r)
}
