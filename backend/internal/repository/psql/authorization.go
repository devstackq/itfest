package psql

import (
	"context"
	"database/sql"
	"time"

	"bimbo/internal/model"
)

type AuthorizationRepository struct {
	db *sql.DB
}

func AuthRepositoryInit(db *sql.DB) *AuthorizationRepository {
	return &AuthorizationRepository{
		db: db,
	}
}

func (ur AuthorizationRepository) CreateUser(ctx context.Context, user *model.User) (id int, err error) {
	query := `INSERT INTO bimbo_user(full_name, phone, email,  password,  created_at,  updated_at,  company_id,  departament_id, position_id,  role_id)values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	row := ur.db.QueryRowContext(ctx, query, user.FullName, user.Phone, user.Email, user.Password, time.Now(), time.Now(), user.Company.ID, user.Departament.ID, user.Position.ID, user.Role.ID)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ur AuthorizationRepository) GetUser(ctx context.Context, username, password string) (lastID int, err error) {
	sqlQuery := `SELECT id FROM bimbo_user WHERE email = $1 AND password = $2`
	row := ur.db.QueryRowContext(ctx, sqlQuery, username, password)
	err = row.Scan(&lastID)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (ur AuthorizationRepository) CreateSession(ctx context.Context, token *model.TokenDetails) error {
	query := `INSERT INTO bazar_session(access_uuid, refresh_uuid, user_id)values($1, $2, $3)`
	row := ur.db.QueryRowContext(ctx, query, token.AccessUuid, token.RefreshUuid, token.UserID)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (ur AuthorizationRepository) UpdateSession(ctx context.Context, token *model.TokenDetails) error {
	query := `UPDATE bazar_session SET access_uuid=$1, refresh_uuid=$2  WHERE user_ID=$3`
	row := ur.db.QueryRowContext(ctx, query, token.AccessUuid, token.RefreshUuid, token.UserID)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (ur AuthorizationRepository) DeleteSession(ctx context.Context, token *model.TokenDetails) error {
	query := `DELETE bazar_session WHERE user_ID=$1`
	row := ur.db.QueryRowContext(ctx, query, token.UserID)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}
