package queries

import (
	"database/sql"
	"fmt"
	"net/url"
	cursors "shared-library/utils"
)

// Auth account
func IsAccountValidated(formData url.Values) (int, error) {
	accountCursor := cursors.AccountCursorTurn()

	row := accountCursor.QueryRow(`
        SELECT id FROM public.account 
        WHERE username = $1 and password = $2`, formData.Get("Username"), formData.Get("Password"))

	var id int
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return id, nil
}

// For check presence account
func IsThereAccount(accountId float64) (*sql.Rows, error) {
	accountCursor := cursors.AccountCursorTurn()

	rows, err := accountCursor.Query(`
		SELECT * FROM public.account 
		WHERE id = $1`, accountId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var id int
	var username, password string

	//  Set as nil when not found account info parameter
	// resFound := false

	for rows.Next() {
		err := rows.Scan(&id, &username, &password)
		if err != nil {
			fmt.Println(err)
			return nil, nil
		}

		fmt.Println(id, username, password)
		return rows, nil

	}

	return rows, nil
}

// Check presence of room by account id
func IsHaveARoomAccount(accountId float64) (bool, error) {
	accountCursor := cursors.AccountCursorTurn()

	rows, err := accountCursor.Query("SELECT * FROM public.account_rooms WHERE accountid = $1", accountId)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	var id, accountid, roomid int
	var is_owner bool

	for rows.Next() {
		err := rows.Scan(&id, &accountid, &roomid, &is_owner)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}
