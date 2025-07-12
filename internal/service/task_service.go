package service

import (
	"errors"
	"taskmaster/internal/domain"
	"taskmaster/internal/repository"
	"time"

	"github.com/sirupsen/logrus"
)

type taskService struct {
	taskRepo repository.TaskRepository
	logger   *logrus.Logger
}

func NewTaskService(taskRepo repository.TaskRepository, logger *logrus.Logger) TaskService {
	return &taskService{
		taskRepo: taskRepo,
		logger:   logger,
	}
}
func (s *taskService) CreateTask(title, description string, priority domain.TaskPriority, dueDate *time.Time) (*domain.Task, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	task := &domain.Task{
		Title:       title,
		Description: description,
		Status:      domain.TaskStatusPending,
		Priority:    priority,
		DueDate:     dueDate,
	}

	if err := s.taskRepo.Create(task); err != nil {
		s.logger.WithError(err).Error("failed to create task")
		return nil, err
	}

	s.logger.WithField("task_id", task.ID).Info("task created successfully")
	return task, nil
}

func (s *taskService) GetTask(id uint) (*domain.Task, error) {
	task, err := s.taskRepo.GetByID(id)
	if err != nil {
		s.logger.WithError(err).WithField("task_id", id).Error("failed to get task")
		return nil, err
	}
	return task, nil
}

func (s *taskService) GetAllTasks() ([]*domain.Task, error) {
	tasks, err := s.taskRepo.GetAll()
	if err != nil {
		s.logger.WithError(err).Error("failed to get all tasks")
		return nil, err
	}
	return tasks, err
}

func (s *taskService) GetTaskByStatus(status domain.TaskStatus) ([]*domain.Task, error) {
	tasks, err := s.taskRepo.GetByStatus(status)
	if err != nil {
		s.logger.WithError(err).WithField("status", status).Error("failed to get tasks")
		return nil, err
	}
	return tasks, nil
}

func (s *taskService) UpdateTask(id uint, title, description string, priority domain.TaskPriority, dueDate *time.Time) (*domain.Task, error) {
	task, err := s.taskRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if title != "" {
		task.Title = title
	}
	if description != "" {
		task.Description = description
	}
	if priority != "" {
		task.Priority = priority
	}
	if dueDate != nil {
		task.DueDate = dueDate
	}

	if err := s.taskRepo.Update(task); err != nil {
		s.logger.WithError(err).WithField("task_id", id).Error("failed to update task")
		return nil, err
	}

	s.logger.WithField("task_id", task.ID).Info("task updated successfully")
	return task, nil
}

func (s *taskService) CompleteTask(id uint) error {
	task, err := s.taskRepo.GetByID(id)
	if err != nil {
		return err
	}

	task.Status = domain.TaskStatusCompleted
	if err := s.taskRepo.Update(task); err != nil {
		s.logger.WithError(err).WithField("task_id", id).Error("failed to complete task")
		return err
	}

	s.logger.WithField("task_id", task.ID).Info("task completed successfully")
	return nil
}

func (s *taskService) DeleteTask(id uint) error {
	if err := s.taskRepo.Delete(id); err != nil {
		s.logger.WithError(err).WithField("task_id", id).Error("failed to delete task")
		return err
	}

	s.logger.WithField("task_id", id).Info("task deleted successfully")
	return nil
}

func (s *taskService) GetOverdueTasks() ([]*domain.Task, error) {
	tasks, err := s.taskRepo.GetOverDueTasks()
	if err != nil {
		s.logger.WithError(err).Error("failed to get overdue tasks")
		return nil, err
	}

	return tasks, nil
}
