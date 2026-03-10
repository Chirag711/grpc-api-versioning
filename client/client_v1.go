package main

import (
	"context"
	"fmt"

	pb "grpc-api-versioning/grpc-api-versioning/pb/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, _ := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	client := pb.NewUserServiceClient(conn)

	res, _ := client.GetUser(context.Background(), &pb.UserRequest{
		Id: 1,
	})

	fmt.Println("User:", res)
}
