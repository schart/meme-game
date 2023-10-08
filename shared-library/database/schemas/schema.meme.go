package schemas

import (
	"database/sql"
	cursors "shared-library/utils"
)

var memeCursor *sql.DB

func MemeTextTable() {
	memeCursor = cursors.MemeCursorTurn()

	memeCursor.Exec(`
		CREATE TABLE memeText (
			id serial PRIMARY KEY,
			text varchar(50) NOT NULL
			 
 		)
	`)

}

func MemePhotoTable() {
	memeCursor = cursors.MemeCursorTurn()

	memeCursor.Exec(`
		CREATE TABLE memePhoto (
			id serial PRIMARY KEY,
			photoid varchar(60)
		)
	`)

}

func MemeCreateTables() {
	MemeTextTable()
	MemePhotoTable()
}
