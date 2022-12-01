package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	connect := "user=postgres dbname=Crud-GO password=ics270699@ host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", connect)

	if err != nil {
		panic(err.Error())
	}

	return db
}
