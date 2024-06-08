package repo

import (
	"context"
	"fmt"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (i *impl) GetPost(ctx context.Context, title string) (*model.DBPost, error) {
	var post model.DBPost
	filter := bson.M{"Title": title}
	err := i.client.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("post not found")
		}
		return nil, fmt.Errorf("failed to find post: %v", err)
	}

	return &post, nil
}
