package service

import (
	"context"

	"bimbo/internal/model"
	"bimbo/internal/repository"
)

type TemplateService struct {
	repo repository.TemplateRepoInterface
}

func TemplateServiceInit(r repository.TemplateRepoInterface) TemplateServiceInterface {
	return TemplateService{repo: r}
}

func (cuc TemplateService) Create(c model.Template) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.repo.Create(ctx, c)
}

func (s TemplateService) GetList() ([]model.Template, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return s.repo.GetList(ctx)
}
