package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//HandleLogout ...
func HandleLogout() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Clear()
		session.Save()
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	}
}
