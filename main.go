package server

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")

	router.LoadHTMLGlob("client/main/Login/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", gin.H{})
	})

	router.Run()
}
