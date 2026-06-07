package models

import (
	"time"
)

// WorkLog represents a time entry against a task.
type WorkLog struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	TaskID      uint      `json:"task_id" gorm:"not null;index"`
	UserID      uint      `json:"user_id" gorm:"not null;index"`
	Hours       float64   `json:"hours" gorm:"not null"`
	Description string    `json:"description" gorm:"type:text"`
	LogDate     time.Time `json:"log_date" gorm:"type:date"`
	Task        Task      `json:"task,omitempty" gorm:"foreignKey:TaskID"`
	User        User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time `json:"created_at"`
}

// CreateWorkLogInput is the payload for adding a work log entry.
type CreateWorkLogInput struct {
	Hours       float64 `json:"hours" binding:"required,gt=0"`
	Description string  `json:"description"`
	LogDate     string  `json:"log_date"`
}
