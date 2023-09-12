package utils

import (
	"database/sql"
	db_connections "shared-library/database/connections"
)

func TurnAccountCursor() *sql.DB {
	return db_connections.AccountDbConnection()
}

func TurnGameCursor() *sql.DB {
	return db_connections.GameDbConnection()
}

func TurnMemeCursor() *sql.DB {
	return db_connections.MemeDbConnection()
}
