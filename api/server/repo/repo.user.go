package repo

import (
	"context"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegistryUser interface {
	GetUser(ctx context.Context, objId primitive.ObjectID) (*model.UserDB, error)
	UpdateUser(ctx context.Context, objId primitive.ObjectID, user *model.User) error
}
type implUser struct {
	collection *mongo.Collection
}

func NewUserCollection(collection *mongo.Database) RegistryUser {
	return &implUser{collection: collection.Collection("users")}
}
