package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "practice-gRPC/go-as-serverside/pb_golang"
)

func main() {
	conn, err := grpc.Dial(":9999", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)
	resp, _ := client.GetUserInfo(context.Background(), &pb.UserRequest{
		Name: "Guobin",
	})
	log.Printf("Received user name: %v", resp.Name)
}
