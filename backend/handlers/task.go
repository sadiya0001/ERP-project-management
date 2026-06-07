package handlers

import (
	"net/http"
	"strconv"
	"time"

	"erp-project-management/database"
	"erp-project-management/models"

	"github.com/gin-gonic/gin"
)

// GetTasks returns all tasks for a specific project.
func GetTasks(c *gin.Context) {
	projectID := c.Param("id")
	status := c.Query("status")
	priority := c.Query("priority")

	// Verify project exists
	var project models.Project
	if err := database.DB.First(&project, projectID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	query := database.DB.Where("project_id = ?", projectID)

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if priority != "" {
		query = query.Where("priority = ?", priority)
	}

	var tasks []models.Task
	if err := query.Preload("Assignee").Preload("Creator").
		Order("created_at DESC").
		Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// CreateTask creates a new task within a project.
func CreateTask(c *gin.Context) {
	projectID := c.Param("id")

	// Verify project exists
	var project models.Project
	if err := database.DB.First(&project, projectID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	var input models.CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")
	pid, _ := strconv.ParseUint(projectID, 10, 32)

	task := models.Task{
		ProjectID:   uint(pid),
		Title:       input.Title,
		Description: input.Description,
		Status:      "pending",
		Priority:    "medium",
		CreatedBy:   userID.(uint),
		AssignedTo:  input.AssignedTo,
	}

	if input.Status != "" {
		task.Status = input.Status
	}
	if input.Priority != "" {
		task.Priority = input.Priority
	}
	if input.Deadline != "" {
		t, err := time.Parse("2006-01-02", input.Deadline)
		if err == nil {
			task.Deadline = &t
		}
	}

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	database.DB.Preload("Assignee").Preload("Creator").First(&task, task.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task":    task,
	})
}

// UpdateTask updates an existing task.
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input models.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{}

	if input.Title != nil {
		updates["title"] = *input.Title
	}
	if input.Description != nil {
		updates["description"] = *input.Description
	}
	if input.Status != nil {
		updates["status"] = *input.Status
	}
	if input.Priority != nil {
		updates["priority"] = *input.Priority
	}
	if input.AssignedTo != nil {
		updates["assigned_to"] = *input.AssignedTo
	}
	if input.TimeSpent != nil {
		updates["time_spent"] = *input.TimeSpent
	}
	if input.Deadline != nil {
		t, err := time.Parse("2006-01-02", *input.Deadline)
		if err == nil {
			updates["deadline"] = t
		}
	}

	if err := database.DB.Model(&task).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	database.DB.Preload("Assignee").Preload("Creator").First(&task, task.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"task":    task,
	})
}

// DeleteTask soft-deletes a task.
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := database.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// AssignTask assigns a task to a specific user.
func AssignTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input models.AssignTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify user exists
	var user models.User
	if err := database.DB.First(&user, input.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := database.DB.Model(&task).Update("assigned_to", input.UserID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign task"})
		return
	}

	database.DB.Preload("Assignee").Preload("Creator").First(&task, task.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Task assigned successfully",
		"task":    task,
	})
}

// GetDashboardStats returns aggregated statistics for the dashboard.
func GetDashboardStats(c *gin.Context) {
	var totalProjects int64
	var activeProjects int64
	var completedProjects int64
	var totalTasks int64
	var pendingTasks int64
	var inProgressTasks int64
	var completedTasks int64
	var totalUsers int64
	var totalWorkLogHours float64

	database.DB.Model(&models.Project{}).Count(&totalProjects)
	database.DB.Model(&models.Project{}).Where("status = ?", "active").Count(&activeProjects)
	database.DB.Model(&models.Project{}).Where("status = ?", "completed").Count(&completedProjects)
	database.DB.Model(&models.Task{}).Count(&totalTasks)
	database.DB.Model(&models.Task{}).Where("status = ?", "pending").Count(&pendingTasks)
	database.DB.Model(&models.Task{}).Where("status = ?", "in_progress").Count(&inProgressTasks)
	database.DB.Model(&models.Task{}).Where("status = ?", "completed").Count(&completedTasks)
	database.DB.Model(&models.User{}).Count(&totalUsers)

	database.DB.Model(&models.WorkLog{}).Select("COALESCE(SUM(hours), 0)").Scan(&totalWorkLogHours)

	// Recent tasks
	var recentTasks []models.Task
	database.DB.Preload("Assignee").Preload("Creator").
		Order("created_at DESC").Limit(5).Find(&recentTasks)

	// Recent projects
	var recentProjects []models.Project
	database.DB.Preload("Creator").
		Order("created_at DESC").Limit(5).Find(&recentProjects)

	c.JSON(http.StatusOK, gin.H{
		"projects": gin.H{
			"total":     totalProjects,
			"active":    activeProjects,
			"completed": completedProjects,
		},
		"tasks": gin.H{
			"total":       totalTasks,
			"pending":     pendingTasks,
			"in_progress": inProgressTasks,
			"completed":   completedTasks,
		},
		"team_size":        totalUsers,
		"total_hours":      totalWorkLogHours,
		"recent_tasks":     recentTasks,
		"recent_projects":  recentProjects,
	})
}

// GetIntegrationProjects returns all project data for ERP integration.
func GetIntegrationProjects(c *gin.Context) {
	var projects []models.Project
	if err := database.DB.
		Preload("Creator").
		Preload("Members").
		Preload("Members.User").
		Preload("Tasks").
		Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

// GetIntegrationResources returns resource allocation data for ERP integration.
func GetIntegrationResources(c *gin.Context) {
	type ResourceAllocation struct {
		UserID      uint    `json:"user_id"`
		FirstName   string  `json:"first_name"`
		LastName    string  `json:"last_name"`
		Email       string  `json:"email"`
		Designation string  `json:"designation"`
		ProjectCount int64  `json:"project_count"`
		TaskCount    int64  `json:"task_count"`
		TotalHours   float64 `json:"total_hours"`
	}

	var users []models.User
	database.DB.Find(&users)

	var resources []ResourceAllocation
	for _, user := range users {
		var projectCount int64
		var taskCount int64
		var totalHours float64

		database.DB.Model(&models.ProjectMember{}).Where("user_id = ?", user.ID).Count(&projectCount)
		database.DB.Model(&models.Task{}).Where("assigned_to = ?", user.ID).Count(&taskCount)
		database.DB.Model(&models.WorkLog{}).Where("user_id = ?", user.ID).
			Select("COALESCE(SUM(hours), 0)").Scan(&totalHours)

		resources = append(resources, ResourceAllocation{
			UserID:       user.ID,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
			Designation:  user.Designation,
			ProjectCount: projectCount,
			TaskCount:    taskCount,
			TotalHours:   totalHours,
		})
	}

	c.JSON(http.StatusOK, gin.H{"resources": resources})
}
