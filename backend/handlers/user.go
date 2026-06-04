package handlers

import (
	"net/http"

	"erp-project-management/database"
	"erp-project-management/models"

	"github.com/gin-gonic/gin"
)

// GetUsers returns a list of all users.
func GetUsers(c *gin.Context) {
	search := c.Query("search")
	role := c.Query("role")

	query := database.DB.Model(&models.User{})

	if search != "" {
		query = query.Where(
			"first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%",
		)
	}
	if role != "" {
		query = query.Where("role = ?", role)
	}

	var users []models.User
	if err := query.Order("created_at DESC").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUser returns details for a single user.
func GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUser updates a user's profile.
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{}

	if input.FirstName != nil {
		updates["first_name"] = *input.FirstName
	}
	if input.LastName != nil {
		updates["last_name"] = *input.LastName
	}
	if input.Phone != nil {
		updates["phone"] = *input.Phone
	}
	if input.Nationality != nil {
		updates["nationality"] = *input.Nationality
	}
	if input.Designation != nil {
		updates["designation"] = *input.Designation
	}
	if input.Skills != nil {
		updates["skills"] = *input.Skills
	}
	if input.Avatar != nil {
		updates["avatar"] = *input.Avatar
	}
	if input.Role != nil {
		// Only admins can change roles
		currentRole, _ := c.Get("role")
		if currentRole == "admin" {
			updates["role"] = *input.Role
		}
	}

	if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	database.DB.First(&user, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    user,
	})
}
