package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret string
	JWTExpiry string

	ServerPort          string
	Environment         string
	UrlFe               string
	GoogleRedirectURL   string
	GoogleClientID      string
	GoogleClientSecret  string
	GoogleApiOauth      string
	FrontendRedirectURL string
}

func LoadConfig() *Config {
	err := godotenv.Load() // Load from .env
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),

		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpiry: os.Getenv("JWT_EXPIRY"),

		ServerPort:          os.Getenv("PORT"),
		Environment:         os.Getenv("ENVIRONMENT"),
		UrlFe:               os.Getenv("URL_FE"),
		GoogleRedirectURL:   os.Getenv("GOOGLE_REDIRECT_URL"),
		GoogleClientID:      os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret:  os.Getenv("GOOGLE_CLIENT_SECRET"),
		GoogleApiOauth:      os.Getenv("GOOGLE_API_OAUTH"),
		FrontendRedirectURL: os.Getenv("FRONTEND_REDIRECT_URL"),
	}
}
