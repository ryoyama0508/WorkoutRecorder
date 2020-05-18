package tools

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
)

//UserDataRecord is tool func for recording userdata
func UserDataRecord(
	ctx context.Context,
	db *sql.DB,
	dbname,
	userName, email, password string,
) error {
	fmt.Println(dbname)
	fmt.Println(userName)
	fmt.Println(email)
	fmt.Println(password)
	result, err := squirrel.Insert(dbname).
		Columns("email", "hashed_password", "name").
		Values(email, password, userName).
		RunWith(db).
		ExecContext(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
