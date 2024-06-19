package controller

import (
	"context"
	"errors"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i *implPost) GetPost(ctx context.Context, id string) (*model.DBPost, error) {
	findId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	post, err := i.repo.GetPost(ctx, findId)
	if err != nil {
		return nil, errors.New("post not found")
	}
	if post == nil {
		return nil, errors.New("post not found")
	}
	return post, nil
}
