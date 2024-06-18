package repo

import (
	"context"
	"log"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (i *impl) GetPost(ctx context.Context, id primitive.ObjectID) (*model.DBPost, error) {
	var post model.DBPost
	filter := bson.M{"_id": id}
	err := i.client.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("repo-GetPost1", err)
			return nil, err
		}
		log.Println("repo-GetPost2", err)
		return nil, err
	}

	return &post, nil
}
