package config

import (
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

type HttpServerConfig struct {
	Host    string
	Port    string
	CertPem string
	CertKey string
	Static  string
}

type DBConfig struct {
	Host           string
	Port           string
	Username       string
	Password       string
	DBName         string
	SSLMode        string
	MigrationsPath string
}

type Secret struct {
	Salt     string
	TokenKey string
}

type Config struct {
	Env              string
	HttpServerConfig HttpServerConfig
	DBConfig         DBConfig
	Secret           Secret
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}
	dir, file := filepath.Split(path)
	viper.AddConfigPath(dir)
	viper.SetConfigType(file[strings.LastIndex(file, ".")+1:])
	viper.SetConfigName(file[:strings.LastIndex(file, ".")])
	setDefaults()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	loadEnvs(config)
	return config, nil
}

func setDefaults() {
	viper.SetDefault("server.key", "")
	viper.SetDefault("server.cert", "")
	viper.SetDefault("server.host", "")
	viper.SetDefault("server.port", "8443")
	viper.SetDefault("server.static", "")

	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.username", "postgres")
	viper.SetDefault("db.password", "postgres")
	viper.SetDefault("db.name", "test")
	viper.SetDefault("db.sslmode", "disable")
	viper.SetDefault("db.migrationspath", "./migrations")

	viper.SetDefault("secret.salt", "secretSalt")
	viper.SetDefault("secret.tokenkey", "secretToken")
}

func loadEnvs(c *Config) {
	c.Env = viper.GetString("env")

	serverConfig := HttpServerConfig{}
	serverConfig.CertKey = viper.GetString("server.key")
	serverConfig.CertPem = viper.GetString("server.cert")
	serverConfig.Host = viper.GetString("server.host")
	serverConfig.Port = viper.GetString("server.port")
	serverConfig.Static = viper.GetString("server.static")

	dbConfig := DBConfig{}
	dbConfig.Host = viper.GetString("db.host")
	dbConfig.Port = viper.GetString("db.port")
	dbConfig.Username = viper.GetString("db.username")
	dbConfig.Password = viper.GetString("db.password")
	dbConfig.DBName = viper.GetString("db.name")
	dbConfig.SSLMode = viper.GetString("db.sslmode")
	dbConfig.MigrationsPath = viper.GetString("db.migrationspath")

	secret := Secret{}
	secret.Salt = viper.GetString("secret.salt")
	secret.TokenKey = viper.GetString("secret.tokenkey")

	c.HttpServerConfig = serverConfig
	c.DBConfig = dbConfig
}
