package controller

import (
	"context"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"github.com/trantho123/CRUD-gRPC/api/server/repo"
)

type PostController interface {
	GetPost(ctx context.Context, id string) (*model.DBPost, error)
	GetPosts(ctx context.Context, page int, limit int) (*[]model.DBPost, error)
	CreatePost(ctx context.Context, post *model.CreatePostRequest) error
	UpdatePostById(ctx context.Context, id string, post *model.UpdatePost) error
	DeletePostById(ctx context.Context, id string) error
}

type implPost struct {
	repo repo.RegistryPost
}

func NewPostController(repo repo.RegistryPost) PostController {
	return &implPost{repo: repo}
}
