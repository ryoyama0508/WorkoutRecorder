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

//HandleRecordInput is input for HandleRecord
type HandleRecordInput struct {
	Weight string `json:"weight"`
	Reps   string `json:"reps"`
	Sets   string `json:"sets"`
}

func decodeJSONInBody(r *http.Request, d interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := json.Unmarshal(data, d); err != nil {
		fmt.Println(err)
	}

	return nil
}

//HandleRecord ...
func HandleRecord(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input []HandleRecordInput //need userid

		if err := decodeJSONInBody(ctx.Request, &input); err != nil {
			errors.WithStack(err)
			return
		}

		_, err := usecases.StoreAndGetData(ctx.Request.Context(), db, input[0].Weight, input[0].Reps, input[0].Sets)
		if err != nil {
			errors.WithStack(err)
		}
	}
}
