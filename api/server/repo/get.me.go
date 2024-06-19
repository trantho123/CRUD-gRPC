package repo

import (
	"context"
	"log"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (i *implUser) GetUser(ctx context.Context, objId primitive.ObjectID) (*model.UserDB, error) {
	var user model.UserDB
	filter := bson.M{"_id": objId}
	err := i.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("repo-GetUser1", err)
			return nil, err
		}
		log.Println("repo-GetUser2", err)
		return nil, err
	}
	return &user, nil
}
