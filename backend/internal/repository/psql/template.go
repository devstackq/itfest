package psql

import (
	"context"
	"database/sql"

	"bimbo/internal/model"
)

type TemplateRepository struct {
	db *sql.DB
}

func TemplateRepoInit(db *sql.DB) *TemplateRepository {
	return &TemplateRepository{
		db: db,
	}
}

func (cr TemplateRepository) Create(ctx context.Context, cat model.Template) error {
	sqlQuery := `INSERT INTO bimbo_template(name) VALUES($1) RETURNING id`

	err := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cr TemplateRepository) GetList(ctx context.Context) ([]model.Template, error) {
	query := `SELECT id, name FROM bimbo_template`
	result := []model.Template{}
	rows, err := cr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := model.Template{}
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
