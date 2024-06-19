package repo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i *implPost) DeletePost(ctx context.Context, objId primitive.ObjectID) error {
	_, err := i.collection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		log.Println("repo-DeletePost", err)
		return err
	}
	return nil
}
