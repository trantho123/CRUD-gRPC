package main

import (
	"context"
	"log"
	"time"

	pb "github.com/trantho123/CRUD-gRPC/api/pb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("client service started...")
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPostServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetPost(ctx, &pb.PostRequest{Id: "First Post"})
	if err != nil {
		log.Fatalf("Could not get post: %v", err)
	}

	log.Printf("Post: %v", resp)
}
