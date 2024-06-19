package controller

import (
	"context"
	"errors"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i *implUser) GetUser(ctx context.Context, id string) (*model.UserDB, error) {
	findId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	user, err := i.repo.GetUser(ctx, findId)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
