package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a system user in the ERP.
type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Email       string         `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	PasswordHash string        `json:"-" gorm:"type:varchar(255);not null"`
	FirstName   string         `json:"first_name" gorm:"type:varchar(100)"`
	LastName    string         `json:"last_name" gorm:"type:varchar(100)"`
	Role        string         `json:"role" gorm:"type:varchar(50);default:'member'"`
	Avatar      string         `json:"avatar" gorm:"type:varchar(500)"`
	Phone       string         `json:"phone" gorm:"type:varchar(20)"`
	Nationality string         `json:"nationality" gorm:"type:varchar(100)"`
	Designation string         `json:"designation" gorm:"type:varchar(100)"`
	Skills      string         `json:"skills" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// RegisterInput is the payload for user registration.
type RegisterInput struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

// LoginInput is the payload for user login.
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UpdateUserInput is the payload for updating a user profile.
type UpdateUserInput struct {
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	Phone       *string `json:"phone"`
	Nationality *string `json:"nationality"`
	Designation *string `json:"designation"`
	Skills      *string `json:"skills"`
	Avatar      *string `json:"avatar"`
	Role        *string `json:"role"`
}

// HashPassword hashes the given plaintext password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword compares the hashed password with the given plaintext password.
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
