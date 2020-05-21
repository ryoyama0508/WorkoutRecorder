package usecases

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"strings"

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
func UserLogin(ctx context.Context, db *sql.DB, keyValueData url.Values) (string, bool, error) {
	pw := keyValueData["passcode"]

	password := strings.Join(pw, "")

	userName := strings.Join(keyValueData["username"], "")
	if userName == "" {
		fmt.Println("username couldn't find")
	}
	email := strings.Join(keyValueData["mymail"], "")
	if email == "" {
		fmt.Println("mymail couldn't find")
	}
	passWord := string(password)
	if passWord == "" {
		fmt.Println("password couldn't find")
	}

	isCorrespond, err := tools.UserDataCheck(
		ctx,
		db,
		"users",
		userName,
		email,
		passWord,
	)
	if err != nil {
		return "", false, errors.WithStack(err)
	}

	return userName, isCorrespond, nil
}
