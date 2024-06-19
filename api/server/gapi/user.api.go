package gapi

import (
	"context"

	pb "github.com/trantho123/CRUD-gRPC/api/pb"
	"github.com/trantho123/CRUD-gRPC/api/server/controller"
	model "github.com/trantho123/CRUD-gRPC/api/server/models"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	i controller.UserController
}

func NewUserServiceServer(i controller.UserController) *UserServiceServer {
	return &UserServiceServer{i: i}
}

// GetUser implements pb.UserServiceServer
func (s *UserServiceServer) GetMe(ctx context.Context, in *pb.GetMeRequest) (*pb.UserResponse, error) {
	resp, err := s.i.GetUser(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		User: &pb.User{
			Id:       resp.ID,
			Name:     resp.Username,
			Email:    resp.Email,
			Password: resp.Password,
		}}, nil
}

// GetUser implements pb.UserServiceServer

func (s *UserServiceServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := &model.User{
		Username: in.GetName(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}
	err := s.i.UpdateUser(ctx, in.GetId(), user)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
