package database

import (
	"database/sql"
	"fmt"

	//for mysql
	_ "github.com/go-sql-driver/mysql"
)

//DBinit is used for initializing DB connection or other manipulation of database.
func DBinit(dbName string) *sql.DB {
	db, err := sql.Open("mysql", dbName)
	if err != nil {
		fmt.Println(err)
	}

	return db
}
