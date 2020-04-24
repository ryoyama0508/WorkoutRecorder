package handler

import "github.com/gin-gonic/gin"

func recordHundler(ctx *gin.Context) {
	ctx.HTML(200, "record.html", "")

}
