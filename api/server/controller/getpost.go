package controller

import (
	"context"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
)

func (i *impl) GetPost(ctx context.Context, id string) (*model.DBPost, error) {
	post, err := i.repo.GetPost(ctx, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}
