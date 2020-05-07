package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/usecases"
)

func handleRecord(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input usecases.ExerciseRecord

		if err := decodeJSONInBody(ctx.Request, &input); err != nil {
			errors.WithStack(err)
			return
		}
		fmt.Println(&input)

		_, err := usecases.StoreAndGetData(ctx.Request.Context(), db, input)
		if err != nil {
			errors.WithStack(err)
		}
	}
}
