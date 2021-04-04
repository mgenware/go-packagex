package filepathx

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// TrimExt removes the extension from the given path.
func TrimExt(basename string) string {
	return strings.TrimSuffix(basename, filepath.Ext(basename))
}

// TempFilePath returns a unique tmp file path (note the file is not created).
func TempFilePath(dir, pattern, extension string) (string, error) {
	s, err := os.MkdirTemp(dir, pattern)
	if err != nil {
		return "", err
	}
	return path.Join(s, time.Now().Format("20060102150405")+extension), nil
}
