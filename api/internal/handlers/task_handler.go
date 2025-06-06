package handlers

import (
	"go-task-manager-api/internal/model"
	"go-task-manager-api/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

type TaskPayload struct {
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      string     `json:"status"`
	StartDate   time.Time  `json:"start_date"`
	DueDate     *time.Time `json:"due_date"`
	EndDate     *time.Time `json:"end_date"`
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var payload TaskPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := &model.Task{
		UserID:      payload.UserID,
		Title:       payload.Title,
		Description: payload.Description,
		Status:      payload.Status,
		StartDate:   payload.StartDate,
		DueDate:     payload.DueDate,
		EndDate:     payload.EndDate,
	}

	t, err := h.service.CreateTask(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": t})
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	paramID := c.Param("taskID")
	taskID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.service.GetTask(c, taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}
