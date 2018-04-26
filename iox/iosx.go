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

// FileExist checks if a file exists.
func FileExist(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// DirectoryExist checks if a directory exists.
func DirectoryExist(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return info.IsDir()
}
