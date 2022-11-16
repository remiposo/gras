package repository

import (
	"context"

	"github.com/remiposo/gras/domain/model"
)

type TaskRepo interface {
	Create(ctx context.Context, task *model.Task) error
	List(ctx context.Context) (model.Tasks, error)
}
