package controller

import (
	"context"
	"errors"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
)

func (i *impl) GetPosts(ctx context.Context, page int, limit int) (*[]model.DBPost, error) {
	posts, err := i.repo.GetPosts(ctx, page, limit)
	if err != nil {
		return nil, errors.New("failed to get posts")
	}
	return posts, nil
}
