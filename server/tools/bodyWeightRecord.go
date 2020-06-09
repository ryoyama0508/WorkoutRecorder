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

type rawTime []byte

func (t rawTime) Time() (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", string(t))
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
	queryRow := squirrel.Select("weight", "reps", "created_at").
		From(exercise).
		Where(squirrel.Eq{"user_id": userID, "deleted_at": nil})

	row, err := queryRow.OrderBy("created_at desc").
		RunWith(db).
		QueryContext(ctx)

	if err != nil {
		return err
	}

	rec := make([][]float64, 8)
	i := 0
	j := 0
	var start time.Time
	for row.Next() {
		var weight uint
		var reps uint
		var createdAtRaw rawTime
		if err := row.Scan(
			&weight,
			&reps,
			&createdAtRaw,
		); err != nil {
			return err
		}

		oneRM := float64(weight) * (1 + (float64(reps) / 30))

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

		rec[i] = append(rec[i], oneRM)
	}

	fmt.Println(rec, "output")

	return nil
}
