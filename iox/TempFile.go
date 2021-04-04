package iox

import (
	"os"
)

// TempFile helps you create and remove an temporary file.
type TempFile struct {
	file *os.File
}

// NewTempFile creates an TempFile.
func NewTempFile(dir, pattern string) (*TempFile, error) {
	file, err := os.CreateTemp(dir, pattern)
	if err != nil {
		return nil, err
	}
	return &TempFile{file: file}, nil
}

// NewTempFileWithContent creates an TempFile with the given content.
func NewTempFileWithContent(dir, pattern string, content []byte) (*TempFile, error) {
	tf, err := NewTempFile(dir, pattern)
	if err != nil {
		return nil, err
	}

	err = tf.SetContent(content)
	if err != nil {
		return nil, err
	}
	return tf, nil
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
