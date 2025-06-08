package handlers

import (
	"go-task-manager/internal/auth"
	"go-task-manager/internal/service"
	"go-task-manager/views"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	User UserHandler
	Task *TaskHandler
}

func New(service service.Service, jwt auth.IJWTManger) *Handlers {
	return &Handlers{
		User: NewUserHandler(service.User, jwt),
		Task: NewTaskHandler(service.Task),
	}
}

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func (h *Handlers) IndexPage(c *gin.Context) {
	render(c, http.StatusOK, views.Index())
}

func (h *Handlers) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Good!"})
}
