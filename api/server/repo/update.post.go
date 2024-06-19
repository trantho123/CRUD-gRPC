package repo

import (
	"context"
	"log"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i *implPost) UpdatePost(ctx context.Context, objId primitive.ObjectID, post *model.UpdatePost) error {
	data, err := bson.Marshal(post)
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
	log.Println("update", update)
	_, err = i.collection.UpdateByID(ctx, objId, update)
	if err != nil {
		log.Println("repo-UpdatePost", err)
		return err
	}

	return nil
}
