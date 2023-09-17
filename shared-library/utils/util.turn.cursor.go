package utils

import (
	"database/sql"
	db_connections "shared-library/database/connections"
)

func AccountCursorTurn() *sql.DB {
	return db_connections.AccountDbConnection()
}

func GameCursorTurn() *sql.DB {
	return db_connections.GameDbConnection()
}

func MemeCursorTurn() *sql.DB {
	return db_connections.MemeDbConnection()
}
