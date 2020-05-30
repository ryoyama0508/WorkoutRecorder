package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//HandleHome ...
func HandleHome() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if session.Get("username") == nil {
			fmt.Println("cannot get username")
		} else {
			ctx.HTML(http.StatusOK, "home.html", gin.H{
				"username": session.Get("username"),
			})
		}
	}
}
