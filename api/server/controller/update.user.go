package controller

import (
	"context"
	"errors"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i *implUser) UpdateUser(ctx context.Context, id string, user *model.User) error {
	findId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}
	err = i.repo.UpdateUser(ctx, findId, user)
	if err != nil {
		return errors.New("user not found")
	}
	return nil
}
