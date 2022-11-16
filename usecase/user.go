package usecase

import (
	"context"
	"time"

	"github.com/remiposo/gras/domain/model"
	"github.com/remiposo/gras/domain/repository"
)

type UserUc struct {
	userRepo repository.UserRepo
}

func NewUserUc(userRepo repository.UserRepo) *UserUc {
	return &UserUc{
		userRepo: userRepo,
	}
}

type UserRegisterDto struct {
	Name     string
	Password string
	Role     string
}

func (uuc *UserUc) Register(ctx context.Context, data UserRegisterDto) (string, error) {
	name, err := model.NewUserName(data.Name)
	if err != nil {
		return "", err
	}
	pw, err := model.NewUserPassword(data.Password)
	if err != nil {
		return "", err
	}
	role, err := model.NewUserRole(data.Role)
	if err != nil {
		return "", err
	}
	user := model.NewUser(name, pw, role, time.Now())
	if err := uuc.userRepo.Create(ctx, user); err != nil {
		return "", err
	}

	return string(user.ID), nil
}
