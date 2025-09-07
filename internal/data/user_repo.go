package data

import (
	"context"
	"my-project/internal/biz"
	"my-project/internal/data/sqlc"
)

type UserRepo struct {
	queries *sqlc.Queries
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &UserRepo{
		queries: sqlc.New(data.DB),
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, name, email string) (int64, error) {
	params := sqlc.CreateUserParams{
		Name:  name,
		Email: email,
	}
	id, err := r.queries.CreateUser(ctx, params)
	return id, err
}

func (r *UserRepo) GetUser(ctx context.Context, id int64) (biz.User, error) {
	user, err := r.queries.GetUser(ctx, int32(id))
	if err != nil {
		return biz.User{}, err
	}
	return biz.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (r *UserRepo) ListUsers(ctx context.Context) ([]biz.User, error) {
	users, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	var result []biz.User
	for _, u := range users {
		result = append(result, biz.User{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}
	return result, nil
}
