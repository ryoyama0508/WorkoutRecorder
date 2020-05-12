package tools

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

//ChestRecord is tool func for recording
func ChestRecord(
	ctx context.Context,
	db *sql.DB,
	/* userID int, */ weight, rep, set string,
) (*int, error) {
	/* if userID == 0 { //something weird
		return nil, nil
	} */
	result, err := squirrel.Insert("bench_press").
		Columns("weight", "rep", "set").
		Values(weight, rep, set).
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
