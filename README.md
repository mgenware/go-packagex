# go-packagex

[![Build Status](https://travis-ci.org/mgenware/go-packagex.svg?branch=master)](http://travis-ci.org/mgenware/go-packagex)

Extra helper methods for Go. Requires **Go 1.6+**.

## Packages

### strconvx

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/strconvx?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/strconvx)

Installation:
```sh
go get github.com/mgenware/go-packagex/strconvx
```

Overview:
* Quick `ParseXXX` functions without `base` or `bitSize` parameters.

### stringsx

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/stringsx?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/stringsx)

Installation:
```sh
go get github.com/mgenware/go-packagex/stringsx
```

Overview:
* `SubString`, `SubStringFromStart`, `SubStringToEnd`.
* `Truncate`.

### mathx

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/mathx?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/mathx)

Installation:
```sh
go get github.com/mgenware/go-packagex/mathx
```

Overview:
* `Min`, `Max` for `int`, `int64`, `uint`, `uint64` and `byte` types.
* `Abs` for `int` and `int64`.

### slicex

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/slicex?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/slicex)

Installation:
```sh
go get github.com/mgenware/go-packagex/slicex
```

Overview:
* Check for deep equality for slices (types supported: `byte`, `int`, `uint`, `int64`, `uint64`).

### iox

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/iox?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/iox)

Installation:
```sh
go get github.com/mgenware/go-packagex/iox
```

Overview:
* Read text file.
* Check if a file/directory exists.


### httpx

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/httpx?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/httpx)

Installation:
```sh
go get github.com/mgenware/go-packagex/httpx
```

Overview:
* Set HTTP response `Content-Type` to common MIME types or from an file extension.

### database/sqlx

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/database/sqlx?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/database/sqlx)

Installation:
```sh
go get github.com/mgenware/go-packagex/database/sqlx
```

Overview:
* Helper function to start a database transaction.
* Convert `Result.LastInsertId` to `uint64`, `int`, `uint`.
* Convert `Result.RowsAffected` to `uint64`, `int`, `uint`.
* Limit number of rows affected.
* Common interface definition for `database/sql` package.

### templatex

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/templatex?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/templatex)

Installation:
```sh
go get github.com/mgenware/go-packagex/templatex
```

Overview:
* **Note that `templatex` replies on the `text/template`, not the `html/template` package.**
* Template parsing helper methods.
* Execute a template to string.

# License
MIT
