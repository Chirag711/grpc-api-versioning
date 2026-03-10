package main

import (
	"context"
	"log"
	"net"

	pbV1 "grpc-api-versioning/grpc-api-versioning/pb/v1"
	pbV2 "grpc-api-versioning/grpc-api-versioning/pb/v2"

	"google.golang.org/grpc"
)

/*
   V1 SERVER
*/

type serverV1 struct {
	pbV1.UnimplementedUserServiceServer
}

func (s *serverV1) GetUser(ctx context.Context, req *pbV1.UserRequest) (*pbV1.UserResponse, error) {

	return &pbV1.UserResponse{
		Id:   req.Id,
		Name: "Chirag (v1)",
	}, nil
}

/*
   V2 SERVER
*/

type serverV2 struct {
	pbV2.UnimplementedUserServiceServer
}

func (s *serverV2) GetUser(ctx context.Context, req *pbV2.UserRequest) (*pbV2.UserResponse, error) {

	return &pbV2.UserResponse{
		Id:    req.Id,
		Name:  "Chirag",
		Email: "chirag@email.com",
	}, nil
}

/*
   MAIN
*/

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pbV1.RegisterUserServiceServer(grpcServer, &serverV1{})
	pbV2.RegisterUserServiceServer(grpcServer, &serverV2{})

	log.Println("gRPC API Versioning Server running on port 50051")

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatal(err)
	}
}
