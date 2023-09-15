package tasks

import (
	"errors"

	"github.com/GoWeb-Challenge/internal/domain"
)

// REPOSITORY
// This is the layer that manages the database.

// This is our "database" in memory. It is an array of tasks.
// We have the definition of a "task" at our domain.
type inMemoryRepository struct {
	db []domain.Task
}

// This are the functions that our repository is going to do for us.
type TaskRepositoryInterface interface {
	CreateTask(task domain.Task) error
	GetAll() []domain.Task
	GetById(id int) (domain.Task, error)
	UpdateById(id int, task domain.Task) error
	DeleteById(id int) error
}

func (r *inMemoryRepository) CreateTask(task domain.Task) error {
	r.db = append(r.db, task)
	return nil
}

func (r *inMemoryRepository) GetAll() []domain.Task {
	return r.db
}

func (r *inMemoryRepository) GetById(id int) (domain.Task, error) {
	for _, t := range r.db {
		if t.Id == id {
			return t, nil
		}
	}
	return domain.Task{}, errors.New("Task not found")
}

func (r *inMemoryRepository) UpdateById(id int, task domain.Task) error {
	for i, t := range r.db {
		if t.Id == id {
			r.db[i] = task
			return nil
		}
	}
	return errors.New("Task not found")
}

func (r *inMemoryRepository) DeleteById(id int) error {
	for i, t := range r.db {
		if t.Id == id {
			r.db = append(r.db[:i], r.db[i+1:]...)
			return nil
		}
	}
	return errors.New("Task not found")
}

// This is the way we initialize our "database".
// Notice that we give it an array. That array comes from our tasks.csv file.
// It returns our "db" with the functions we loaded into the interface.
func NewTasksRepository(db []domain.Task) TaskRepositoryInterface {
	return &inMemoryRepository{
		db: db,
	}
}
