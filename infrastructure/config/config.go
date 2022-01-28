package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

type HttpServerConfig struct {
	Host    string
	Port    string
	CertPem string
	CertKey string
}

type DBConfig struct {
	Host           string
	Port           string
	Username       string
	Pasword        string
	DBName         string
	SSLMode        string
	MigrationsPath string
}

type Secret struct {
	Salt     string
	TokenKey string
}

type Config struct {
	ENV              string
	HttpServerConfig HttpServerConfig
	DBConfig         DBConfig
	Secret           Secret
}

var instance *Config
var once sync.Once

// GetInstance Возвращает текущий экземплар настроек
func GetInstance() *Config {
	once.Do(func() {
		instance = new(Config)
	})
	return instance
}

// Reload Перечитывает настроки из ..env файла и
// сохраняет по переданному указателю
// в случает ошибки вернет ошибку, при
// успехе - nil
func Reload(c *Config) error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	loadEnvs(c)
	return nil
}

// Load читает настроки из ..env файла и
// сохраняет по переданному указателю
// в случает ошибки вернет ошибку, при
// успехе - nil
// Если настройки были ранее прочитаны, вернет ошибку,
// для этой задачи используется метод Reload
func Load(c *Config) error {
	if c.ENV == "" {
		return Reload(c)
	}
	return errors.New("already Loaded")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func loadEnvs(c *Config) {
	c.ENV = os.Getenv("ENV")

	serverConfig := HttpServerConfig{}
	serverConfig.CertKey = getEnv("ssl.key", "")
	serverConfig.CertPem = getEnv("ssl.cert", "")
	serverConfig.Host = getEnv("server.host", "")
	serverConfig.Port = getEnv("server.port", "8443")

	dbConfig := DBConfig{}
	dbConfig.Host = getEnv("db.host", "localhost")
	dbConfig.Port = getEnv("db.port", "5432")
	dbConfig.Username = getEnv("db.username", "postgres")
	dbConfig.Pasword = getEnv("db.password", "postgres")
	dbConfig.DBName = getEnv("db.name", "test")
	dbConfig.SSLMode = getEnv("db.sslmode", "disable")
	dbConfig.MigrationsPath = getEnv("db.migrationspath", "./migrations")

	secret := Secret{}
	secret.Salt = getEnv("secret.salt", "secretSalt")
	secret.TokenKey = getEnv("secret.tokenkey", "secretToken")

	c.HttpServerConfig = serverConfig
	c.DBConfig = dbConfig
}
