package service

import (
	"context"

	"bimbo/internal/config"
	"bimbo/internal/model"

	"bimbo/internal/repository"
)

type Services struct {
	AuthService
	CompanyServiceInterface
	DepartamentServiceInterface
	PositionServiceInterface
	RoleServiceInterface
	ChoiceServiceInterface
	TemplateServiceInterface
	DocumentServiceInterface
}

func ServiceInit(repos repository.Repositories, cfg *config.Config) Services {
	return Services{
		AuthService:                 AuthServiceInit(repos.AuthRepositoryInterface, cfg),
		CompanyServiceInterface:     CompanyServiceInit(repos.CompanyRepoInterface),
		DepartamentServiceInterface: DepartamentServiceInit(repos.DepartamentRepoInterface),
		PositionServiceInterface:    PositionServiceInit(repos.PositionRepoInterface),
		RoleServiceInterface:        RoleServiceInit(repos.RoleRepoInterface),
		TemplateServiceInterface:    TemplateServiceInit(repos.TemplateRepoInterface),
		ChoiceServiceInterface:      ChoiceServiceInit(repos.ChoiceRepoInterface),
	}
}

// todo
type DocumentServiceInterface interface {
	Create(model.Document) error
}

type ChoiceServiceInterface interface {
	Create([]model.Choice, int) error
	GetList(int) ([]model.Choice, error)
}

type TemplateServiceInterface interface {
	Create(model.Template) error
	GetList() ([]model.Template, error)
}

type AuthServiceInterface interface {
	SignUp(ctx context.Context, user *model.User) (int, error)
	SignIn(ctx context.Context, username, password string) (int, error)
}

type DepartamentServiceInterface interface {
	Create(model.Departament) error
	GetList() ([]model.Departament, error)
}

type PositionServiceInterface interface {
	Create(model.Position) error
	GetList() ([]model.Position, error)
}

type CompanyServiceInterface interface {
	Create(model.Company) error
	GetList() ([]model.Company, error)
}
type RoleServiceInterface interface {
	Create(model.Role) error
	GetList() ([]model.Role, error)
}
