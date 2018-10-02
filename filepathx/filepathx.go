package filepathx

import (
	"path/filepath"
	"strings"
)

// TrimExt removes the extension from the given path.
func TrimExt(basename string) string {
	return strings.TrimSuffix(basename, filepath.Ext(basename))
}
