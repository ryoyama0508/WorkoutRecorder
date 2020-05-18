package tools

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

//FreeWeightRecord is tool func for recording
func FreeWeightRecord(
	ctx context.Context,
	db *sql.DB,
	/* userID int, */ exercise, weightStr, repStr, setStr string,
) (*int, error) {
	/* if userID == 0 {
		return nil, nil
	} */
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
		Columns("weight", "reps", "sets").
		Values(float32(weight), reps, sets).
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
