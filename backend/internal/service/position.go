package service

import (
	"context"

	"bimbo/internal/model"
	"bimbo/internal/repository"
)

type PositionService struct {
	repo repository.PositionRepoInterface
}

func PositionServiceInit(r repository.PositionRepoInterface) PositionServiceInterface {
	return PositionService{repo: r}
}

func (s PositionService) Create(c model.Position) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return s.repo.Create(ctx, c)
}

func (s PositionService) GetList() ([]model.Position, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return s.repo.GetList(ctx)
}
