package schemas

import (
	"database/sql"
	cursors "shared-library/utils"
)

var gameCursor *sql.DB

func RoomTable() {
	gameCursor = cursors.GameCursorTurn()
	gameCursor.Exec(`
		CREATE TABLE rooms( 
			id serial PRIMARY KEY,
			accountid INT NOT NULL,
			link VARCHAR NOT NULL 

 		)
	`)
}

func GameCreateTables() {
	RoomTable()

}
