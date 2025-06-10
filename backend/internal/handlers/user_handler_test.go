package handlers

import (
	"bytes"
	"encoding/json"

	"go-task-manager-backend/internal/auth"
	"go-task-manager-backend/internal/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.MockUserService)
	mockJWT := new(mocks.MockJWTManager)

	var jwtManager auth.IJWTManger = mockJWT
	handler := NewUserHandler(mockService, jwtManager)
	endpoint := "/v1/auth/register"
	router := gin.Default()
	router.POST(endpoint, handler.Register)

	payload := RegisterPayload{
		Email:    "test@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", endpoint, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}
