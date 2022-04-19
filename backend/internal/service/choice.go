package service

import (
	"context"

	"bimbo/internal/model"
	"bimbo/internal/repository"
)

type ChoiceService struct {
	repo repository.ChoiceRepoInterface
}

func ChoiceServiceInit(r repository.ChoiceRepoInterface) ChoiceServiceInterface {
	return ChoiceService{repo: r}
}

func (cuc ChoiceService) Create(c []model.Choice, id int) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.repo.Create(ctx, c, id)
}

func (s ChoiceService) GetList(id int) ([]model.Choice, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return s.repo.GetList(ctx, id)
}
