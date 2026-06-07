package models

import (
	"time"

	"gorm.io/gorm"
)

// Task represents a task within a project.
type Task struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ProjectID   uint           `json:"project_id" gorm:"not null;index"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null"`
	Description string         `json:"description" gorm:"type:text"`
	Status      string         `json:"status" gorm:"type:varchar(50);default:'pending'"`
	Priority    string         `json:"priority" gorm:"type:varchar(50);default:'medium'"`
	AssignedTo  *uint          `json:"assigned_to" gorm:"index"`
	CreatedBy   uint           `json:"created_by"`
	Deadline    *time.Time     `json:"deadline" gorm:"type:date"`
	TimeSpent   int            `json:"time_spent" gorm:"default:0"`
	Project     Project        `json:"project,omitempty" gorm:"foreignKey:ProjectID"`
	Assignee    *User          `json:"assignee,omitempty" gorm:"foreignKey:AssignedTo"`
	Creator     User           `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
	WorkLogs    []WorkLog      `json:"work_logs,omitempty" gorm:"foreignKey:TaskID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// CreateTaskInput is the payload for creating a task.
type CreateTaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	AssignedTo  *uint  `json:"assigned_to"`
	Deadline    string `json:"deadline"`
}

// UpdateTaskInput is the payload for updating a task.
type UpdateTaskInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	Priority    *string `json:"priority"`
	AssignedTo  *uint   `json:"assigned_to"`
	Deadline    *string `json:"deadline"`
	TimeSpent   *int    `json:"time_spent"`
}

// AssignTaskInput is the payload for assigning a task to a user.
type AssignTaskInput struct {
	UserID uint `json:"user_id" binding:"required"`
}
