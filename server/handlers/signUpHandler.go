package handlers

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/usecases"
	"golang.org/x/crypto/bcrypt"
)

//HandleSignUp ...
func HandleSignUp(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input []usecases.HandleSignUpInput

		if err := decodeJSONInBody(ctx.Request, &input); err != nil {
			errors.WithStack(err)
			return
		}

		pass, err := bcrypt.GenerateFromPassword([]byte(input[0].PassWord), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
		}

		input[0].PassWord = string(pass)

		err = usecases.StoreDataSignUp(ctx.Request.Context(), db, input[0])
		if err != nil {
			errors.WithStack(err)
		}
	}
}
