package usecases

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/tools"
)

//HandleLoginInput is input for HandleRecord
type HandleLoginInput struct {
	UserName string `json:"userName"`
	Email    string `json:"mailAddr"`
	PassWord string `json:"passWord"`
}

//UserLogin ...
func UserLogin(ctx context.Context, db *sql.DB, data HandleLoginInput) error {
	fmt.Println(data.UserName)
	_, err := tools.UserDataCheck(
		ctx,
		db,
		"users",
		data.UserName,
		data.Email,
		data.PassWord,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
