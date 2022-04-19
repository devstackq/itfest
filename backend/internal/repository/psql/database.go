package psql

import (
	"database/sql"

	"bimbo/internal/config"
)

func InitDb(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://"+cfg.DB.Username+":"+cfg.DB.Password+"@"+cfg.DB.Host+"/"+cfg.DB.DBName+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
