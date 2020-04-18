package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("start database")

	db, err := sql.Open("mysql", "")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("success")
}
