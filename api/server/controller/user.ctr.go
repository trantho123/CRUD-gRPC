package controller

import (
	"context"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"github.com/trantho123/CRUD-gRPC/api/server/repo"
)

type UserController interface {
	GetUser(ctx context.Context, id string) (*model.UserDB, error)
	UpdateUser(ctx context.Context, id string, user *model.User) error
}

type implUser struct {
	repo repo.RegistryUser
}

func NewUserController(repo repo.RegistryUser) UserController {
	return &implUser{repo: repo}
}
