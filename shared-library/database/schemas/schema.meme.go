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
			text varchar(50)
		)
	`)
}

func MemePhotoTable() {
	memeCursor = cursors.TurnMemeCursor()
	memeCursor.Query(`
		CREATE TABLE memePhoto (
			photoId varchar(60)
		)
	`)

}

func MemeCreateT() {
	MemeTextTable()
	MemePhotoTable()
}
