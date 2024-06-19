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

	clientPostService := pb.NewPostServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	clientUserService := pb.NewUserServiceClient(conn)
	a := "66728cd690db585bd1a26a1a"
	b := "name"
	c := "email"
	d := "aaaaa"
	ketqua, err := clientUserService.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:       a,
		Name:     &b,
		Email:    &c,
		Password: &d,
	})
	if err != nil {
		log.Fatalf("Could not get user: %v", err)
	} else {
		log.Println("Update user: Susscess", ketqua)
	}

	resp, err := clientPostService.GetPost(ctx, &pb.PostRequest{Id: "66728cd690db585bd1a26a13"})
	if err != nil {
		log.Fatalf("Could not get post: %v", err)
	}
	log.Println("Get user: ", resp)
	// valPage, err := strconv.ParseInt("1", 10, 64)
	// if err != nil {
	// 	// Xử lý lỗi
	// 	panic(err)
	// }
	// valLimit, err := strconv.ParseInt("10", 10, 64)
	// if err != nil {
	// 	// Xử lý lỗi
	// 	panic(err)
	// }
	// stream, err := clientPostService.GetPosts(ctx, &pb.GetPostsRequest{Page: &valPage, Limit: &valLimit})
	// if err != nil {
	// 	log.Fatalf("Could not get post: %v", err)
	// }
	// for {
	// 	// Đọc một phần tử phản hồi từ stream
	// 	resp, err := stream.Recv()
	// 	if err == io.EOF {
	// 		// Khi stream kết thúc, không còn phản hồi nào nữa
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatalf("Error while reading stream: %v", err)
	// 	}

	// 	// Xử lý phản hồi được nhận
	// 	// Ví dụ: In thông tin của bài đăng
	// 	log.Printf("Received post: %v", resp)
	// }

	// stream, err := clientPostService.CreatePost(ctx, &pb.CreatePostRequest{
	// 	Title:   "title",
	// 	Content: "content",
	// 	Image:   "image",
	// 	User:    "user"})
	// if err != nil {
	// 	log.Fatalf("Could not get post: %v", err)
	// }
	// log.Println("Create post: ", stream)

	// stream, err := clientPostService.DeletePost(ctx, &pb.PostRequest{Id: "66716211dde700a9ef686a23"})
	// if err != nil {
	// 	log.Fatalf("Could not get post: %v", err)
	// }
	// log.Println("Delete post: ", stream)

	// title := "Updated Title"
	// content := "abc xyz"
	// image := "updated_image.jpg"
	// user := "user1"
	// stream, err := clientPostService.UpdatePost(ctx, &pb.UpdatePostRequest{
	// 	Id:      "667182493d6fc26603a26a14",
	// 	Title:   &title,
	// 	Content: &content,
	// 	Image:   &image,
	// 	User:    &user})
	// if err != nil {
	// 	log.Fatalf("Could not get post: %v", err)
	// }
	// log.Println(" post: ", stream)
}
