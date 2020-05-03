package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Static("/assets", "./assets")

	engine.LoadHTMLGlob("client/main/*.html")

	//html page responces
	engine.GET("/signup", func(ctx *gin.Context) { //what is that
		ctx.HTML(http.StatusOK, "signup.html", gin.H{})
	})
	engine.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})
	engine.GET("/home", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", gin.H{})
	})

	//record post and page change at the same time
	engine.GET("/record", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "record.html", gin.H{})
	})
	engine.GET("/archive", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "achieve.html", gin.H{})
	})
	engine.GET("/analysis", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "analysis.html", gin.H{})
	})

	engine.POST("/record/post", handleRecord())

	engine.Run()
}
