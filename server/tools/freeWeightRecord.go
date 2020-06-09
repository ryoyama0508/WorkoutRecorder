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

type rawTime []byte

func (t rawTime) Time() (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", string(t))
}

// FreeWeightRecordGet is ...
func FreeWeightRecordGet(
	ctx context.Context,
	db *sql.DB,
	userIDStr, exercise string,
) ([]float64, error) {
	if userIDStr == "" {
		return nil, nil
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, err
	}
	queryRow := squirrel.Select("weight", "reps", "created_at").
		From(exercise).
		Where(squirrel.Eq{"user_id": userID, "deleted_at": nil})

	row, err := queryRow.OrderBy("created_at desc").
		RunWith(db).
		QueryContext(ctx)

	if err != nil {
		return nil, err
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
			return nil, err
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
	fmt.Println(rec, "rec")

	avg := make([]float64, 8)

	for i = 0; i < len(rec); i++ {
		var sum float64
		for j = 0; j < len(rec[i]); j++ {
			sum += rec[i][j]
		}
		fmt.Println(sum, "sum")
		if len(rec[i]) != 0 {
			avg[i] = sum / float64(len(rec[i]))
		} else {
			avg[i] = 0
		}
	}

	fmt.Println(avg, "avg")

	return avg, nil
}
