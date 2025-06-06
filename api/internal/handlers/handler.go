package handlers

import (
	"go-task-manager-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	User UserHandler
}

func New(service service.Service) Handlers {
	return Handlers{
		User: UserHandler{service: service.User},
	}
}

func (h *Handlers) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Good!"})
}
