package repo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i *impl) DeletePost(ctx context.Context, id primitive.ObjectID) error {
	_, err := i.client.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Println("repo-DeletePost", err)
		return err
	}
	return nil
}
