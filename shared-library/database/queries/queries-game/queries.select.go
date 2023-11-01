package queriesgame

import (
	"fmt"
	cursors "shared-library/utils"
)

/*
   @ Queries of rooms table
   - Room available? via link
   - Get all room.
   - Get Room. via link
*/

func RoomAvailable(link string) bool {
	gameCursor := cursors.GameCursorTurn()

	rows := gameCursor.QueryRow("SELECT id FROM public.rooms WHERE link = $1", link)

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

func GetRoomByLink(roomLink string) map[string]interface{} {
	gameCursor := cursors.GameCursorTurn()

	row := gameCursor.QueryRow("SELECT * FROM public.rooms WHERE link = $1", roomLink)

	var room map[string]interface{} = map[string]interface{}{}

	var id int
	var accountid int
	var link string

	err := row.Scan(&id, &accountid, &link)
	if err != nil {
		fmt.Println(err)
	}

	room["id"] = id
	room["accountid"] = accountid
	room["link"] = link

	return room
}
