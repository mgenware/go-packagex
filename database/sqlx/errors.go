package sqlx

import (
	"database/sql"
)

func GetLastInsertIDWithError(result sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetLastInsertIDUint64WithError(result sql.Result, err error) (uint64, error) {
	if err != nil {
		return 0, err
	}
	return GetLastInsertIDUint64(result)
}

func GetLastInsertIDIntWithError(result sql.Result, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	return GetLastInsertIDInt(result)
}

func GetLastInsertIDUintWithError(result sql.Result, err error) (uint, error) {
	if err != nil {
		return 0, err
	}
	return GetLastInsertIDUint(result)
}

func GetRowsAffectedWithError(result sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func GetRowsAffectedUint64WithError(result sql.Result, err error) (uint64, error) {
	if err != nil {
		return 0, err
	}
	return GetRowsAffectedUint64(result)
}

func GetRowsAffectedIntWithError(result sql.Result, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	return GetRowsAffectedInt(result)
}

func GetRowsAffectedUintWithError(result sql.Result, err error) (uint, error) {
	if err != nil {
		return 0, err
	}
	return GetRowsAffectedUint(result)
}

func CheckOneRowAffectedWithError(result sql.Result, err error) error {
	if err != nil {
		return err
	}
	return CheckOneRowAffected(result)
}
