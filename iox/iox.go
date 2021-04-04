package iox

import (
	"os"
)

// ReadFileText behaves like Go's os.ReadFile, but returns a string instead.
func ReadFileText(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func pathExistsCore(path string) (os.FileInfo, error) {
	if fileInfo, err := os.Stat(path); err == nil {
		return fileInfo, nil
	} else if os.IsNotExist(err) {
		return nil, nil
	} else {
		return nil, err
	}
}

// PathExists checks if a specified path exists.
func PathExists(path string) (bool, error) {
	info, err := pathExistsCore(path)
	if err != nil {
		return false, err
	}
	return info != nil, nil
}

// FileExists checks if the specified file exists.
func FileExists(file string) (bool, error) {
	info, err := pathExistsCore(file)
	if err != nil {
		return false, err
	}
	return info != nil && !info.IsDir(), nil
}

// IsFile returns true if the specified path is a file.
func IsFile(path string) bool {
	fileExists, _ := FileExists(path)
	return fileExists
}

// DirectoryExists checks if the specified directory exists.
func DirectoryExists(dir string) (bool, error) {
	info, err := pathExistsCore(dir)
	if err != nil {
		return false, err
	}
	return info != nil && info.IsDir(), nil
}

// IsDirectory returns true if the specified path is a directory.
func IsDirectory(path string) bool {
	dirExists, _ := DirectoryExists(path)
	return dirExists
}
