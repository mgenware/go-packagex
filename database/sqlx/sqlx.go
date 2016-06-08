package sqlx

import "database/sql"

// Transact starts a database transaction and calls Commit() when no errors, otherwise Rollback() will be called.
func Transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}

// GetLastInsertIdUint64 calls result.LastInsertId() and returns its result as uint64.
func GetLastInsertIdUint64(result sql.Result) (uint64, error) {
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), err
}

// GetLastInsertIdInt calls result.LastInsertId() and returns its result as int.
func GetLastInsertIdInt(result sql.Result) (int, error) {
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), err
}

// GetLastInsertIdUint calls result.LastInsertId() and returns its result as uint.
func GetLastInsertIdUint(result sql.Result) (uint, error) {
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint(id), err
}
