package repository

import (
	"context"

	"github.com/remiposo/gras/domain/model"
)

type UserRepo interface {
	Create(ctx context.Context, user *model.User) error
}
