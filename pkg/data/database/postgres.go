package database

import (
	"awcoding.com/back/pkg/infrastructure/config"
	"database/sql"
	"fmt"
)

func ConnectPostgresDB(cfg config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode,
		))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
