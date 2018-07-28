package filepathx

import (
	"path/filepath"
	"strings"
)

func TrimExt(basename string) string {
	return strings.TrimSuffix(basename, filepath.Ext(basename))
}
