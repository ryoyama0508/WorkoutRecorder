package handlers

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/usecases"
)

//HandleLogin ...
func HandleLogin(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Request.ParseForm()

		keyValueData := ctx.Request.Form

		userName, isCorrespond, userID, err := usecases.UserLogin(ctx.Request.Context(), db, keyValueData)
		if err != nil {
			errors.WithStack(err)
		}

		if isCorrespond == true {
			session := sessions.Default(ctx)

			if session.Get("SID") == nil {
				userIDStr := strconv.FormatUint(uint64(userID), 10)
				hx := hex.EncodeToString([]byte(userIDStr))
				session.Set("SID", hx)
				session.Set("username", userName)
				session.Save()
			} else {
				fmt.Println("have session ID somehow")
			}

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
