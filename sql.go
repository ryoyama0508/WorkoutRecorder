package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/w_recorder")
	if err != nil {
		panic(err)
	}

	fmt.Println("start database")

	defer db.Close()

	insert, err := db.Query("INSERT INTO users (email, hashed_password, name) VALUES ('test','test','ryo')")
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	fmt.Println("success")
}
