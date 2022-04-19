package service

import (
	"context"

	"bimbo/internal/model"
	"bimbo/internal/repository"
)

type CompanyService struct {
	repo repository.CompanyRepoInterface
}

func CompanyServiceInit(r repository.CompanyRepoInterface) CompanyServiceInterface {
	return CompanyService{repo: r}
}

func (cuc CompanyService) Create(c model.Company) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.repo.Create(ctx, c)
}

// func (cuc CompanyService) GetCompanyByID(id int) (*model.Company, error) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	return cuc.repo.GetByID(ctx, id)
// }

func (s CompanyService) GetList() ([]model.Company, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return s.repo.GetList(ctx)
}
