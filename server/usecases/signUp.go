package usecases

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/tools"
)

//HandleSignUpInput is input for HandleRecord
type HandleSignUpInput struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	PassWord string `json:"passWord"`
}

//StoreDataSignUp ...
func StoreDataSignUp(ctx context.Context, db *sql.DB, data HandleSignUpInput) error {
	err := tools.UserDataRecord(
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
