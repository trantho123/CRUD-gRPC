package gapi

import (
	"context"

	pb "github.com/trantho123/CRUD-gRPC/api/pb"
	"github.com/trantho123/CRUD-gRPC/api/server/controller"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostServiceServer struct {
	pb.UnimplementedPostServiceServer
	i controller.Controller
}

func NewPostServiceServer(i controller.Controller) *PostServiceServer {
	return &PostServiceServer{i: i}
}

// GetPost implements pb.PostServiceServer
func (s *PostServiceServer) GetPost(ctx context.Context, in *pb.PostRequest) (*pb.PostResponse, error) {
	post, err := s.i.GetPost(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.PostResponse{
		Post: &pb.Post{
			Id:        post.Id.Hex(),
			Title:     post.Title,
			Content:   post.Content,
			Image:     post.Image,
			User:      post.User,
			CreatedAt: timestamppb.New(post.CreateAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
		}}, nil
}

// GetPosts implements pb.PostServiceServer
func (s *PostServiceServer) GetPosts(in *pb.GetPostsRequest, stream pb.PostService_GetPostsServer) error {
	return nil
}

// CreatePost implements pb.PostServiceServer
func (s *PostServiceServer) CreatePost(ctx context.Context, in *pb.CreatePostRequest) (*pb.PostResponse, error) {
	return nil, nil
}

// UpdatePost implements pb.PostServiceServer
func (s *PostServiceServer) UpdatePost(ctx context.Context, in *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	return nil, nil
}

// DeletePost implements pb.PostServiceServer
func (s *PostServiceServer) DeletePost(ctx context.Context, in *pb.PostRequest) (*pb.DeletePostResponse, error) {
	return nil, nil
}
