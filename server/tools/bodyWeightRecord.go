package tools

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

//BodyWeight is tool func for recording
func BodyWeightRecord(
	ctx context.Context,
	db *sql.DB,
	/* userID int, */ exercise, repStr, setStr string,
) (*int, error) {
	/* if userID == 0 {
		return nil, nil
	} */

	reps, err := strconv.Atoi(repStr)
	if err != nil {
		return nil, err
	}

	sets, err := strconv.Atoi(setStr)
	if err != nil {
		return nil, err
	}
	result, err := squirrel.Insert(exercise).
		Columns("reps", "sets").
		Values(reps, sets).
		RunWith(db).
		ExecContext(ctx)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		fmt.Println("get id error")
		return nil, errors.WithStack(err)
	}

	id := int(id64)
	return &id, nil
}
