package service

import (
	"context"
	pb "practice-gRPC/go-as-serverside/pb_golang"
)

type UserService struct{}

//Implementation code
func (u *UserService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	if req.Name == "Guobin" {
		return &pb.UserResponse{
			Id:    1,
			Name:  "Guobin",
			Age:   60,
			Title: []string{"Engineer", "Architect"},
		}, nil
	}
	return &pb.UserResponse{}, nil
}
