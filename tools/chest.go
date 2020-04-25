package tools

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func ChestRecord(
	ctx context.Context,
	userID int, weight float32, rep, set int8,
) *int {
	if userID == 0 { //something weird
		return nil
	}
	result := squirrel.Insert("bench_press").
		Columns("user_id", "weight", "rep", "set").
		Values(userID, weight, rep, set).
		RunWith().
		ExecContext(ctx)

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.WithStack(err)
	}

}
