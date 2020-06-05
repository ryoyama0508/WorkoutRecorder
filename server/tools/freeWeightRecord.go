package tools

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

//FreeWeightRecordInsert is tool func for recording
func FreeWeightRecordInsert(
	ctx context.Context,
	db *sql.DB,
	userIDStr, exercise, weightStr, repStr, setStr string,
) (*int, error) {
	if userIDStr == "" {
		return nil, nil
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, err
	}
	weight, err := strconv.ParseFloat(weightStr, 32)
	if err != nil {
		fmt.Println("float error")
	}

	reps, err := strconv.Atoi(repStr)
	if err != nil {
		return nil, err
	}

	sets, err := strconv.Atoi(setStr)
	if err != nil {
		return nil, err
	}
	result, err := squirrel.Insert(exercise).
		Columns("user_id", "weight", "reps", "sets").
		Values(userID, float32(weight), reps, sets).
		RunWith(db).
		ExecContext(ctx)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	id := int(id64)
	return &id, nil
}
