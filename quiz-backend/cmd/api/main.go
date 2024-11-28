package main

import (
	"log"
	"quiz-backend/internal/config"
	"quiz-backend/internal/handler"
	"quiz-backend/internal/middleware"
	"quiz-backend/internal/repository/sqlite"
	"quiz-backend/internal/service"
	"quiz-backend/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.NewSQLiteDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	userRepo := sqlite.NewUserRepository(db)

	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	userService := service.NewUserService(userRepo)

	authHandler := handler.NewAuthHandler(authService)
	homeHandler := handler.NewHomeHandler(userService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.FrontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	protected := r.Group("/api")
	protected.Use(middleware.Auth(cfg.JWTSecret))
	{
		protected.GET("/home", homeHandler.GetHomeData)
	}

	log.Printf("Servidor iniciado na porta %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
