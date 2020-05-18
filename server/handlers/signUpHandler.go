package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/usecases"
)

//HandleSignUp ...
func HandleSignUp(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input []usecases.HandleSignUpInput

		if err := decodeJSONInBody(ctx.Request, &input); err != nil {
			errors.WithStack(err)
			return
		}

		err := usecases.StoreDataSignUp(ctx.Request.Context(), db, input[0])
		if err != nil {
			errors.WithStack(err)
		}
	}
}
