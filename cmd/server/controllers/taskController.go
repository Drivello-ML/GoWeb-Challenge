package controllers

import (
	"net/http"
	"strconv"

	"github.com/GoWeb-Challenge/internal/domain"
	"github.com/GoWeb-Challenge/internal/tasks"
	"github.com/gin-gonic/gin"
)

// CONTROLLER
// This is the middleman between the client request and the logic of the bussiness.
// It validates the request, calls the service and generate a proper response.
// It also handles the errors that might occur during the process.

type TaskHandler struct {
	service tasks.TaskServiceInterface
}

func checkInternalError(c *gin.Context, err error) {
	if err != nil {
		c.String(http.StatusInternalServerError, "Something failed in our system", nil)
		return
	}
}

func getIdFromParam(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid id", nil)
	}
	return id
}

func (s *TaskHandler) CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody domain.Task
		if err := c.BindJSON(&requestBody); err != nil {
			c.String(http.StatusBadRequest, "Invalid request body", nil)
			return
		}
		s.service.CreateTask(requestBody)
		c.JSON(http.StatusCreated, "")
	}
}

func (s *TaskHandler) GetAllTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks := s.service.GetAllTasks()
		c.JSON(http.StatusOK, tasks)
	}
}
func (s *TaskHandler) GetTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		task, err := s.service.GetTask(getIdFromParam(c))
		checkInternalError(c, err)
		c.JSON(http.StatusCreated, task)
	}
}
func (s *TaskHandler) UpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody domain.Task
		if err := c.BindJSON(&requestBody); err != nil {
			c.String(http.StatusBadRequest, "Invalid request body", nil)
			return
		}
		err := s.service.UpdateTask(getIdFromParam(c), requestBody)
		checkInternalError(c, err)
		c.JSON(http.StatusNoContent, "")
	}
}
func (s *TaskHandler) DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := s.service.DeleteTask(getIdFromParam(c))
		checkInternalError(c, err)
		c.JSON(http.StatusNoContent, "")

	}
}

func NewTaskHandler(s tasks.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{
		service: s,
	}
}
