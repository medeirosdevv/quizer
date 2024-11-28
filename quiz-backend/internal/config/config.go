package config

import "os"

type Config struct {
    Port        string
    DatabaseURL string
    JWTSecret   string
    FrontendURL string
}

func Load() *Config {
    return &Config{
        Port:        getEnv("PORT", "8080"),
        DatabaseURL: getEnv("DATABASE_URL", "quiz.db"),
        JWTSecret:   getEnv("JWT_SECRET", "seu_segredo_aqui"),
        FrontendURL: getEnv("FRONTEND_URL", "http://localhost:3000"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}