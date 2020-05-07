package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

//DBinit is used for initializing DB connection or other manipulation of database.
func DBinit(dbName string) *sql.DB {
	db, err := sql.Open("mysql", dbName)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("initialized database")

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	http.ListenAndServe(":8080", nil)

	defer db.Close()

	return db
}
