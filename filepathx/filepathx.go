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
func TempFilePath(ext, prefix string) string {
	if prefix == "" {
		prefix = "tmp"
	}
	return path.Join(os.TempDir(), prefix+"-"+time.Now().Format("20060102150405")+ext)
}
