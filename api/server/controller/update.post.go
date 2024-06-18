package controller

import (
	"context"
	"errors"
	"log"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i *impl) UpdatePostById(ctx context.Context, id string, post *model.UpdatePost) error {
	log.Println("controller-UpdatePostById: ", post)
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
	if err = i.repo.UpdatePost(ctx, findId, post); err != nil {
		return errors.New("failed to update post")
	}
	return nil
}
