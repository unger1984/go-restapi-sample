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

type preferences struct {
	ENV              string
	HttpServerConfig HttpServerConfig
}

var instance *preferences
var once sync.Once

// GetInstance Возвращает текущий экземплар настроек
func GetInstance() *preferences {
	once.Do(func() {
		instance = new(preferences)
	})
	return instance
}

// Reload Перечитывает настроки из .env файла и
// сохраняет по переданному указателю
// в случает ошибки вернет ошибку, при
// успехе - nil
func Reload(c *preferences) error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	loadEnvs(c)
	return nil
}

// Load читает настроки из .env файла и
// сохраняет по переданному указателю
// в случает ошибки вернет ошибку, при
// успехе - nil
// Если настройки были ранее прочитаны, вернет ошибку,
// для этой задачи используется метод Reload
func Load(c *preferences) error {
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

func loadEnvs(c *preferences) {
	c.ENV = os.Getenv("ENV")

	serverConfig := HttpServerConfig{}
	serverConfig.CertKey = getEnv("ssl.key", "")
	serverConfig.CertPem = getEnv("ssl.cert", "")
	serverConfig.Host = getEnv("server.host", "")
	serverConfig.Port = getEnv("server.port", "8443")

	c.HttpServerConfig = serverConfig
}
