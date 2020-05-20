package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/usecases"
)

//HandleLogin ...
func HandleLogin(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Request.ParseForm()

		fmt.Println(ctx.Request.Form)

		keyValueData := ctx.Request.Form

		isCorrespond, err := usecases.UserLogin(ctx.Request.Context(), db, keyValueData)
		if err != nil {
			errors.WithStack(err)
		}
		if isCorrespond == true {
			ctx.HTML(http.StatusOK, "home.html", gin.H{})
		}
	}
}
