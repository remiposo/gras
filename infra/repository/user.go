package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/remiposo/gras/domain/model"
	"github.com/remiposo/gras/infra/dao"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) Create(ctx context.Context, user *model.User) error {
	userDao := &dao.User{
		ID:        string(user.ID),
		Name:      string(user.Name),
		Password:  string(user.Password),
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt,
	}
	sql := `INSERT INTO users (id, name, password, role, created_at)
          VALUES (:id, :name, :password, :role, :created_at);`
	_, err := ur.db.NamedExecContext(ctx, sql, userDao)
	if err != nil {
		return err
	}
	return nil

}
