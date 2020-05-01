package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")

	router.LoadHTMLGlob("client/main/*.html")

	//html page responces
	router.GET("/signup", func(ctx *gin.Context) { ctx.HTML(200, "signup.html", "") })
	router.GET("/login", func(ctx *gin.Context) { ctx.HTML(200, "login.html", "") })
	router.GET("/home", func(ctx *gin.Context) { ctx.HTML(200, "home.html", "") })

	//record post and page change at the same time
	router.GET("/record", func(ctx *gin.Context) { ctx.HTML(200, "record.html", "") })
	router.GET("/archive", func(ctx *gin.Context) { ctx.HTML(200, "archive.html", "") })
	router.GET("/analysis", func(ctx *gin.Context) { ctx.HTML(200, "analysis.html", "") })

	router.POST("/record/post", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbInsert(text, status)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
