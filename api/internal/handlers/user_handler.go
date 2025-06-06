package handlers

import (
	"go-task-manager-api/internal/auth"
	"go-task-manager-api/internal/model"
	"go-task-manager-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
	jwt     *auth.JWTManager
}

func NewUserHandler(service *service.UserService, jwt *auth.JWTManager) *UserHandler {
	return &UserHandler{
		service: service,
		jwt:     jwt,
	}
}

type RegisterPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var payload RegisterPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := &model.User{}
	user.Email = payload.Email
	user.HashPassword(payload.Password)

	if err := h.service.CreateUser(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

type LoginPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" bindin:"required"`
}

func (h *UserHandler) Login(c *gin.Context) {
	var payload LoginPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get User by Email
	user, err := h.service.GetUserByEmail(c.Request.Context(), payload.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Verify the user's password
	isPasswordValid := user.VerifyPassword(payload.Password)
	if !isPasswordValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user"})
		return
	}

	token, err := h.jwt.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": token})
}
