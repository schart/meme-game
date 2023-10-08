package schemas

import (
	"database/sql"
	cursors "shared-library/utils"
)

var accountCursor *sql.DB

func AccountTable() {
	accountCursor = cursors.AccountCursorTurn()

	accountCursor.Exec(`
		CREATE TABLE account(
			id serial PRIMARY KEY,
			username varchar(30) UNIQUE NOT NULL,
			password varchar(80) NOT NULL
		)
	`)

}

func AccountRoomsTable() {
	accountCursor = cursors.AccountCursorTurn()

	accountCursor.Exec(`
		CREATE TABLE account_rooms(
			id SERIAL PRIMARY KEY,
			accountid INT NOT NULL,
			roomid INT NOT NULL,
			is_owner BOOLEAN NOT NULL 
		)
	`)
}

func AccountCreateTables() {
	AccountTable()
	AccountRoomsTable()
}
