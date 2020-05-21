package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/usecases"
)

//HandleLogin ...
func HandleLogin(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Request.ParseForm()

		keyValueData := ctx.Request.Form

		userName, isCorrespond, err := usecases.UserLogin(ctx.Request.Context(), db, keyValueData)
		if err != nil {
			errors.WithStack(err)
		}
		if isCorrespond == true {
			ctx.HTML(http.StatusOK, "home.html", gin.H{
				"username": userName,
			})
		} else {
			ctx.HTML(http.StatusOK, "login.html", gin.H{
				"message": "your input data is incorrect",
			})
		}
	}
}
