package iox

import (
	"io/ioutil"
	"os"
)

type TempFile struct {
	file *os.File
}

func NewTempFile(prefix string) (*TempFile, error) {
	file, err := ioutil.TempFile("", prefix)
	if err != nil {
		return nil, err
	}
	return &TempFile{file: file}, nil
}

func MustNewTempFileWithContent(prefix string, content []byte) *TempFile {
	tf, err := NewTempFile(prefix)
	if err != nil {
		panic(err)
	}

	tf.MustWriteContent(content)
	return tf
}

func (t *TempFile) Dispose() error {
	return os.Remove(t.file.Name())
}

func (t *TempFile) File() *os.File {
	return t.file
}

func (t *TempFile) WriteContent(content []byte) error {
	file := t.file
	if _, err := file.Write(content); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}

func (t *TempFile) MustWriteContent(content []byte) {
	err := t.WriteContent(content)
	if err != nil {
		panic(err)
	}
}
