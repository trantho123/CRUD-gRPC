package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/trantho123/CRUD-gRPC/api/pb"
	controller "github.com/trantho123/CRUD-gRPC/api/server/controller"
	gapi "github.com/trantho123/CRUD-gRPC/api/server/gapi"
	repo "github.com/trantho123/CRUD-gRPC/api/server/repo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	println("server service started...")
	config := LoadEnv()
	con, err := ConnectToDatabase(config)
	if err != nil {
		log.Fatal(err)
	} else {
		println("Connected to database")
	}
	repo := repo.New(con.Collection("posts"))
	ctr := controller.New(repo)
	lis, err := net.Listen("tcp", ":"+config["PORT"])
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	postServiceServer := gapi.NewPostServiceServer(ctr)
	pb.RegisterPostServiceServer(s, postServiceServer)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func ConnectToDatabase(config map[string]string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(config["MONGO_URI"])
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(config["MONGO_DB"]), nil

}

func LoadEnv() map[string]string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
	config := make(map[string]string)
	config["MONGO_URI"] = os.Getenv("MONGO_URI")
	log.Println(config["MONGO_URI"])
	config["MONGO_USERNAME"] = os.Getenv("MONGO_USERNAME")
	config["MONGO_DB"] = os.Getenv("MONGO_DB")
	config["PORT"] = os.Getenv("PORT")
	return config
}
