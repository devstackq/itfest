package psql

import (
	"context"
	"database/sql"

	"bimbo/internal/model"
)

type PositionRepository struct {
	db *sql.DB
}

func PositionRepoInit(db *sql.DB) *PositionRepository {
	return &PositionRepository{
		db: db,
	}
}

func (cr PositionRepository) Create(ctx context.Context, cat model.Position) error {
	sqlQuery := `INSERT INTO bimbo_position(name) VALUES($1) RETURNING id`

	err := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cr PositionRepository) GetList(ctx context.Context) ([]model.Position, error) {
	query := `SELECT id, name FROM bimbo_position`
	result := []model.Position{}
	rows, err := cr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := model.Position{}
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
