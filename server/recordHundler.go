package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/usecases"
)

func handleRecord() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input usecases.Exercise

		if err := decodeJSONnInBody(ctx.Request, &input); err != nil {
			errors.WithStack(err)
			return
		}

		usecases.StoreAndGetData(ctx.Request.Context(), input) //what is that?
	}
}
