package database

import (
	"database/sql"
	"fmt"

	//for mysql
	_ "github.com/go-sql-driver/mysql"
)

//DBinit is used for initializing DB connection or other manipulation of database.
func DBinit(dbName string) *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/w_recorder")
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Query("INSERT INTO bench_press (user_id, weight, rep, set)VALUES (1, 50, 5, 5)")
	if err != nil {
		fmt.Println("experient")
		fmt.Println(err)
	}

	fmt.Println("initialized database")

	return db
}
