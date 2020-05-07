package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/usecases"
)

func decodeJSONInBody(r *http.Request, d interface{}) error {
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return errors.WithStack(err)
	}

	if err := json.Unmarshal(data, d); err != nil {
		println("unmarshal error")
		return errors.WithStack(err)
	}

	return nil
}

//HandleRecord ...
func HandleRecord(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input usecases.ExerciseRecord //need userid

		if err := decodeJSONInBody(ctx.Request, &input); err != nil {
			errors.WithStack(err)
			return
		}
		fmt.Println(&input)

		_, err := usecases.StoreAndGetData(ctx.Request.Context(), db, input)
		if err != nil {
			errors.WithStack(err)
		}
	}
}
