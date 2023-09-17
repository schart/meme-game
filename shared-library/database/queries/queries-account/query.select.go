package queries

import (
	"database/sql"
	"fmt"
	"net/url"
	cursors "shared-library/utils"
)

func AccountAllGet() (*sql.Rows, error) {
	res, err := accountCursor.Query(`SELECT * FROM public.account`)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AccountOneGet(formData url.Values) (*sql.Rows, error) {
	accountCursor = cursors.AccountCursorTurn()

	res, err := accountCursor.Query(`
		SELECT * FROM public.account 
		WHERE username = $1 and password = $2`, formData.Get("Username"), formData.Get("Password"))

	if err != nil {
		return nil, err
	}
	defer res.Close()

	var id int
	var username, password string

	//  Set as nil when not found account info  parameter
	resFound := false

	for res.Next() {
		err := res.Scan(&id, &username, &password)
		if err != nil {
			fmt.Println("Hata:", err)
		}
		fmt.Printf("ID: %d, Username: %s, Password: %s\n", id, username, password)
		resFound = true
	}

	if !resFound {
		// Account not founded
		return nil, nil
	}

	return res, nil
}
