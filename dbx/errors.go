package dbx

import (
	"database/sql"
)

// GetLastInsertIDWithError checks a given error before calling GetLastInsertID.
func GetLastInsertIDWithError(result sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// GetLastInsertIDUint64WithError checks a given error before calling GetLastInsertIDUint64.
func GetLastInsertIDUint64WithError(result sql.Result, err error) (uint64, error) {
	if err != nil {
		return 0, err
	}
	return GetLastInsertIDUint64(result)
}

// GetLastInsertIDIntWithError checks a given  error before calling GetLastInsertIDInt.
func GetLastInsertIDIntWithError(result sql.Result, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	return GetLastInsertIDInt(result)
}

// GetLastInsertIDUintWithError checks a given error before calling GetLastInsertIDUint.
func GetLastInsertIDUintWithError(result sql.Result, err error) (uint, error) {
	if err != nil {
		return 0, err
	}
	return GetLastInsertIDUint(result)
}

// GetLastInsertIDInt64WithError checks a given error before calling GetLastInsertIDInt64.
func GetLastInsertIDInt64WithError(result sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// GetRowsAffectedWithError checks a given error before calling GetRowsAffected.
func GetRowsAffectedWithError(result sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// GetRowsAffectedUint64WithError checks a given error before calling GetRowsAffectedUint64.
func GetRowsAffectedUint64WithError(result sql.Result, err error) (uint64, error) {
	if err != nil {
		return 0, err
	}
	return GetRowsAffectedUint64(result)
}

// GetRowsAffectedIntWithError checks a given error before calling GetRowsAffectedInt.
func GetRowsAffectedIntWithError(result sql.Result, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	return GetRowsAffectedInt(result)
}

// GetRowsAffectedUintWithError checks a given error before calling GetRowsAffectedUint.
func GetRowsAffectedUintWithError(result sql.Result, err error) (uint, error) {
	if err != nil {
		return 0, err
	}
	return GetRowsAffectedUint(result)
}

// GetRowsAffectedInt64WithError checks a given error before calling GetRowsAffectedInt64.
func GetRowsAffectedInt64WithError(result sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// CheckOneRowAffectedWithError checks a given error before calling CheckOneRowAffected.
func CheckOneRowAffectedWithError(result sql.Result, err error) error {
	if err != nil {
		return err
	}
	return CheckOneRowAffected(result)
}
