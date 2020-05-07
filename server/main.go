package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func main() {
	databaseName := "root:@tcp(127.0.0.1:3306)/w_recorder"

	var db *sql.DB
	db = dbInit(databaseName)

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

	engine.POST("/record/post", handleRecord(db))

	engine.Run()

}

func decodeJSONInBody(r *http.Request, d interface{}) error {
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return errors.WithStack(err)
	}

	if err := json.Unmarshal(data, d); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
