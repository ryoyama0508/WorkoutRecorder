package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/database"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/handlers"
)

func main() {
	databaseName := "root:@tcp(127.0.0.1:3306)/w_recorder"

	var db *sql.DB
	db = database.DBinit(databaseName)
	fmt.Println("finish db")

	engine := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("mysession", store))

	engine.Static("/assets", "./assets")

	engine.LoadHTMLGlob("../client/templates/*.html")

	//html page responces
	engine.GET("/signup", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signup.html", gin.H{})
	})
	engine.POST("/signup/post", handlers.HandleSignUp(db))

	engine.GET("/login/see", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})
	engine.POST("/login/enter", handlers.HandleLogin(db))

	engine.GET("/logout", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Clear()
		session.Save()
		ctx.String(http.StatusOK, "log out")
	})

	engine.GET("/home", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", gin.H{})
	})

	//record post and page change at the same time
	engine.GET("/record/see", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if session.Get("SID") == nil {
			fmt.Println("error! doesnt have sess id")
		} else {
			fmt.Println(session.Get("SID"))
			fmt.Println(session.Get("userName"))
		}
		ctx.HTML(http.StatusOK, "record.html", gin.H{"username": session.Get("userName")})
	})
	engine.GET("/archive", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "archive.html", gin.H{})
	})
	engine.GET("/analysis", handlers.HandleAnalysis(db))

	engine.POST("/record/post", handlers.HandleRecord(db))

	engine.Run(":8080")

}
