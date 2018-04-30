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

// FileExists checks if a file exists.
func FileExists(file string) (bool, error) {
	info, err := os.Stat(file)
	if err != nil {
		return false, err
	}
	return !info.IsDir(), nil
}

// IsFile returns true if path points to a file and false otherwise.
func IsFile(path string) bool {
	fileExists, _ := FileExists(path)
	return fileExists
}

// DirectoryExists checks if a directory exists.
func DirectoryExists(dir string) (bool, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

// IsDirectory returns true if path points to a directory and false otherwise.
func IsDirectory(path string) bool {
	fileExists, _ := DirectoryExists(path)
	return fileExists
}
