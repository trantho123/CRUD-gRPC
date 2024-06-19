package repo

import (
	"context"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegistryPost interface {
	GetPost(ctx context.Context, objId primitive.ObjectID) (*model.DBPost, error)
	GetPosts(ctx context.Context, page int, limit int) (*[]model.DBPost, error)
	CreatePost(ctx context.Context, post *model.CreatePostRequest) error
	UpdatePost(ctx context.Context, objId primitive.ObjectID, post *model.UpdatePost) error
	DeletePost(ctx context.Context, objId primitive.ObjectID) error
}
type implPost struct {
	collection *mongo.Collection
}

func NewPostCollection(collection *mongo.Database) RegistryPost {
	return &implPost{collection: collection.Collection("posts")}
}
