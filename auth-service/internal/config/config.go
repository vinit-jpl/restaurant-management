package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	SSLMode    string
	JWTSecret  string
}

func LoadConfig() *Config {

	_ = godotenv.Load()

	cfg := Config{

		Port:       getEnv("AUTH_PORT", "8081"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASS", "postgres@123"),
		DBName:     getEnv("DB_NAME", "users_db"),
		DBPort:     getEnv("DB_PORT", "5432"),
		SSLMode:    getEnv("SSL_MODE", "disable"),
		JWTSecret:  getEnv("JWT_SECRET", ""),
	}

	return &cfg

}

func getEnv(key, defaultValue string) string {

	value, ok := os.LookupEnv(key)

	if !ok {
		return defaultValue
	}

	return value
}
