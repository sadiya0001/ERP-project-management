package handlers

import (
	"net/http"
	"strconv"
	"time"

	"erp-project-management/database"
	"erp-project-management/models"

	"github.com/gin-gonic/gin"
)

// GetProjects returns a paginated list of projects with optional search.
func GetProjects(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")
	status := c.Query("status")

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit

	query := database.DB.Model(&models.Project{})

	if search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var projects []models.Project
	if err := query.Preload("Creator").Preload("Members").Preload("Members.User").
		Order("created_at DESC").
		Offset(offset).Limit(limit).
		Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}

// GetProject returns a single project with its members and tasks.
func GetProject(c *gin.Context) {
	id := c.Param("id")

	var project models.Project
	if err := database.DB.
		Preload("Creator").
		Preload("Members").
		Preload("Members.User").
		Preload("Tasks").
		Preload("Tasks.Assignee").
		First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"project": project})
}

// CreateProject creates a new project.
func CreateProject(c *gin.Context) {
	var input models.CreateProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")

	project := models.Project{
		Title:       input.Title,
		ProjectType: input.ProjectType,
		Description: input.Description,
		Status:      "active",
		CreatedBy:   userID.(uint),
	}

	if input.Status != "" {
		project.Status = input.Status
	}

	if input.StartDate != "" {
		t, err := time.Parse("2006-01-02", input.StartDate)
		if err == nil {
			project.StartDate = &t
		}
	}
	if input.EndDate != "" {
		t, err := time.Parse("2006-01-02", input.EndDate)
		if err == nil {
			project.EndDate = &t
		}
	}

	if err := database.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	// Add creator as team_lead by default
	member := models.ProjectMember{
		ProjectID: project.ID,
		UserID:    userID.(uint),
		Role:      "team_lead",
	}
	database.DB.Create(&member)

	// Reload with associations
	database.DB.Preload("Creator").Preload("Members").Preload("Members.User").First(&project, project.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Project created successfully",
		"project": project,
	})
}

// UpdateProject updates an existing project.
func UpdateProject(c *gin.Context) {
	id := c.Param("id")

	var project models.Project
	if err := database.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	var input models.UpdateProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{}

	if input.Title != nil {
		updates["title"] = *input.Title
	}
	if input.ProjectType != nil {
		updates["project_type"] = *input.ProjectType
	}
	if input.Description != nil {
		updates["description"] = *input.Description
	}
	if input.Status != nil {
		updates["status"] = *input.Status
	}
	if input.StartDate != nil {
		t, err := time.Parse("2006-01-02", *input.StartDate)
		if err == nil {
			updates["start_date"] = t
		}
	}
	if input.EndDate != nil {
		t, err := time.Parse("2006-01-02", *input.EndDate)
		if err == nil {
			updates["end_date"] = t
		}
	}

	if err := database.DB.Model(&project).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}

	database.DB.Preload("Creator").Preload("Members").Preload("Members.User").First(&project, project.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Project updated successfully",
		"project": project,
	})
}

// DeleteProject soft-deletes a project.
func DeleteProject(c *gin.Context) {
	id := c.Param("id")

	var project models.Project
	if err := database.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	if err := database.DB.Delete(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// AddProjectMember adds a user to a project.
func AddProjectMember(c *gin.Context) {
	projectID := c.Param("id")

	// Verify project exists
	var project models.Project
	if err := database.DB.First(&project, projectID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	var input models.AddMemberInput
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

	// Check if user is already a member
	var existing models.ProjectMember
	pid, _ := strconv.ParseUint(projectID, 10, 32)
	if err := database.DB.Where("project_id = ? AND user_id = ?", uint(pid), input.UserID).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User is already a member of this project"})
		return
	}

	member := models.ProjectMember{
		ProjectID: uint(pid),
		UserID:    input.UserID,
		Role:      input.Role,
	}

	if err := database.DB.Create(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add member"})
		return
	}

	database.DB.Preload("User").First(&member, member.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Member added successfully",
		"member":  member,
	})
}

// RemoveProjectMember removes a user from a project.
func RemoveProjectMember(c *gin.Context) {
	projectID := c.Param("id")
	userID := c.Param("userId")

	var member models.ProjectMember
	if err := database.DB.Where("project_id = ? AND user_id = ?", projectID, userID).First(&member).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found in this project"})
		return
	}

	if err := database.DB.Delete(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove member"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member removed successfully"})
}
