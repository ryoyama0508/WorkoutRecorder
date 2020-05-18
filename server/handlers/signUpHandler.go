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
		var input []usecases.HandleRecordInput

		if err := decodeJSONInBody(ctx.Request, &input); err != nil {
			errors.WithStack(err)
			return
		}

		_, err := usecases.StoreAndGetData(ctx.Request.Context(), db, input)
		if err != nil {
			errors.WithStack(err)
		}
	}
}
