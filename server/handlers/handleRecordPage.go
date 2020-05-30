package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//HandleRecordPage ...
func HandleRecordPage() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if session.Get("SID") == nil {
			fmt.Println("error! doesnt have sess id")
		} else {
			fmt.Println(session.Get("SID"))
			fmt.Println(session.Get("userName"))
		}
		ctx.HTML(http.StatusOK, "record.html", gin.H{"username": session.Get("userName")})
	}
}
