package handler

import (
    "quiz-backend/internal/service"
    "github.com/gin-gonic/gin"
)

type HomeHandler struct {
    userService *service.UserService
}

func NewHomeHandler(userService *service.UserService) *HomeHandler {
    return &HomeHandler{
        userService: userService,
    }
}

func (h *HomeHandler) GetHomeData(c *gin.Context) {
    userID, _ := c.Get("user_id")
    username, _ := c.Get("username")

    user, err := h.userService.GetUserByID(userID.(uint))
    if err != nil {
        c.JSON(404, gin.H{"error": "Usuário não encontrado"})
        return
    }

    c.JSON(200, gin.H{
        "message": "Bem-vindo à página inicial",
        "user": gin.H{
            "id": user.ID,
            "name": user.Name,
            "email": user.Email,
            "username": username,
        },
    })
}