package tools

import (
	"context"
	"database/sql"

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
		Columns("userName", "email", "password").
		Values(userName, email, password).
		RunWith(db).
		ExecContext(ctx)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
