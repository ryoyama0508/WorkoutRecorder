package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

//HandleAnalysis ...
func HandleAnalysis(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		/* session := sessions.Default(ctx)
		if session.Get("SID") == "" {
			session.Set("SID", userID)
			session.Save()
		} else {
			fmt.Println("have session ID somehow")
		}

		tmpl := template.Must(template.ParseFiles("analysis.html"))
		tmpl.Execute(ctx.Writer, cookie) */

	}
}
