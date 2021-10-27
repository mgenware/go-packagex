package iox

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/mgenware/go-packagex/v6/test"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func newTempDir() string {
	s, err := os.MkdirTemp("", "iox_test")
	if err != nil {
		panic(err)
	}
	return s
}

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randUniqueDirPath() string {
	return filepath.Join(os.TempDir(), fmt.Sprintf("%v-%v", randString(5), time.Now().UnixNano()))
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

func TestMkdirp(t *testing.T) {
	dir := randUniqueDirPath()
	err := Mkdirp(dir)
	test.PanicIfErr(err)

	exist, err := DirectoryExists(dir)
	test.PanicIfErr(err)
	test.Assert(t, exist, true)
}

func TestCreateFile(t *testing.T) {
	file := filepath.Join(randUniqueDirPath(), randString(5))
	_, err := CreateFile(file)
	test.PanicIfErr(err)

	exist, err := FileExists(file)
	test.PanicIfErr(err)
	test.Assert(t, exist, true)
}
