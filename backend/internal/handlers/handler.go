package handlers

import (
	"go-task-manager-backend/internal/auth"
	"go-task-manager-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	User UserHandler
	Task TaskHandler
}

func New(service service.Service, jwt auth.IJWTManger) *Handlers {
	return &Handlers{
		User: NewUserHandler(service.User, jwt),
		Task: NewTaskHandler(service.Task),
	}
}

func (h *Handlers) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Good!"})
}
