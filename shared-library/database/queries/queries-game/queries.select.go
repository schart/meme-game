package queriesgame

import (
	"fmt"
	cursors "shared-library/utils"
)

// Check presence of room by id -> returing id
func IsThereRoom(link string) (int, error) {
	gameCursor := cursors.GameCursorTurn()

	rows, err := gameCursor.Query("SELECT * FROM public.rooms WHERE link = $1", link)
	defer rows.Close()

	if err != nil {
		return 0, err
	}

	var id int
	var accountid, _link string

	for rows.Next() {
		err := rows.Scan(&id, &accountid, &_link)
		if err != nil {
			return 0, err
		}
		return id, nil

	}

	return 0, err
}

// Get all room
func GetAllRoom() []struct {
	id        int
	accountId int
	link      string
} {
	gameCursor := cursors.GameCursorTurn()
	rows, err := gameCursor.Query("SELECT * FROM public.rooms")

	if err != nil {
		return nil
	}

	defer rows.Close()

	// Use slice for keep data
	rooms := make([]struct {
		id        int
		accountId int
		link      string
	}, 0)

	for rows.Next() {
		var id int
		var accountid int
		var link string

		err := rows.Scan(&id, &accountid, &link)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Add to struct data of keeped
		room := struct {
			id        int
			accountId int
			link      string
		}{
			id:        id,
			accountId: accountid,
			link:      link,
		}

		rooms = append(rooms, room)
	}

	return rooms
}
