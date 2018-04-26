package iox

import (
	"io/ioutil"
	"os"
)

// ReadFileText behaves like Go's ioutil.ReadFile, but returns a string instead.
func ReadFileText(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// IsFile checks if a file exists.
func IsFile(file string) (bool, error) {
	info, err := os.Stat(file)
	if err != nil {
		return false, err
	}
	return !info.IsDir(), nil
}

// IsDirectory checks if a directory exists.
func IsDirectory(dir string) (bool, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}
