package queries

import (
	"database/sql"
	"fmt"
	"net/url"
	cursors "shared-library/utils"
)

// Get room of account
func GetRoomOfAccount(accountId float64) []struct {
	id        int
	accountid float64
	roomid    int
	is_owner  bool
} {
	accountCursor := cursors.AccountCursorTurn()

	rows, err := accountCursor.Query(`
        SELECT * FROM public.account_rooms 
        WHERE accountid = $1`, accountId)

	if err != nil {
		fmt.Errorf("Error: ", err.Error())
		return nil
	}

	rooms := make([]struct {
		id        int
		accountid float64
		roomid    int
		is_owner  bool
	}, 0)

	var accountid float64
	var id, roomid int
	var is_owner bool

	for rows.Next() {
		if err := rows.Scan(&id, &accountid, &roomid, &is_owner); err != nil {
			if err == sql.ErrNoRows {
				fmt.Errorf(err.Error())
				return nil
			}
		}
		// Add to struct data of keeped
		room := struct {
			id        int
			accountid float64
			roomid    int
			is_owner  bool
		}{
			id:        id,
			accountid: accountid,
			roomid:    roomid,
			is_owner:  is_owner,
		}

		rooms = append(rooms, room)
	}

	return rooms
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

// Check room is own the of room
func IsAccountOwnerRoom(accountId int) (bool, error) {
	accountCursor := cursors.AccountCursorTurn()
	fmt.Println(accountId) // 1

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

	fmt.Println("is own", is_owner)
	if is_owner == true {
		return true, nil
	}
	return false, nil
}

// How much there user in the room ?
func CheckMinAccountInRoom(roomId int) (int, error) {
	accountCursor := cursors.AccountCursorTurn()

	rows, err := accountCursor.Query("SELECT * FROM public.account_rooms WHERE roomid = $1", roomId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id, accountid, roomid int
	var is_owner bool

	counter := 0

	for rows.Next() {
		err := rows.Scan(&id, &accountid, &roomid, &is_owner)
		if err != nil {
			return 0, err
		}
		counter += 1
	}

	fmt.Println("counter: ", counter)
	if counter < 0 {
		return 0, nil
	}

	return counter, nil
}
