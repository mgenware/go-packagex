package dbx

import (
	"database/sql"
)

func NewNullBool() sql.NullBool {
	return sql.NullBool{}
}

func NewNullFloat64() sql.NullFloat64 {
	return sql.NullFloat64{}
}

func NewNullInt64() sql.NullInt64 {
	return sql.NullInt64{}
}

func NewNullString() sql.NullString {
	return sql.NullString{}
}

func NewBool(value bool) sql.NullBool {
	return sql.NullBool{Bool: value, Valid: true}
}

func NewFloat64(value float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: value, Valid: true}
}

func NewInt64(value int64) sql.NullInt64 {
	return sql.NullInt64{Int64: value, Valid: true}
}

func NewString(value string) sql.NullString {
	return sql.NullString{String: value, Valid: true}
}
