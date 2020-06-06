package tools

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

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

type BodyWeightRecordGetOutput struct {
	weight *[]int8
	reps   *[]int8
}

// BodyWeightRecordGet is ...
func BodyWeightRecordGet(
	ctx context.Context,
	db *sql.DB,
	userIDStr, exercise string,
) (*int, error) {
	if userIDStr == "" {
		return nil, nil
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, err
	}
	output := &BodyWeightRecordGetOutput{}
	rowScanner := squirrel.Select("weight", "reps").
		From(exercise).
		Where(squirrel.Eq{"user_id": userID, "deleted_at": nil}).
		RunWith(db).
		QueryRow()

	if err := rowScanner.Scan(
		&output.weight,
		&output.reps,
	); err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(output)

	return nil, nil
}
