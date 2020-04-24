package usecases

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/pkg/errors"
)

record

func (u *commentUpdateUseCase) Execute(ctx context.Context, input CommentUpdateUseCaseInput) (Comment, error) {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return Comment{}, errors.WithStack(ErrValidation)
	}

	comment, _, err := u.repository.CommentFind(ctx, input.CommentID)
	if err != nil {
		return Comment{}, err
	}

	if comment.UserID != input.AuthUserID {
		return Comment{}, errors.WithStack(ErrPermission)
	}

	if err := u.repository.CommentUpdate(ctx, input.CommentID, input.Content); err != nil {
		return Comment{}, err
	}

	outputUser, err := getUserOfCommentByID(ctx, u.repository, input.AuthUserID, input.AuthUserID)
	if err != nil {
		return Comment{}, err
	}

	return Comment{
		ID:        input.CommentID,
		User:      &outputUser,
		Content:   input.Content,
		Timestamp: uint(comment.CreatedAt.Unix()),
	}, nil
}
