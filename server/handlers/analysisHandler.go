package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//HandleAnalysis ...
func HandleAnalysis(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		if session.Get("SID") != nil {
			ctx.HTML(http.StatusOK, "analysis.html", gin.H{"username": session.Get("username")})
		} else {
			fmt.Println("have session ID somehow")
		}
	}
}
