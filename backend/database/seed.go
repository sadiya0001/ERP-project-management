package database

import (
	"log"
	"time"

	"erp-project-management/models"
)

// Seed populates the database with initial data for development.
func Seed() {
	var userCount int64
	DB.Model(&models.User{}).Count(&userCount)
	if userCount > 0 {
		log.Println("Database already seeded, skipping...")
		return
	}

	log.Println("Seeding database...")

	// ---- Users ----
	adminHash, _ := models.HashPassword("admin123")
	managerHash, _ := models.HashPassword("manager123")
	memberHash, _ := models.HashPassword("member123")

	users := []models.User{
		{
			Email:        "admin@erp.com",
			PasswordHash: adminHash,
			FirstName:    "Admin",
			LastName:     "User",
			Role:         "admin",
			Phone:        "+1-555-0100",
			Nationality:  "American",
			Designation:  "System Administrator",
			Skills:       "System Administration, DevOps, Go, PostgreSQL",
		},
		{
			Email:        "sarah.johnson@erp.com",
			PasswordHash: managerHash,
			FirstName:    "Sarah",
			LastName:     "Johnson",
			Role:         "manager",
			Phone:        "+1-555-0101",
			Nationality:  "Canadian",
			Designation:  "Project Manager",
			Skills:       "Project Management, Agile, Scrum, Leadership",
		},
		{
			Email:        "james.wilson@erp.com",
			PasswordHash: memberHash,
			FirstName:    "James",
			LastName:     "Wilson",
			Role:         "member",
			Phone:        "+1-555-0102",
			Nationality:  "British",
			Designation:  "Senior Developer",
			Skills:       "Go, React, TypeScript, PostgreSQL, Docker",
		},
		{
			Email:        "emily.chen@erp.com",
			PasswordHash: memberHash,
			FirstName:    "Emily",
			LastName:     "Chen",
			Role:         "member",
			Phone:        "+1-555-0103",
			Nationality:  "Chinese",
			Designation:  "UI/UX Designer",
			Skills:       "Figma, Adobe XD, CSS, Vue.js, User Research",
		},
		{
			Email:        "michael.brown@erp.com",
			PasswordHash: memberHash,
			FirstName:    "Michael",
			LastName:     "Brown",
			Role:         "member",
			Phone:        "+1-555-0104",
			Nationality:  "Australian",
			Designation:  "QA Engineer",
			Skills:       "Testing, Selenium, Cypress, CI/CD, Python",
		},
	}

	for i := range users {
		if err := DB.Create(&users[i]).Error; err != nil {
			log.Printf("Failed to seed user %s: %v", users[i].Email, err)
		}
	}

	// ---- Projects ----
	startDate1 := time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)
	endDate1 := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	startDate2 := time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC)
	endDate2 := time.Date(2026, 9, 30, 0, 0, 0, 0, time.UTC)
	startDate3 := time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC)
	endDate3 := time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)
	startDate4 := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)
	endDate4 := time.Date(2026, 4, 30, 0, 0, 0, 0, time.UTC)

	projects := []models.Project{
		{
			Title:       "ERP System Development",
			ProjectType: "Software Development",
			Description: "Main ERP system with modules for HR, Finance, Inventory, and Project Management.",
			Status:      "active",
			StartDate:   &startDate1,
			EndDate:     &endDate1,
			CreatedBy:   1,
		},
		{
			Title:       "Mobile App Redesign",
			ProjectType: "UI/UX Design",
			Description: "Complete redesign of the company mobile application with modern UI patterns.",
			Status:      "active",
			StartDate:   &startDate2,
			EndDate:     &endDate2,
			CreatedBy:   2,
		},
		{
			Title:       "Cloud Infrastructure Migration",
			ProjectType: "DevOps",
			Description: "Migrate existing on-premises infrastructure to AWS cloud services.",
			Status:      "active",
			StartDate:   &startDate3,
			EndDate:     &endDate3,
			CreatedBy:   1,
		},
		{
			Title:       "API Documentation Portal",
			ProjectType: "Documentation",
			Description: "Build a comprehensive API documentation portal for all internal services.",
			Status:      "completed",
			StartDate:   &startDate4,
			EndDate:     &endDate4,
			CreatedBy:   2,
		},
	}

	for i := range projects {
		if err := DB.Create(&projects[i]).Error; err != nil {
			log.Printf("Failed to seed project %s: %v", projects[i].Title, err)
		}
	}

	// ---- Project Members ----
	members := []models.ProjectMember{
		{ProjectID: 1, UserID: 1, Role: "team_lead"},
		{ProjectID: 1, UserID: 2, Role: "developer"},
		{ProjectID: 1, UserID: 3, Role: "developer"},
		{ProjectID: 1, UserID: 4, Role: "designer"},
		{ProjectID: 1, UserID: 5, Role: "tester"},
		{ProjectID: 2, UserID: 2, Role: "team_lead"},
		{ProjectID: 2, UserID: 4, Role: "designer"},
		{ProjectID: 2, UserID: 3, Role: "developer"},
		{ProjectID: 3, UserID: 1, Role: "team_lead"},
		{ProjectID: 3, UserID: 3, Role: "developer"},
		{ProjectID: 3, UserID: 5, Role: "tester"},
		{ProjectID: 4, UserID: 2, Role: "team_lead"},
		{ProjectID: 4, UserID: 3, Role: "developer"},
	}

	for i := range members {
		if err := DB.Create(&members[i]).Error; err != nil {
			log.Printf("Failed to seed project member: %v", err)
		}
	}

	// ---- Tasks ----
	deadline1 := time.Date(2026, 3, 15, 0, 0, 0, 0, time.UTC)
	deadline2 := time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC)
	deadline3 := time.Date(2026, 4, 15, 0, 0, 0, 0, time.UTC)
	deadline4 := time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC)
	deadline5 := time.Date(2026, 5, 15, 0, 0, 0, 0, time.UTC)
	deadline6 := time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)
	deadline7 := time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC)
	deadline8 := time.Date(2026, 7, 1, 0, 0, 0, 0, time.UTC)

	user3 := uint(3)
	user4 := uint(4)
	user5 := uint(5)

	tasks := []models.Task{
		// Project 1 tasks
		{ProjectID: 1, Title: "Design database schema", Description: "Create the complete database schema for all ERP modules.", Status: "completed", Priority: "high", AssignedTo: &user3, CreatedBy: 1, Deadline: &deadline1, TimeSpent: 480},
		{ProjectID: 1, Title: "Implement authentication module", Description: "Build JWT-based authentication with role-based access control.", Status: "completed", Priority: "critical", AssignedTo: &user3, CreatedBy: 1, Deadline: &deadline2, TimeSpent: 960},
		{ProjectID: 1, Title: "Build project management API", Description: "Create REST API endpoints for project management module.", Status: "in_progress", Priority: "high", AssignedTo: &user3, CreatedBy: 2, Deadline: &deadline3, TimeSpent: 360},
		{ProjectID: 1, Title: "Design dashboard UI", Description: "Create wireframes and mockups for the main dashboard.", Status: "in_progress", Priority: "medium", AssignedTo: &user4, CreatedBy: 2, Deadline: &deadline4, TimeSpent: 240},
		{ProjectID: 1, Title: "Write unit tests", Description: "Create comprehensive unit tests for all API endpoints.", Status: "pending", Priority: "medium", AssignedTo: &user5, CreatedBy: 1, Deadline: &deadline5, TimeSpent: 0},
		{ProjectID: 1, Title: "Setup CI/CD pipeline", Description: "Configure GitHub Actions for automated testing and deployment.", Status: "pending", Priority: "low", AssignedTo: &user3, CreatedBy: 1, Deadline: &deadline6, TimeSpent: 0},

		// Project 2 tasks
		{ProjectID: 2, Title: "User research and interviews", Description: "Conduct user research sessions to identify pain points.", Status: "completed", Priority: "high", AssignedTo: &user4, CreatedBy: 2, Deadline: &deadline2, TimeSpent: 600},
		{ProjectID: 2, Title: "Create design system", Description: "Build a comprehensive design system with reusable components.", Status: "in_progress", Priority: "high", AssignedTo: &user4, CreatedBy: 2, Deadline: &deadline4, TimeSpent: 480},
		{ProjectID: 2, Title: "Implement new navigation", Description: "Code the redesigned navigation bar and sidebar.", Status: "pending", Priority: "medium", AssignedTo: &user3, CreatedBy: 2, Deadline: &deadline6, TimeSpent: 0},

		// Project 3 tasks
		{ProjectID: 3, Title: "AWS account setup", Description: "Set up AWS accounts, IAM roles, and VPC configuration.", Status: "in_progress", Priority: "critical", AssignedTo: &user3, CreatedBy: 1, Deadline: &deadline5, TimeSpent: 180},
		{ProjectID: 3, Title: "Database migration plan", Description: "Create a detailed plan for migrating PostgreSQL databases to RDS.", Status: "pending", Priority: "high", AssignedTo: &user3, CreatedBy: 1, Deadline: &deadline7, TimeSpent: 0},
		{ProjectID: 3, Title: "Load testing", Description: "Perform load testing on the cloud infrastructure.", Status: "pending", Priority: "medium", AssignedTo: &user5, CreatedBy: 1, Deadline: &deadline8, TimeSpent: 0},
	}

	for i := range tasks {
		if err := DB.Create(&tasks[i]).Error; err != nil {
			log.Printf("Failed to seed task %s: %v", tasks[i].Title, err)
		}
	}

	// ---- Work Logs ----
	workLogs := []models.WorkLog{
		{TaskID: 1, UserID: 3, Hours: 4.0, Description: "Designed initial schema for users, projects, and tasks tables.", LogDate: time.Date(2026, 2, 10, 0, 0, 0, 0, time.UTC)},
		{TaskID: 1, UserID: 3, Hours: 4.0, Description: "Added work_logs table and refined foreign key relationships.", LogDate: time.Date(2026, 2, 11, 0, 0, 0, 0, time.UTC)},
		{TaskID: 2, UserID: 3, Hours: 6.0, Description: "Implemented JWT token generation and validation.", LogDate: time.Date(2026, 2, 20, 0, 0, 0, 0, time.UTC)},
		{TaskID: 2, UserID: 3, Hours: 5.0, Description: "Added role-based middleware and login/register endpoints.", LogDate: time.Date(2026, 2, 21, 0, 0, 0, 0, time.UTC)},
		{TaskID: 2, UserID: 3, Hours: 5.0, Description: "Fixed edge cases in token refresh and password hashing.", LogDate: time.Date(2026, 2, 22, 0, 0, 0, 0, time.UTC)},
		{TaskID: 3, UserID: 3, Hours: 6.0, Description: "Built CRUD endpoints for projects with pagination.", LogDate: time.Date(2026, 3, 5, 0, 0, 0, 0, time.UTC)},
		{TaskID: 4, UserID: 4, Hours: 4.0, Description: "Created low-fidelity wireframes for the dashboard.", LogDate: time.Date(2026, 3, 10, 0, 0, 0, 0, time.UTC)},
		{TaskID: 7, UserID: 4, Hours: 5.0, Description: "Conducted 5 user interview sessions.", LogDate: time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC)},
		{TaskID: 7, UserID: 4, Hours: 5.0, Description: "Synthesised research findings and created persona documents.", LogDate: time.Date(2026, 3, 2, 0, 0, 0, 0, time.UTC)},
		{TaskID: 8, UserID: 4, Hours: 8.0, Description: "Created colour palette, typography scale, and spacing tokens.", LogDate: time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC)},
		{TaskID: 10, UserID: 3, Hours: 3.0, Description: "Created AWS organisation and configured root account security.", LogDate: time.Date(2026, 5, 5, 0, 0, 0, 0, time.UTC)},
	}

	for i := range workLogs {
		if err := DB.Create(&workLogs[i]).Error; err != nil {
			log.Printf("Failed to seed work log: %v", err)
		}
	}

	log.Println("Database seeded successfully!")
}
