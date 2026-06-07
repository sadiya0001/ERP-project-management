package handlers

import (
	"net/http"
	"time"

	"erp-project-management/database"
	"erp-project-management/models"

	"github.com/gin-gonic/gin"
)

// GetWorkLogs returns all work logs for a specific task.
func GetWorkLogs(c *gin.Context) {
	taskID := c.Param("id")

	// Verify task exists
	var task models.Task
	if err := database.DB.First(&task, taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var workLogs []models.WorkLog
	if err := database.DB.Where("task_id = ?", taskID).
		Preload("User").
		Order("log_date DESC").
		Find(&workLogs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch work logs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"work_logs": workLogs})
}

// CreateWorkLog adds a new work log entry to a task.
func CreateWorkLog(c *gin.Context) {
	taskID := c.Param("id")

	// Verify task exists
	var task models.Task
	if err := database.DB.First(&task, taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input models.CreateWorkLogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")

	workLog := models.WorkLog{
		TaskID:      task.ID,
		UserID:      userID.(uint),
		Hours:       input.Hours,
		Description: input.Description,
		LogDate:     time.Now(),
	}

	if input.LogDate != "" {
		t, err := time.Parse("2006-01-02", input.LogDate)
		if err == nil {
			workLog.LogDate = t
		}
	}

	if err := database.DB.Create(&workLog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create work log"})
		return
	}

	// Update task time_spent (add hours converted to minutes)
	additionalMinutes := int(input.Hours * 60)
	database.DB.Model(&task).Update("time_spent", task.TimeSpent+additionalMinutes)

	database.DB.Preload("User").First(&workLog, workLog.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Work log added successfully",
		"work_log": workLog,
	})
}
