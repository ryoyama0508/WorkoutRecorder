package handlers

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/usecases"
)

//HandleAnalysis ...
func HandleAnalysis(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		var userID string
		if session.Get("SID") != nil {
			encodedUserID := session.Get("SID")
			str := fmt.Sprintf("%v", encodedUserID)

			dec, err := hex.DecodeString(str)
			if err != nil {
				fmt.Println("decode error")
			}
			userID = string(dec)
			print(userID)
		} else {
			fmt.Println("have session ID somehow")
		}

		usecases.Analysis(ctx, db, userID)

		if session.Get("SID") != nil {
			ctx.HTML(http.StatusOK, "analysis.html", gin.H{"username": session.Get("username")})
		} else {
			fmt.Println("have session ID somehow")
		}
	}
}
