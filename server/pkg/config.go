package pkg

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database PostgresConfig
}

type PostgresConfig struct {
	Port     int
	User     string
	Password string
	Name     string
	Host     string
	SSLMode  string
}

func (p PostgresConfig) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		p.User, p.Password, p.Host, p.Port, p.Name, p.SSLMode)
}

func GetConfig() (Config, error) {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if err, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			log.Fatalf("Failed to parse config failed: %v", err)
		}
	}

	err := viper.Unmarshal(&cfg)

	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func bindEnv(key string, defaultValue interface{}, envName string) {
	viper.SetDefault(key, defaultValue)
	_ = viper.BindEnv(key, envName)
}

func init() {
	bindEnv("database.dbname", "postgres", "POSTGRES_DBNAME")
	bindEnv("database.user", "postgres", "POSTGRES_USER")
	bindEnv("database.password", "password", "POSTGRES_PASSWORD")
	bindEnv("database.host", "localhost", "POSTGRES_HOST")
	bindEnv("database.port", 5432, "POSTGRES_PORT")
	bindEnv("database.sslmode", "disable", "POSTGRES_SSLMODE")

	viper.AutomaticEnv()
}
