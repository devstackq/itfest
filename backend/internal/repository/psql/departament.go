package psql

import (
	"context"
	"database/sql"

	"bimbo/internal/model"
)

type DepartamentRepository struct {
	db *sql.DB
}

func DepartamentRepoInit(db *sql.DB) *DepartamentRepository {
	return &DepartamentRepository{
		db: db,
	}
}

func (cr DepartamentRepository) Create(ctx context.Context, cat model.Departament) error {
	sqlQuery := `INSERT INTO bimbo_departament(name) VALUES($1) RETURNING id`

	err := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cr DepartamentRepository) GetList(ctx context.Context) ([]model.Departament, error) {
	query := `SELECT id, name FROM bimbo_departament`
	result := []model.Departament{}
	rows, err := cr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := model.Departament{}
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
