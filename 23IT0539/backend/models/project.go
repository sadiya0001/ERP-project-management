package models

import (
	"time"

	"gorm.io/gorm"
)

// Project represents a project in the ERP system.
type Project struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null"`
	ProjectType string         `json:"project_type" gorm:"type:varchar(100)"`
	Description string         `json:"description" gorm:"type:text"`
	Status      string         `json:"status" gorm:"type:varchar(50);default:'active'"`
	StartDate   *time.Time     `json:"start_date" gorm:"type:date"`
	EndDate     *time.Time     `json:"end_date" gorm:"type:date"`
	CreatedBy   uint           `json:"created_by"`
	Creator     User           `json:"creator" gorm:"foreignKey:CreatedBy"`
	Members     []ProjectMember `json:"members,omitempty" gorm:"foreignKey:ProjectID"`
	Tasks       []Task         `json:"tasks,omitempty" gorm:"foreignKey:ProjectID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// ProjectMember represents the association between a project and a user.
type ProjectMember struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ProjectID uint      `json:"project_id" gorm:"not null;index"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	Role      string    `json:"role" gorm:"type:varchar(50)"`
	JoinedAt  time.Time `json:"joined_at" gorm:"autoCreateTime"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Project   Project   `json:"-" gorm:"foreignKey:ProjectID"`
}

// CreateProjectInput is the payload for creating a project.
type CreateProjectInput struct {
	Title       string `json:"title" binding:"required"`
	ProjectType string `json:"project_type"`
	Description string `json:"description"`
	Status      string `json:"status"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

// UpdateProjectInput is the payload for updating a project.
type UpdateProjectInput struct {
	Title       *string `json:"title"`
	ProjectType *string `json:"project_type"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	StartDate   *string `json:"start_date"`
	EndDate     *string `json:"end_date"`
}

// AddMemberInput is the payload for adding a member to a project.
type AddMemberInput struct {
	UserID uint   `json:"user_id" binding:"required"`
	Role   string `json:"role" binding:"required"`
}
