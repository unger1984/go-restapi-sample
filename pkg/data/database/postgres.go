package database

import (
	"awcoding.com/back/pkg/infrastructure/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

func ConnectPostgresDB(cfg config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode,
		))
	if err != nil {
		return nil, err
	}

	db.MapperFunc(func(s string) string {
		//fmt.Println(strings.ToLower(s[:1]) + s[1:])
		return strings.ToLower(s[:1]) + s[1:]
	})
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
