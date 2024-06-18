package repo

import (
	"context"
	"log"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
)

func (i *impl) CreatePost(ctx context.Context, post *model.CreatePostRequest) error {
	_, err := i.client.InsertOne(ctx, post)
	if err != nil {
		log.Println("repo-CreatePost", err)
		return err
	}
	return nil
}
