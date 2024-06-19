package controller

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i *implPost) DeletePostById(ctx context.Context, id string) error {
	findId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}
	isPostExist, err := i.repo.GetPost(ctx, findId)
	if err != nil {
		return errors.New("post not found")
	}
	if isPostExist == nil {
		return errors.New("post not found")
	}
	if err := i.repo.DeletePost(ctx, findId); err != nil {
		return errors.New("failed to delete post")
	}
	return nil
}
