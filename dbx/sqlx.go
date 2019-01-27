package dbx

import (
	"database/sql"
	"fmt"
)

// Transact starts a database transaction and calls Commit() when no errors, otherwise Rollback() will be called.
func Transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}

// GetLastInsertIDUint64 calls result.LastInsertId() and returns its result as uint64.
func GetLastInsertIDUint64(result sql.Result) (uint64, error) {
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), err
}

// GetLastInsertIDInt calls result.LastInsertId() and returns its result as int.
func GetLastInsertIDInt(result sql.Result) (int, error) {
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), err
}

// GetLastInsertIDUint calls result.LastInsertId() and returns its result as uint.
func GetLastInsertIDUint(result sql.Result) (uint, error) {
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint(id), err
}

// GetRowsAffectedUint64 calls result.RowsAffected() and returns its result as uint64.
func GetRowsAffectedUint64(result sql.Result) (uint64, error) {
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return uint64(id), err
}

// GetRowsAffectedInt calls result.RowsAffected() and returns its result as int.
func GetRowsAffectedInt(result sql.Result) (int, error) {
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(id), err
}

// GetRowsAffectedUint calls result.RowsAffected() and returns its result as uint.
func GetRowsAffectedUint(result sql.Result) (uint, error) {
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return uint(id), err
}

// CheckRowsAffected returns an error if rows affected does not equal to a given number.
func CheckRowsAffected(result sql.Result, num int) error {
	rows, err := GetRowsAffectedInt(result)
	if err != nil {
		return err
	}
	if rows != num {
		return fmt.Errorf("Expected %v rows affected, got %v.", num, rows)
	}
	return nil
}

// CheckOneRowAffected returns an error if rows affected does not not equal to 1.
func CheckOneRowAffected(result sql.Result) error {
	return CheckRowsAffected(result, 1)
}
