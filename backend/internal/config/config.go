package config

import (
	"time"
)

const (
	defaultAppMode             = "debug"
	defaultAppPort             = ":6969"
	defaultAppSecret           = "some_secret_key"
	defaultAppReadTimeout      = 10
	defaultAppWriteTimeout     = 10
	defaultAppAllowCredentials = true
)

type AppCfg struct {
	Mode          string
	Port          string
	SecretAccess  string
	SecretRefresh string
	HashSalt      string
	TokenTTL      time.Duration
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
}

type DBConf struct {
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type Config struct {
	// AuthConfig
	App *AppCfg
	DB  *DBConf
}

func GetConfig() *Config {
	return &Config{
		App: &AppCfg{
			HashSalt:      getEnvAsStr("APP_HASH_SALT", "keyForSalt"),
			TokenTTL:      time.Duration(getEnvAsInt("APP_TOKEN_TTL", 86400)),
			Mode:          getEnvAsStr("APP_MODE", defaultAppMode),
			Port:          getEnvAsStr("APP_PORT", defaultAppPort),
			SecretAccess:  getEnvAsStr("APP_SECRET_ACCESS", "accessx"),
			SecretRefresh: getEnvAsStr("APP_SECRET_REFRESH", "refreshx"),

			ReadTimeout:  time.Duration(getEnvAsInt("APP_READ_TIMEOUT", defaultAppReadTimeout)) * time.Second,
			WriteTimeout: time.Duration(getEnvAsInt("APP_WRITE_TIMEOUT", defaultAppWriteTimeout)) * time.Second,
		},

		DB: &DBConf{
			Dialect:  getEnvAsStr("POSTGRES_DIALECT", "pgx"),
			Host:     getEnvAsStr("POSTGRES_URI", "127.0.0.1"), // postgresdb - for compose
			Port:     getEnvAsStr("POSTGRES_PORT", "5432"),
			Username: getEnvAsStr("POSTGRES_USER", "postgres"),
			Password: getEnvAsStr("POSTGRES_PASSWORD", "postgres"),
			DBName:   getEnvAsStr("POSTGRES_DB", "testdb"),
		},
	}
}
