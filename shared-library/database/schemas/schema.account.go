package schemas

import (
	"database/sql"
	cursors "shared-library/utils"
)

var accountCursor *sql.DB

func AccountTable() {
	accountCursor = cursors.AccountCursorTurn()
	accountCursor.Query(`
		CREATE TABLE account (
			id serial PRIMARY KEY,
			username varchar(30) UNIQUE NOT NULL,
			password varchar(80) NOT NULL
 		)
	`)

}

func AccountCreateTables() {
	AccountTable()
}
