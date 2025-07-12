package repository

import (
	"gorm.io/gorm"
	"taskmaster/internal/domain"
	"time"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *domain.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetByID(id uint) (*domain.Task, error) {
	var task domain.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepository) GetAll() ([]*domain.Task, error) {
	var tasks []*domain.Task
	err := r.db.Order("created_at desc").Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetByStatus(status domain.TaskStatus) ([]*domain.Task, error) {
	var tasks []*domain.Task
	err := r.db.Where("status = ?", status).Order("created_at desc").Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) Update(task *domain.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Task{}, id).Error
}

func (r *taskRepository) GetOverDueTasks() ([]*domain.Task, error) {
	var tasks []*domain.Task
	now := time.Now()
	err := r.db.Where("due_date < ? AND status = ?", now, domain.TaskStatusPending).Order("due_date asc").Find(&tasks).Error

	return tasks, err
}
