package handlers

import (
	"go-task-manager-api/internal/auth"
	"go-task-manager-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	User *UserHandler
	Task *TaskHandler
}

func New(service *service.Service, jwt *auth.JWTManager) *Handlers {
	return &Handlers{
		User: NewUserHandler(service.User, jwt),
		Task: NewTaskHandler(service.Task),
	}
}

func (h *Handlers) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Good!"})
}
