package controllers

import (
	"project-management/config"

	"github.com/gin-gonic/gin"
)

// ==========================
// CREATE PROJECT
// ==========================

func CreateProject(c *gin.Context) {

	var project struct {
		Name   string  `json:"name"`
		Budget float64 `json:"budget"`
	}

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})
		return
	}

	var id int

	query := `
	INSERT INTO projects (name, budget)
	VALUES ($1, $2)
	RETURNING project_id
	`

	err := config.DB.QueryRow(
		query,
		project.Name,
		project.Budget,
	).Scan(&id)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":    "Project created successfully",
		"project_id": id,
	})
}

// ==========================
// GET PROJECTS
// ==========================

func GetProjects(c *gin.Context) {

	rows, err := config.DB.Query(
		"SELECT project_id, name, budget FROM projects",
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	var projects []gin.H

	for rows.Next() {

		var id int
		var name string
		var budget float64

		err := rows.Scan(&id, &name, &budget)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		projects = append(projects, gin.H{
			"project_id": id,
			"name":       name,
			"budget":     budget,
		})
	}

	c.JSON(200, gin.H{
		"data": projects,
	})
}

// ==========================
// UPDATE PROJECT
// ==========================

func UpdateProject(c *gin.Context) {

	id := c.Param("id")

	var project struct {
		Name   string  `json:"name"`
		Budget float64 `json:"budget"`
	}

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})
		return
	}

	query := `
	UPDATE projects
	SET name=$1, budget=$2
	WHERE project_id=$3
	`

	_, err := config.DB.Exec(
		query,
		project.Name,
		project.Budget,
		id,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Project updated successfully",
	})
}

// ==========================
// DELETE PROJECT
// ==========================

func DeleteProject(c *gin.Context) {

	// Get ID from URL
	id := c.Param("id")

	// SQL delete query
	query := `
	DELETE FROM projects
	WHERE project_id=$1
	`

	_, err := config.DB.Exec(query, id)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Success response
	c.JSON(200, gin.H{
		"message": "Project deleted successfully",
	})
}