package repository

import (
	"context"
	"database/sql"

	"bimbo/internal/model"
	"bimbo/internal/repository/psql"
)

type AuthRepositoryInterface interface {
	CreateUser(ctx context.Context, user *model.User) (int, error)
	GetUser(ctx context.Context, username, password string) (int, error) // todo: rename
	// CreateSession(context.Context, *model.TokenDetails) error            // todo : remove token , end time
	// UpdateSession(context.Context, *model.TokenDetails) error
	// DeleteSession(context.Context, *model.TokenDetails) error
}

type CompanyRepoInterface interface {
	Create(context.Context, model.Company) error
	GetList(context.Context) ([]model.Company, error)
}

type DepartamentRepoInterface interface {
	Create(context.Context, model.Departament) error
	GetList(context.Context) ([]model.Departament, error)
}
type PositionRepoInterface interface {
	Create(context.Context, model.Position) error
	GetList(context.Context) ([]model.Position, error)
}

type RoleRepoInterface interface {
	Create(context.Context, model.Role) error
	GetList(context.Context) ([]model.Role, error)
}

type TemplateRepoInterface interface {
	Create(context.Context, model.Template) error
	GetList(context.Context) ([]model.Template, error)
}

type ChoiceRepoInterface interface {
	Create(context.Context, []model.Choice, int) error
	GetList(context.Context, int) ([]model.Choice, error)
}

type Repositories struct {
	AuthRepositoryInterface
	CompanyRepoInterface
	DepartamentRepoInterface
	PositionRepoInterface
	RoleRepoInterface
	TemplateRepoInterface
	ChoiceRepoInterface
}

func RepositoryInit(db *sql.DB) Repositories {
	return Repositories{
		AuthRepositoryInterface:  psql.AuthRepositoryInit(db),
		CompanyRepoInterface:     psql.CompanyRepositoryInit(db),
		DepartamentRepoInterface: psql.DepartamentRepoInit(db),
		PositionRepoInterface:    psql.PositionRepoInit(db),
		RoleRepoInterface:        psql.RoleRepoInit(db),
		TemplateRepoInterface:    psql.TemplateRepoInit(db),
		ChoiceRepoInterface:      psql.ChoiceRepoInit(db),
	}
}
