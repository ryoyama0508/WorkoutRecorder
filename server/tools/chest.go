package tools

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func ChestRecord(
	db *sql.DB,
	ctx context.Context,
	userID int, weight float32, rep, set int8,
) (*int, error) {
	if userID == 0 { //something weird
		return nil, nil
	}
	result, err := squirrel.Insert("bench_press").
		Columns("user_id", "weight", "rep", "set").
		Values(userID, weight, rep, set).
		RunWith(db).
		ExecContext(ctx)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	id := int(id64)

	return &id, nil
}
