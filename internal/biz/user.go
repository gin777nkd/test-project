package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
)

type User struct {
	ID    int32
	Name  string
	Email string
}

type UserRepo interface {
	CreateUser(ctx context.Context, name, email string) (int64, error)
	GetUser(ctx context.Context, id int64) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) Register(ctx context.Context, name, email string) (int64, error) {
	users, _ := uc.repo.ListUsers(ctx)
	for _, user := range users {
		if user.Email == email {
			return 0, errors.New(409, "USER_EXISTS", "user with this email already exists")
		}
	}
	return uc.repo.CreateUser(ctx, name, email)
}

func (uc *UserUsecase) Get(ctx context.Context, id int64) (User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUsecase) List(ctx context.Context) ([]User, error) {
	return uc.repo.ListUsers(ctx)
}
