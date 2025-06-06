package router

import (
	"fmt"
	"go-task-manager-api/internal/app"
	"go-task-manager-api/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *app.Application) http.Handler {
	handler := handlers.New(app.Service)

	r := gin.Default()
	v1Route := r.Group("/v1")
	{
		v1Route.GET("/health", handler.Health)

		v1Route.POST("/auth/register", handler.User.Register)
		v1Route.POST("/auth/login", handler.User.Login)

		v1Route.POST("/tasks", handler.Task.CreateTask)
		v1Route.GET("/tasks/:taskID", handler.Task.GetTask)
	}
	r.Run(fmt.Sprintf(":%s", app.Config.Port))
	return r.Handler()
}
