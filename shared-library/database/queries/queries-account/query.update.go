package queries

import (
	"database/sql"
	cursors "shared-library/utils"
)

var accountCursor *sql.DB

func AccountPasswordUpdate(username, password string) error {
	accountCursor := cursors.AccountCursorTurn()

	tx, err := accountCursor.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback() // İşlem hata alınırsa geri alınacak

	_, err = tx.Exec(`
		UPDATE account 
		SET password = $1  
		WHERE username = $2;`, password, username)

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
