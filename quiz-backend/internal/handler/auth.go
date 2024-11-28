package handler

import (
    "quiz-backend/internal/service"
    "github.com/gin-gonic/gin"
)

type AuthHandler struct {
    authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
    return &AuthHandler{
        authService: authService,
    }
}

func (h *AuthHandler) Register(c *gin.Context) {
    var input service.RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": "Dados inválidos"})
        return
    }

    if err := h.authService.Register(input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(201, gin.H{"message": "Usuário criado com sucesso"})
}

func (h *AuthHandler) Login(c *gin.Context) {
    var input service.LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": "Dados inválidos"})
        return
    }

    token, user, err := h.authService.Login(input)
    if err != nil {
        c.JSON(401, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{
        "token": token,
        "user": user,
    })
}