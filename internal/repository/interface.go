package repository

import "taskmaster/internal/domain"

type TaskRepository interface {
	Create(task *domain.Task) error
	GetByID(id uint) (*domain.Task, error)
	GetAll() ([]*domain.Task, error)
	GetByStatus(status domain.TaskStatus) ([]*domain.Task, error)
	Update(task *domain.Task) error
	Delete(id uint) error
	GetOverDueTasks() ([]*domain.Task, error)
}
