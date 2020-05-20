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

//UserDataCheck is tool func for checking userdata
func UserDataCheck(
	ctx context.Context,
	db *sql.DB,
	dbname,
	userName, email, password string,
) (bool, error) {
	_, err := squirrel.Select("name", "email", "hashed_password").
		From("users").
		Where(squirrel.Eq{"name": userName, "email": email, "hashed_password": password, "deleted_at": nil}).
		RunWith(db).
		ExecContext(ctx)

	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return false, err
	}
	if err == sql.ErrNoRows {
		fmt.Println("No rows")
		return false, err
	}
	fmt.Println("correspond")

	return true, nil
}
