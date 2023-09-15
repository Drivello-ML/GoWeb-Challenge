package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/GoWeb-Challenge/cmd/server/controllers"
	"github.com/GoWeb-Challenge/internal/domain"
	"github.com/GoWeb-Challenge/internal/tasks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("USER", "admin")
	_ = os.Setenv("PASSWORD", "imadmin")
	taskList := []domain.Task{
		{
			Id:          1,
			Description: "GetAll Test",
			Status:      false,
		},
	}
	taskRepository := tasks.NewTasksRepository(taskList)
	taskService := tasks.NewTaskService(taskRepository)
	taskHandler := controllers.NewTaskHandler(taskService)

	server := gin.Default()

	server.GET("/tasks", taskHandler.GetAllTasks())

	return server
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "imadmin")
	return req, httptest.NewRecorder()
}

func Test_GetAllTask_OK(t *testing.T) {
	// create server
	server := createServer()
	// create request and recorder
	req, recorder := createRequestTest(http.MethodGet, "/tasks", "")

	// indicate the server that can process the request
	server.ServeHTTP(recorder, req)
	var objRes []domain.Task
	assert.Equal(t, 200, recorder.Code)
	err := json.Unmarshal(recorder.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes) > 0)
}
