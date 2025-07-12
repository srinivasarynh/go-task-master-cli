package domain

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null; size:255" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Status      TaskStatus     `gorm:"default:pending" json:"status"`
	Priority    TaskPriority   `gorm:"default:medium" json:"priority"`
	DueDate     *time.Time     `json:"due_date,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusCanceled  TaskStatus = "canceled"
)

type TaskPriority string

const (
	TaskPriorityLow    TaskPriority = "low"
	TaskPriorityMedium TaskPriority = "medium"
	TaskPriorityHigh   TaskPriority = "high"
)

func (t *Task) IsCompleted() bool {
	return t.Status == TaskStatusCompleted
}

func (t *Task) IsOverdue() bool {
	if t.DueDate == nil {
		return false
	}
	return t.DueDate.Before(time.Now()) && !t.IsCompleted()
}
