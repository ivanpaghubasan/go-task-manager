package router

import (
	"go-task-manager-backend/internal/app"
	"go-task-manager-backend/internal/handlers"
	"go-task-manager-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *app.Application) *gin.Engine {
	handler := handlers.New(*app.Service, app.JWTManager)
	authMiddlerware := middleware.JWTAuthMiddleware(app.JWTManager)

	r := gin.Default()

	v1Route := r.Group("/v1")
	{
		v1Route.GET("/health", handler.Health)

		v1Route.POST("/auth/register", handler.User.Register)
		v1Route.POST("/auth/login", handler.User.Login)

		protected := v1Route.Group("/")
		protected.Use(authMiddlerware)
		{
			protected.POST("/tasks", handler.Task.CreateTask)
			protected.GET("/tasks/:taskID", handler.Task.GetTask)
		}
	}

	return r
}
