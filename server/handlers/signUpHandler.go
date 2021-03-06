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
		ctx.Request.ParseForm()

		keyValueData := ctx.Request.Form

		err := usecases.StoreDataSignUp(ctx.Request.Context(), db, keyValueData)
		if err != nil {
			errors.WithStack(err)
		}
	}
}
