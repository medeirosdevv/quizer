package middleware

import (
    "strings"
    "quiz-backend/pkg/jwt"
    "github.com/gin-gonic/gin"
)

func Auth(jwtSecret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "Token não fornecido"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        claims, err := jwt.ValidateToken(tokenString, jwtSecret)
        if err != nil {
            c.JSON(401, gin.H{"error": "Token inválido"})
            c.Abort()
            return
        }

        c.Set("user_id", claims.UserID)
        c.Set("username", claims.Username)

        c.Next()
    }
}