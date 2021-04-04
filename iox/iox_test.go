package iox

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mgenware/go-packagex/v5/test"
)

func newTempDir() string {
	s, err := os.MkdirTemp("", "iox_test")
	if err != nil {
		panic(err)
	}
	return s
}

func nonExistingPath() string {
	return filepath.Join(newTempDir(), "___")
}

func TestPathExists(t *testing.T) {
	r, err := PathExists(newTempDir())
	test.PanicIfErr(err)
	test.Assert(t, r, true)
	r, err = PathExists(nonExistingPath())
	test.PanicIfErr(err)
	test.Assert(t, r, false)
}

func TestFileExists(t *testing.T) {
	f, err := os.CreateTemp("", "go-packagex.iox")
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
	r, err := DirectoryExists(newTempDir())
	test.PanicIfErr(err)
	test.Assert(t, r, true)

	r, err = DirectoryExists(nonExistingPath())
	test.PanicIfErr(err)
	test.Assert(t, r, false)
}
