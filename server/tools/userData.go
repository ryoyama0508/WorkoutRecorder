package tools

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"
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
	fmt.Println(userName)
	fmt.Println(email)
	fmt.Println(password)
	var pw string
	if err := squirrel.Select("hashed_password").
		From("users").
		Where(squirrel.Eq{"name": userName, "email": email, "deleted_at": nil}).
		RunWith(db).
		QueryRowContext(ctx).Scan(&pw); err != nil {
		log.Fatal(err)
	}

	err := bcrypt.CompareHashAndPassword([]byte(pw), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}
