package gapi

import (
	"context"
	"errors"
	"time"

	pb "github.com/trantho123/CRUD-gRPC/api/pb"
	"github.com/trantho123/CRUD-gRPC/api/server/controller"
	model "github.com/trantho123/CRUD-gRPC/api/server/models"
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
	if in == nil || in.GetId() == "" {
		return nil, errors.New("request is invalid")
	}
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
	var page = in.GetPage()
	var limit = in.GetLimit()
	if page != 0 || limit != 0 {
		posts, err := s.i.GetPosts(stream.Context(), int(page), int(limit))
		if err != nil {
			return err
		}
		for _, post := range *posts {
			response := &pb.PostResponse{
				Post: &pb.Post{
					Id:        post.Id.Hex(),
					Title:     post.Title,
					Content:   post.Content,
					Image:     post.Image,
					User:      post.User,
					CreatedAt: timestamppb.New(post.CreateAt),
					UpdatedAt: timestamppb.New(post.UpdatedAt),
				},
			}
			if err := stream.Send(response.Post); err != nil {
				return errors.New("failed to send response")
			}
		}
	}

	return nil
}

// CreatePost implements pb.PostServiceServer
func (s *PostServiceServer) CreatePost(ctx context.Context, in *pb.CreatePostRequest) (*pb.PostResponse, error) {
	newPost := &model.CreatePostRequest{
		Title:     in.GetTitle(),
		Content:   in.GetContent(),
		Image:     in.GetImage(),
		User:      in.GetUser(),
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.i.CreatePost(ctx, newPost)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdatePost implements pb.PostServiceServer
func (s *PostServiceServer) UpdatePost(ctx context.Context, in *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	updatePost := &model.UpdatePost{
		Title:     in.GetTitle(),
		Content:   in.GetContent(),
		Image:     in.GetImage(),
		User:      in.GetUser(),
		UpdatedAt: time.Now(),
	}
	err := s.i.UpdatePostById(ctx, in.Id, updatePost)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// DeletePost implements pb.PostServiceServer
func (s *PostServiceServer) DeletePost(ctx context.Context, in *pb.PostRequest) (*pb.DeletePostResponse, error) {
	resp := &pb.DeletePostResponse{
		Success: false,
	}
	if in == nil || in.GetId() == "" {
		return resp, errors.New("request is invalid")
	}
	err := s.i.DeletePostById(ctx, in.Id)
	if err != nil {
		return resp, err
	}
	resp.Success = true
	return resp, nil
}
