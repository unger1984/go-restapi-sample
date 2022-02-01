package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := LoadConfig("./config.test.yaml")
	assert.Nilf(t, err, "No config.test.yaml: %v", err)

	assert.Equalf(t, cfg.Env, "production",
		"Incorrect Env. Expect \"%s\", got \"%v\"", "production", cfg.Env,
	)

	assert.Equalf(t, cfg.HttpServerConfig.CertKey, "key.pem",
		"Incorrect server.key. Expect \"%s\", got \"%v\"", "key.pem", cfg.HttpServerConfig.CertKey,
	)
	assert.Equalf(t, cfg.HttpServerConfig.CertPem, "chain.pem",
		"Incorrect server.cert. Expect \"%s\", got \"%v\"", "chain.pem", cfg.HttpServerConfig.CertPem,
	)
	assert.Equalf(t, cfg.HttpServerConfig.Host, "0.0.0.0",
		"Incorrect server.host. Expect \"%s\", got \"%v\"", "0.0.0.0", cfg.HttpServerConfig.Host,
	)
	assert.Equalf(t, cfg.HttpServerConfig.Port, "8080",
		"Incorrect server.post. Expect \"%s\", got \"%v\"", "8080", cfg.HttpServerConfig.Port,
	)
	assert.Equalf(t, cfg.HttpServerConfig.Static, "/upload",
		"Incorrect server.static. Expect \"%s\", got \"%v\"", "/upload", cfg.HttpServerConfig.Static,
	)
	assert.Equalf(t, cfg.DBConfig.Host, "localhost",
		"Incorrect db.host. Expect \"%s\", got \"%v\"", "localhost", cfg.DBConfig.Host,
	)
	assert.Equalf(t, cfg.DBConfig.Port, "5432",
		"Incorrect db.port. Expect \"%s\", got \"%v\"", "5432", cfg.DBConfig.Port,
	)
	assert.Equalf(t, cfg.DBConfig.Username, "postgres",
		"Incorrect db.username. Expect \"%s\", got \"%v\"", "postgres", cfg.DBConfig.Username,
	)
	assert.Equalf(t, cfg.DBConfig.Password, "postgres",
		"Incorrect db.pasword. Expect \"%s\", got \"%v\"", "postgres", cfg.DBConfig.Password,
	)
	assert.Equalf(t, cfg.DBConfig.DBName, "test",
		"Incorrect db.name. Expect \"%s\", got \"%v\"", "test", cfg.DBConfig.DBName,
	)
	assert.Equalf(t, cfg.DBConfig.SSLMode, "disable",
		"Incorrect db.sslmode. Expect \"%s\", got \"%v\"", "disable", cfg.DBConfig.SSLMode,
	)
	assert.Equalf(t, cfg.DBConfig.MigrationsPath, "/migrations",
		"Incorrect db.migrationspath. Expect \"%s\", got \"%v\"", "/migrations", cfg.DBConfig.MigrationsPath,
	)
	assert.Equalf(t, cfg.Secret.Salt, "_superSecretSalt_",
		"Incorrect secret.salt. Expect \"%s\", got \"%v\"", "_superSecretSalt_", cfg.Secret.Salt,
	)
	assert.Equalf(t, cfg.Secret.TokenKey, "_superSecretTokenKey_",
		"Incorrect secret.tokenkey. Expect \"%s\", got \"%v\"", "_superSecretTokenKey_", cfg.Secret.TokenKey,
	)

	_, err = LoadConfig("./config.broken.yaml")
	assert.NotNilf(t, err, "No config.test.yaml: %v", err)
}
