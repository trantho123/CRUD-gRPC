package repo

import (
	"context"
	"log"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
)

func (i *implPost) CreatePost(ctx context.Context, post *model.CreatePostRequest) error {
	_, err := i.collection.InsertOne(ctx, post)
	if err != nil {
		log.Println("repo-CreatePost", err)
		return err
	}
	return nil
}
