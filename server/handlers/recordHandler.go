package handlers

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
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
		fmt.Println("unmarshal")
		fmt.Println(err)
	}

	return nil
}

//HandleRecord ...
func HandleRecord(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input []usecases.HandleRecordInput

		if err := decodeJSONInBody(ctx.Request, &input); err != nil {
			errors.WithStack(err)
			return
		}

		session := sessions.Default(ctx)

		if session.Get("SID") != nil {
			encodedUserID := fmt.Sprintf("%v", session.Get("SID"))
			userID, err := hex.DecodeString(encodedUserID)
			if err != nil {
				fmt.Println("failed decode UserName")
			}

			_, err = usecases.StoreAndGetData(
				ctx.Request.Context(), db, string(userID), input,
			)
			if err != nil {
				errors.WithStack(err)
			}

		} else {
			fmt.Println("doesnt have session ID")
		}

	}
}
