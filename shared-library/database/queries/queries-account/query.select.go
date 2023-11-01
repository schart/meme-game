package queries

import (
	"database/sql"
	"fmt"
	"net/url"
	cursors "shared-library/utils"
)

// Get room of account
func GetRoomOfAccount(accountId float64) map[string]interface{} {
	accountCursor := cursors.AccountCursorTurn()

	row := accountCursor.QueryRow(`SELECT * FROM public.account_rooms WHERE accountid = $1`, accountId)

	var accountid float64
	var id, roomid int
	var is_owner bool

	data := make(map[string]interface{})

	if err := row.Scan(&id, &accountid, &roomid, &is_owner); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err.Error())
			return nil
		}
	}
	// Add to struct data of keeped

	data["id"] = id
	data["accountid"] = accountid
	data["roomid"] = roomid
	data["is_owner"] = is_owner

	return data
}

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
func IsThereAccount(accountId float64) (*sql.Row, error) {
	accountCursor := cursors.AccountCursorTurn()

	rows := accountCursor.QueryRow(`SELECT * FROM public.account  WHERE id = $1`, accountId)

	var id int

	err := rows.Scan(&id)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	if id == 0 {
		return nil, nil
	}

	return rows, nil
}

// Check presence of room by account id
func IsHaveARoomAccount(accountId float64) (bool, error) {
	accountCursor := cursors.AccountCursorTurn()

	rows := accountCursor.QueryRow("SELECT id FROM public.account_rooms WHERE accountid = $1", accountId)

	var id int

	err := rows.Scan(&id)
	if err != nil {
		return false, err
	}

	if id == 0 {
		return false, nil
	}

	return true, nil
}

// Check room is own the of room
func IsAccountOwnerRoom(accountId int) (bool, error) {
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

	}

	if is_owner == true {
		return true, nil
	}
	return false, nil
}

// How much there user in the room ?
func CheckMinAccountInRoom(roomId int) (int, error) {
	accountCursor := cursors.AccountCursorTurn()

	rows, err := accountCursor.Query("SELECT id FROM public.account_rooms WHERE roomid = $1", roomId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id int

	counter := 0

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}
		counter += 1
	}

	if counter < 0 {
		return 0, nil
	}

	return counter, nil
}

// Get account in the room with room link
func GetAccountsInRoom(roomId int) []float64 {
	accountCursor := cursors.AccountCursorTurn()

	rows, err := accountCursor.Query(`
        SELECT accountid FROM public.account_rooms 
        WHERE roomid = $1`, roomId)

	if err != nil {
		fmt.Errorf("Error: ", err.Error())
		return nil
	}

	accounts := []float64{}

	var accountid float64

	for rows.Next() {
		if err := rows.Scan(&accountid); err != nil {
			if err == sql.ErrNoRows {
				fmt.Errorf(err.Error())
				return nil
			}
		}

		// Add to struct data of keeped
		accounts = append(accounts, accountid)
	}

	return accounts
}

// Is account there via username?
func IsAccountThereUsername(username string) bool {
	accountCursor := cursors.AccountCursorTurn()

	/*
		@ We can just take id in the query response for do efficent the query
	*/
	result := accountCursor.QueryRow("SELECT id FROM public.account WHERE username = $1", username)

	var id int
	err := result.Scan(&id)
	if err != nil {
		fmt.Println("Error when check presence account:  ", err.Error())
		return false
	}

	if id == 0 {
		return false
	}

	return true
}
