package iox

import (
	"io/ioutil"
	"os"
)

// TempFile helps you create and remove an temporary file.
type TempFile struct {
	file *os.File
}

// NewTempFile creates an TempFile.
func NewTempFile(prefix string) (*TempFile, error) {
	file, err := ioutil.TempFile("", prefix)
	if err != nil {
		return nil, err
	}
	return &TempFile{file: file}, nil
}

// MustNewTempFileWithContent creates an TempFile with the given content, and panics if error occurred.
func MustNewTempFileWithContent(prefix string, content []byte) *TempFile {
	tf, err := NewTempFile(prefix)
	if err != nil {
		panic(err)
	}

	tf.MustSetContent(content)
	return tf
}

// Dispose removes the temporary file.
func (t *TempFile) Dispose() error {
	return os.Remove(t.file.Name())
}

// File returns the underlying os.File pointer.
func (t *TempFile) File() *os.File {
	return t.file
}

// Path returns the absolute path of the temporary file.
func (t *TempFile) Path() string {
	return t.file.Name()
}

// SetContent writes the given content to the temporary file.
func (t *TempFile) SetContent(content []byte) error {
	file := t.file
	if _, err := file.Write(content); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}

// MustSetContent calls SetContent and panics if error occurred.
func (t *TempFile) MustSetContent(content []byte) {
	err := t.SetContent(content)
	if err != nil {
		panic(err)
	}
}
