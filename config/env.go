package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAdress               string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		// PublicHost:             getEnv("PUBLIC_HOST", ""),
		// Port:                   getEnv("PORT", "8080"),
		// DBUser:                 getEnv("DB_USER", ""),
		// DBPassword:             getEnv("DB_PASSWORD", ""),
		// DBAdress:               fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		// DBName:                 getEnv("DB_NAME", ""),
		// JWTSecret:              getEnv("JWT_SECRET", ""),
		// JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		PublicHost:             getEnv("PUBLIC_HOST", ""),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("MYSQLUSER", ""),
		DBPassword:             getEnv("MYSQL_ROOT_PASSWORD", ""),
		DBAdress:               fmt.Sprintf("%s:%s", getEnv("MYSQL_HOST", ""), getEnv("MYSQL_PORT", "3306")),
		DBName:                 getEnv("MYSQL_DATABASE", ""),
		JWTSecret:              getEnv("JWT_SECRET", ""),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
