package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func AccountDbConnection() *sql.DB {
	connStr := "user=" + os.Getenv("DB_ADMIN_NAME") + " password=" + os.Getenv("DB_ADMIN_PASSWORD") + " dbname=" + os.Getenv("DB_ACCOUNT") + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to **%v** Over PostgreSQL", os.Getenv("DB_ACCOUNT"))
	fmt.Println()
	return db
}
