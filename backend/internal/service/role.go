package service

import (
	"context"

	"bimbo/internal/model"
	"bimbo/internal/repository"
)

type RoleService struct {
	repo repository.RoleRepoInterface
}

func RoleServiceInit(r repository.RoleRepoInterface) RoleServiceInterface {
	return RoleService{repo: r}
}

func (s RoleService) Create(c model.Role) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return s.repo.Create(ctx, c)
}

func (s RoleService) GetList() ([]model.Role, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return s.repo.GetList(ctx)
}
