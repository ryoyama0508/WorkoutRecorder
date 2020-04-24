package handler

import "github.com/gin-gonic/gin"

func recordHandler(ctx *gin.Context) {
	ctx.HTML(200, "record.html", "")

}
