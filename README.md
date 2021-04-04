# go-packagex

[![Build](https://github.com/mgenware/go-packagex/actions/workflows/build.yml/badge.svg)](https://github.com/mgenware/go-packagex/actions/workflows/build.yml)

Extra helpers for Go (Go 1.11+).

Import path:

```go
import "github.com/mgenware/go-packagex/v5/<sub_package>"
```

## Packages

### strconvx

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/strconvx?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/strconvx)

Overview:

- Quick `ParseXXX` functions without `base` or `bitSize` parameters.

### stringsx

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/stringsx?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/stringsx)

Overview:

- `SubString` for Unicode chars(`rune`), `SubStringFromStart`, `SubStringToEnd`.
- `Truncate`.

### mathx

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/mathx?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/mathx)

Overview:

- `Min`, `Max` for `int`, `int64`, `uint`, `uint64` and `byte` types.
- `Abs` for `int` and `int64`.

### slicex

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/slicex?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/slicex)

Overview:

- Check for deep equality for slices (types supported: `byte`, `int`, `uint`, `int64`, `uint64`).

### iox

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/iox?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/iox)

Overview:

- Read text file.
- Check if a file/directory exists.

### httpx

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/httpx?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/httpx)

Overview:

- Set HTTP response `Content-Type` to common MIME types or from an file extension.

### dbx

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/dbx?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/dbx)

Overview:

- Helper function to start a database transaction.
- Convert `Result.LastInsertId` to `uint64`, `int`, `uint`.
- Convert `Result.RowsAffected` to `uint64`, `int`, `uint`.
- Limit number of rows affected.
- Common interface definition for `database/sql` package.

### templatex

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/templatex?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/templatex)

**Note that `templatex` is built upon the `text/template` package, not the `html/template`.**

Overview:

- Helpers for template parsing and executing.
- A wrapper of text/template Template, which can load template from a file and provide optional reloading support (useful in development mode).

### jsonx

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/jsonx?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/jsonx)

Overview:

- Fetch values of various types from a loosely typed map `map[string]interface{}`.
- Helper to marshal a loosely typed map `map[string]interface{}` from an array of bytes.

# License

MIT
