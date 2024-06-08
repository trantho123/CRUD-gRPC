package controller

import (
	"context"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"github.com/trantho123/CRUD-gRPC/api/server/repo"
)

type Controller interface {
	GetPost(ctx context.Context, id string) (*model.DBPost, error)
}

type impl struct {
	repo repo.Registry
}

func New(repo repo.Registry) Controller {
	return &impl{repo: repo}
}
