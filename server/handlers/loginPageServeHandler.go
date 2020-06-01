package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HandleLoginPageServe ...
func HandleLoginPageServe() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	}
}
