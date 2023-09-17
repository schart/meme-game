package queries

import (
	"database/sql"
	cursors "shared-library/utils"
)

var accountCursor *sql.DB

func AccountPasswordUpdate(username, password string) error {
	accountCursor = cursors.AccountCursorTurn()

	_, err := accountCursor.Query(`
		UPDATE account 
		SET password = $1  
		WHERE password and username = $2;`, username, password)

	if err != nil {
		return err
	}

	return nil
}
