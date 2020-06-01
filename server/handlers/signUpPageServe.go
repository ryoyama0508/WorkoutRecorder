package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HandleSignUpPageServe ...
func HandleSignUpPageServe() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signup.html", gin.H{})
	}
}
