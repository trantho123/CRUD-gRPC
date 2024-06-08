package repo

import (
	"context"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Registry interface {
	GetPost(ctx context.Context, id string) (*model.DBPost, error)
}
type impl struct {
	client *mongo.Collection
}

func New(client *mongo.Collection) Registry {
	return &impl{client: client}
}
