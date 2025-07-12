package service

import (
	"taskmaster/internal/domain"
	"time"
)

type TaskService interface {
	CreateTask(title, description string, priority domain.TaskPriority, dueDate *time.Time) (*domain.Task, error)
	GetTask(id uint) (*domain.Task, error)
	GetAllTasks() ([]*domain.Task, error)
	GetTaskByStatus(status domain.TaskStatus) ([]*domain.Task, error)
	UpdateTask(id uint, title, description string, priority domain.TaskPriority, dueDate *time.Time) (*domain.Task, error)
	CompleteTask(id uint) error
	DeleteTask(id uint) error
	GetOverdueTasks() ([]*domain.Task, error)
}
