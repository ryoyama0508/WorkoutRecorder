package tools

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

//BodyWeightRecordInsert is ...
func BodyWeightRecordInsert(
	ctx context.Context,
	db *sql.DB,
	userIDStr, exercise, repStr, setStr string,
) (*int, error) {
	if userIDStr == "" {
		return nil, nil
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, err
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
		Columns("user_id", "reps", "sets").
		Values(userID, reps, sets).
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

// BodyWeightRecordGet is ...
func BodyWeightRecordGet(
	ctx context.Context,
	db *sql.DB,
	userIDStr, exercise string,
) error {
	if userIDStr == "" {
		return nil
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}
	output := make(map[float64]time.Time)
	queryRow := squirrel.Select("weight", "reps", "created_at").
		From(exercise).
		Where(squirrel.Eq{"user_id": userID, "deleted_at": nil})

	row, err := queryRow.OrderBy("created_at desc").
		RunWith(db).
		QueryContext(ctx)

	if err != nil {
		return err
	}

	for row.Next() {
		var weight uint
		var reps uint
		var createdAt time.Time
		if err := row.Scan(
			&weight,
			&reps,
			&createdAt,
		); err != nil {
			return errors.WithStack(err)
		}
		max := float64(weight) * (1 + (float64(reps) / 30))

		//append into map
	}

	fmt.Println(output, "output")

	return nil
}
