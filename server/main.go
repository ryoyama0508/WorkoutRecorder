package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/database"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/handlers"
)

func main() {
	databaseName := "root:@tcp(127.0.0.1:3306)/w_recorder"

	var db *sql.DB
	db = database.DBinit(databaseName)

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

	engine.POST("/record/post", handlers.HandleRecord(db))

	engine.Run()

}
