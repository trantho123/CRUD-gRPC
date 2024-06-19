package repo

import (
	"context"
	"log"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i *implUser) UpdateUser(ctx context.Context, objId primitive.ObjectID, user *model.User) error {
	data, err := bson.Marshal(user)
	if err != nil {
		return err
	}
	var doc bson.M
	err = bson.Unmarshal(data, &doc)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": doc,
	}
	_, err = i.collection.UpdateByID(ctx, objId, update)
	if err != nil {
		log.Println("repo-UpdateUser", err)
		return err
	}

	return nil
}
