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
) ([]uint, error) {
	if userIDStr == "" {
		return nil, nil
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, err
	}
	queryRow := squirrel.Select("reps", "created_at").
		From(exercise).
		Where(squirrel.Eq{"user_id": userID, "deleted_at": nil})

	row, err := queryRow.OrderBy("created_at desc").
		RunWith(db).
		QueryContext(ctx)

	if err != nil {
		return nil, err
	}

	rec := make([][]uint, 8)
	i := 0
	j := 0
	var start time.Time
	for row.Next() {
		var reps uint
		var createdAtRaw rawTime
		if err := row.Scan(
			&reps,
			&createdAtRaw,
		); err != nil {
			return nil, err
		}

		createdAt, err := createdAtRaw.Time()
		if err != nil {
			panic(err)
		}

		if j == 0 {
			start = time.Now()
		}

		if createdAt.Before(start.Add((-24 * 10) * time.Hour)) {
			start = start.Add((-24 * 10) * time.Hour)
			i++
		}

		rec[i] = append(rec[i], reps)
	}

	fmt.Println(rec, "rec")

	avg := make([]uint, 8)

	for i = 0; i < len(rec); i++ {
		var sum uint
		for j = 0; j < len(rec[i]); j++ {
			fmt.Println(rec[i][j], "print")
			sum += rec[i][j]
		}
		fmt.Println(sum, "sum")
		if len(rec[i]) != 0 {
			avg[i] = sum / uint(len(rec[i]))
		} else {
			avg[i] = 0
		}
	}

	fmt.Println(avg, "avg")

	return avg, nil
}
