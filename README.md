# go-packagex

[![Build Status](https://travis-ci.org/mgenware/go-packagex.svg?branch=master)](http://travis-ci.org/mgenware/go-packagex)

Utils for Go(Golang) standard packages

## Packages

### mathx

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/mathx?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/mathx)

Installation:
```sh
go get github.com/mgenware/go-packagex/mathx
```

Overview:
* `Min`, `Max` for `int`, `int64`, `uint`, `uint64` types.
* `Abs` for `int` and `int64`.

### strconvx

[![GoDoc](https://godoc.org/github.com/mgenware/go-packagex/strconvx?status.svg)](http://godoc.org/github.com/mgenware/go-packagex/strconvx)

Installation:
```sh
go get github.com/mgenware/go-packagex/strconvx
```

Overview:
* Quick `ParseXXX` functions without `base` or `bitSize` parameters.

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

# License
MIT
