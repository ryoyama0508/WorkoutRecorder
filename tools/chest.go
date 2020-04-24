package tools

import (
	"context"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func ChestRecord(
	ctx context.Context,
	userID int, weight float32, rep, set int8,
) (string, error) {
	postIDInInt, err := strconv.Atoi(postID)
	if err != nil {
		return "", err
	}

	insert := squirrel.Insert("goals")
	if reflectionID != "" {
		reflectionIDInInt, err := strconv.Atoi(reflectionID)
		if err != nil {
			return "", err
		}
		insert = insert.Columns("post_id", "what", "how_much", "how", "reflection_id").
			Values(postIDInInt, what, howMuch, how, reflectionIDInInt)
	} else {
		insert = insert.Columns("post_id", "what", "how_much", "how").
			Values(postIDInInt, what, howMuch, how)
	}
	result, err := insert.RunWith(r.DB).
		ExecContext(ctx)

	if err != nil {
		return "", errors.WithStack(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", errors.WithStack(err)
	}

	return strconv.Itoa(int(id)), nil
}
