package repo

import (
	"context"
	"log"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (i *impl) GetPosts(ctx context.Context, page int, limit int) (*[]model.DBPost, error) {
	var posts []model.DBPost
	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))

	cursor, err := i.client.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Println("repo-GetPosts1", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post model.DBPost
		if err := cursor.Decode(&post); err != nil {
			log.Println("repo-GetPosts2", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		log.Println("repo-GetPosts3", err)
		return nil, err
	}

	return &posts, nil
}
