package database

import (
	"database/sql"
	"fmt"
	"log"

	//for mysql
	_ "github.com/go-sql-driver/mysql"
)

//DBinit is used for initializing DB connection or other manipulation of database.
func DBinit(dbName string) *sql.DB {
	db, err := sql.Open("mysql", dbName)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("initialized database")
	defer db.Close()

	return db
}
