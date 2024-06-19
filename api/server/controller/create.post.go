package controller

import (
	"context"
	"errors"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
)

func (i *implPost) CreatePost(ctx context.Context, post *model.CreatePostRequest) error {
	err := i.repo.CreatePost(ctx, post)
	if err != nil {
		return errors.New("failed to create post")
	}
	return nil
}
