package psql

import (
	"context"
	"database/sql"

	"bimbo/internal/model"
)

type CompanyRepository struct {
	db *sql.DB
}

func CompanyRepositoryInit(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (cr CompanyRepository) Create(ctx context.Context, cat model.Company) error {
	sqlQuery := `INSERT INTO bimbo_company(name) VALUES($1) RETURNING id`
	err := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cr CompanyRepository) GetList(ctx context.Context) ([]model.Company, error) {
	query := `SELECT id, name FROM bimbo_company`
	result := []model.Company{}
	rows, err := cr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := model.Company{}
		if err = rows.Scan(
			&temp.ID,
			&temp.Name,
		); err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return result, nil
}
