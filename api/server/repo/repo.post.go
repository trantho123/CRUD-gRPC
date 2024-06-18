package repo

import (
	"context"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Registry interface {
	GetPost(ctx context.Context, id primitive.ObjectID) (*model.DBPost, error)
	GetPosts(ctx context.Context, page int, limit int) (*[]model.DBPost, error)
	CreatePost(ctx context.Context, post *model.CreatePostRequest) error
	UpdatePost(ctx context.Context, id primitive.ObjectID, post *model.UpdatePost) error
	DeletePost(ctx context.Context, id primitive.ObjectID) error
}
type impl struct {
	client *mongo.Collection
}

func New(client *mongo.Database) Registry {
	return &impl{client: client.Collection("posts")}
}
