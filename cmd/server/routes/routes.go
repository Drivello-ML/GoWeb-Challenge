package routes

import (
	"github.com/GoWeb-Challenge/cmd/server/controllers"
	"github.com/GoWeb-Challenge/internal/domain"
	"github.com/GoWeb-Challenge/internal/tasks"
	"github.com/gin-gonic/gin"
)

type RouterInterface interface {
	LoadRoutes()
}

func (r *taskRouter) LoadRoutes() {
	tasksGroup := r.server.Group("/tasks")
	tasksGroup.GET("/", r.controller.GetAllTasks())
	tasksGroup.POST("/", r.controller.CreateTask())
	tasksGroup.GET("/:id", r.controller.GetTask())
	tasksGroup.PUT("/:id", r.controller.UpdateTask())
	tasksGroup.DELETE("/:id", r.controller.DeleteTask())

}

type taskRouter struct {
	controller *controllers.TaskHandler
	service    tasks.TaskServiceInterface
	server     *gin.Engine
}

func NewTaskRouter(s *gin.Engine, taskList []domain.Task) taskRouter {
	// new instances
	taskRepository := tasks.NewTasksRepository(taskList)
	taskService := tasks.NewTaskService(taskRepository)
	taskHandler := controllers.NewTaskHandler(taskService)

	return taskRouter{
		controller: taskHandler,
		service:    taskService,
		server:     s,
	}
}
