package service

import (
	"context"
	pb "my-project/api/user/v1"
	"my-project/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	id, err := s.uc.Register(ctx, req.Name, req.Email)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Id: id,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Id:    int64(user.ID),
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersReply, error) {
	users, err := s.uc.List(ctx)
	if err != nil {
		return nil, err
	}
	var pbUsers []*pb.GetUserReply
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.GetUserReply{
			Id:    int64(user.ID),
			Email: user.Email,
			Name:  user.Name,
		})
	}
	return &pb.ListUsersReply{Users: pbUsers}, nil
}
