package queriesgame

import (
	"fmt"
	cursors "shared-library/utils"
)

// Check presence of room by id -> returing id
func IsThereRoom(link string) (int, error) {
	gameCursor := cursors.GameCursorTurn()

	rows := gameCursor.QueryRow("SELECT id FROM public.rooms WHERE link = $1", link)

	var id int

	err := rows.Scan(&id)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, nil
	}

	return id, nil
}

// Get all room
func GetAllRoom() []interface{} {
	gameCursor := cursors.GameCursorTurn()
	rows, err := gameCursor.Query("SELECT * FROM public.rooms")

	if err != nil {
		return nil
	}

	defer rows.Close()

	var rooms []interface{}
	var room map[string]interface{} = map[string]interface{}{}

	for rows.Next() {
		var id int
		var accountid int
		var link string

		err := rows.Scan(&id, &accountid, &link)
		if err != nil {
			fmt.Println(err)
			continue
		}

		room["id"] = id
		room["accountid"] = accountid
		room["link"] = link

		rooms = append(rooms, room)
	}

	return rooms
}
