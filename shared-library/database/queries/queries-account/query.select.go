package queries

import (
	"database/sql"
	"fmt"
	"net/url"
	cursors "shared-library/utils"
)

/*
	@ Queries of Account table

	- Get account. via username
	- Account verified or unverified?
	- Account available? via username
	- Account available? via id
*/

func GetAccountViaUsername(_username string) map[string]interface{} {
	accountCursor := cursors.AccountCursorTurn()
	row := accountCursor.QueryRow(`SELECT * FROM public.account WHERE username = $1`, _username)

	var id int
	var username, password string

	err := row.Scan(&id, &username, &password)
	if err != nil {
		return nil
	}

	if id == 0 {
		return nil
	}

	room := map[string]interface{}{}

	room["id"] = id
	room["username"] = username
	room["password"] = password

	return room
}

func AccountAvaliableViaId(accountId float64) bool {
	accountCursor := cursors.AccountCursorTurn()

	row := accountCursor.QueryRow(`SELECT id FROM public.account WHERE id = $1`, accountId)

	var id int
	err := row.Scan(&id)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(id)

	if id == 0 {
		return false
	}

	return true
}

func AccountVerified(formData url.Values) bool {
	accountCursor := cursors.AccountCursorTurn()

	row := accountCursor.QueryRow(`
        SELECT id FROM public.account 
        WHERE username = $1 and password = $2`, formData.Get("Username"), formData.Get("Password"))

	var id int

	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
	}

	if id == 0 {
		return false
	}

	return true
}

/*

	@ Queries of Account rooms table
	- Get room of account. via id
	- Account have the room? via id
	- Account owner the room? via id
	- How much account in the room? via room id
	- Get Accounts in Room. via roomid
*/

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

		if id == 0 {
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

func AccountHaveTheRoom(accountId float64) bool {
	accountCursor := cursors.AccountCursorTurn()

	rows := accountCursor.QueryRow("SELECT id FROM public.account_rooms WHERE accountid = $1", accountId)

	var id int

	err := rows.Scan(&id)
	if err != nil {
		return false
	}

	if id == 0 {
		return false
	}

	return true
}

func AccountOwnerTheRoom(accountId int) bool {
	accountCursor := cursors.AccountCursorTurn()

	row := accountCursor.QueryRow("SELECT * FROM public.account_rooms WHERE accountid = $1", accountId)

	var id, accountid, roomid int
	var is_owner bool

	err := row.Scan(&id, &accountid, &roomid, &is_owner)
	if err != nil {
		return false
	}

	if is_owner == true {
		return true
	}

	return false
}

func AccountCountInRoom(roomId int) int {
	accountCursor := cursors.AccountCursorTurn()

	rows, err := accountCursor.Query("SELECT id FROM public.account_rooms WHERE roomid = $1", roomId)
	if err != nil {
		return 0
	}
	defer rows.Close()

	var id int

	counter := 0

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0
		}
		counter += 1
	}

	if counter < 0 {
		return 0
	}

	return counter
}

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
