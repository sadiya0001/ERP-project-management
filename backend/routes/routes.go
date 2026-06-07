package routes

import (
	"erp-project-management/config"
	"erp-project-management/handlers"
	"erp-project-management/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures all routes and middleware for the application.
func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Health check
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "ERP Project Management API is running",
		})
	})

	// Auth handler
	authHandler := handlers.NewAuthHandler(cfg)

	// Public routes
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		// Auth - profile
		protected.GET("/auth/profile", authHandler.GetProfile)

		// Projects
		protected.GET("/projects", handlers.GetProjects)
		protected.GET("/projects/:id", handlers.GetProject)
		protected.POST("/projects", handlers.CreateProject)
		protected.PUT("/projects/:id", handlers.UpdateProject)
		protected.DELETE("/projects/:id", handlers.DeleteProject)
		protected.POST("/projects/:id/members", handlers.AddProjectMember)
		protected.DELETE("/projects/:id/members/:userId", handlers.RemoveProjectMember)

		// Tasks
		protected.GET("/projects/:id/tasks", handlers.GetTasks)
		protected.POST("/projects/:id/tasks", handlers.CreateTask)
		protected.PUT("/tasks/:id", handlers.UpdateTask)
		protected.DELETE("/tasks/:id", handlers.DeleteTask)
		protected.PUT("/tasks/:id/assign", handlers.AssignTask)

		// Users
		protected.GET("/users", handlers.GetUsers)
		protected.GET("/users/:id", handlers.GetUser)
		protected.PUT("/users/:id", handlers.UpdateUser)

		// Work Logs
		protected.GET("/tasks/:id/worklogs", handlers.GetWorkLogs)
		protected.POST("/tasks/:id/worklogs", handlers.CreateWorkLog)

		// Dashboard
		protected.GET("/dashboard/stats", handlers.GetDashboardStats)

		// Integration endpoints (for other ERP modules)
		protected.GET("/integration/projects", handlers.GetIntegrationProjects)
		protected.GET("/integration/resources", handlers.GetIntegrationResources)
	}

	return r
}
