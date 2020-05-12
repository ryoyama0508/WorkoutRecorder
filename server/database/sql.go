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
	defer db.Close()

	_, err = db.Query("INSERT INTO users (`email`, `hashed_password`,`name`) VALUES(`test2`, `testpw`,`testname`)")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("initialized database")

	return db
}
