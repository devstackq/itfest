package service

import (
	"context"

	"bimbo/internal/model"
	"bimbo/internal/repository"
)

type DepartamentService struct {
	repo repository.DepartamentRepoInterface
}

func DepartamentServiceInit(r repository.DepartamentRepoInterface) DepartamentServiceInterface {
	return DepartamentService{repo: r}
}

func (cuc DepartamentService) Create(c model.Departament) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.repo.Create(ctx, c)
}

func (s DepartamentService) GetList() ([]model.Departament, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return s.repo.GetList(ctx)
}
