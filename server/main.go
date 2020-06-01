package main

import (
	"database/sql"
	"fmt"

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

	engine.GET("/signup", handlers.HandleSignUpPageServe())
	engine.POST("/signup/post", handlers.HandleSignUp(db))

	engine.GET("/login", handlers.HandleLoginPageServe())
	engine.POST("/login/enter", handlers.HandleLogin(db))

	engine.GET("/logout", handlers.HandleLogout())

	engine.GET("/home", handlers.HandleHome())

	engine.GET("/record", handlers.HandleRecordPage())
	engine.POST("/record/post", handlers.HandleRecord(db))

	engine.GET("/archive", handlers.HandleArchive())

	engine.GET("/analysis", handlers.HandleAnalysis(db))

	engine.Run(":8080")

}
