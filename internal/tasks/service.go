package tasks

import (
	"github.com/GoWeb-Challenge/internal/domain"
)

// SERVICE
// This layer is in charge of the logic of the business.

// This is our service. It has a repository interface inside.
type taskService struct {
	repository TaskRepositoryInterface
}

// This is our service interface. It allows us to abstract the
type TaskServiceInterface interface {
	CreateTask(task domain.Task) error
	GetAllTasks() []domain.Task
	GetTask(id int) (domain.Task, error)
	UpdateTask(id int, newTask domain.Task) error
	DeleteTask(id int) error
}

func (s *taskService) CreateTask(task domain.Task) error {
	s.repository.CreateTask(task)
	return nil
}

func (s *taskService) GetAllTasks() []domain.Task {
	data := s.repository.GetAll()
	return data
}

func (s *taskService) GetTask(id int) (domain.Task, error) {
	data, err := s.repository.GetById(id)
	if err != nil {
		return domain.Task{}, err
	}
	return data, nil
}

func (s *taskService) UpdateTask(id int, newTask domain.Task) error {

	err := s.repository.UpdateById(id, newTask)
	if err != nil {
		return err
	}
	return nil
}

func (s *taskService) DeleteTask(id int) error {
	err := s.repository.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func NewTaskService(r TaskRepositoryInterface) TaskServiceInterface {
	return &taskService{
		repository: r,
	}
}
