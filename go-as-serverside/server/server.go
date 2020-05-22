package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "practice-gRPC/go-as-serverside/pb_golang"
	"practice-gRPC/go-as-serverside/server/service"
)

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &service.UserService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
