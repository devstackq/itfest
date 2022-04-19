package psql

import (
	"context"
	"database/sql"

	"bimbo/internal/model"
)

type RoleRepository struct {
	db *sql.DB
}

func RoleRepoInit(db *sql.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (cr RoleRepository) Create(ctx context.Context, cat model.Role) error {
	sqlQuery := `INSERT INTO bimbo_role(name) VALUES($1) RETURNING id`

	err := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cr RoleRepository) GetList(ctx context.Context) ([]model.Role, error) {
	query := `SELECT id, name FROM bimbo_role`
	result := []model.Role{}
	rows, err := cr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := model.Role{}
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
