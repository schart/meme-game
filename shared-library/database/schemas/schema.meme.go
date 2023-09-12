package schemas

import (
	"database/sql"
	cursors "shared-library/utils"
)

var memeCursor *sql.DB

func MemeTextTable() {
	memeCursor = cursors.TurnMemeCursor()
	memeCursor.Query(`
		CREATE TABLE memeText (
			id serial PRIMARY KEY,
			text varchar(50) NOT NULL
			 
 		)
	`)

}

func MemePhotoTable() {
	memeCursor = cursors.TurnMemeCursor()
	memeCursor.Query(`
		CREATE TABLE memePhoto (
			id serial PRIMARY KEY,
			photoId varchar(60)
		)
	`)

}

func MemeCreateT() {
	MemeTextTable()
	MemePhotoTable()
}
