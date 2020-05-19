package handlers

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/usecases"
	"golang.org/x/crypto/bcrypt"
)

//HandleLogin ...
func HandleLogin(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input []usecases.HandleLoginInput

		if err := decodeJSONInBody(ctx.Request, &input); err != nil {
			errors.WithStack(err)
			return
		}
		fmt.Println("here")
		pass, err := bcrypt.GenerateFromPassword([]byte(input[0].PassWord), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			// err := ErrorResponse{
			// 	Err: "Password Encryption  failed",
			// }
			// json.NewEncoder(w).Encode(err)
		}

		input[0].PassWord = string(pass)

		err = usecases.UserLogin(ctx.Request.Context(), db, input[0])
		if err != nil {
			errors.WithStack(err)
		}
	}
}
