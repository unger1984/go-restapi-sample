package database

import (
	"awcoding.com/back/infrastructure/config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func ConnectPostgresDB(cfg config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Pasword, cfg.DBName, cfg.SSLMode,
		))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
