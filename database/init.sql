-- ERP Project Management Module - Database Initialization
-- Group 13 | Enterprise Resource Planning System

-- Create database (if not already created by Docker)
-- CREATE DATABASE erp_project_mgmt;

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ============================================
-- USERS TABLE
-- ============================================
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    role VARCHAR(50) DEFAULT 'member', -- admin, manager, member
    avatar VARCHAR(500),
    phone VARCHAR(20),
    nationality VARCHAR(100),
    designation VARCHAR(100),
    skills TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- PROJECTS TABLE
-- ============================================
CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    project_type VARCHAR(100),
    description TEXT,
    status VARCHAR(50) DEFAULT 'active', -- active, completed, on_hold, archived
    start_date DATE,
    end_date DATE,
    created_by INTEGER REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- PROJECT MEMBERS TABLE
-- ============================================
CREATE TABLE IF NOT EXISTS project_members (
    id SERIAL PRIMARY KEY,
    project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(100), -- team_lead, developer, designer, tester, manager
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(project_id, user_id)
);

-- ============================================
-- TASKS TABLE
-- ============================================
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE,
    title VARCHAR(500) NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'pending', -- pending, in_progress, completed, on_hold
    priority VARCHAR(50) DEFAULT 'medium', -- low, medium, high, critical
    assigned_to INTEGER REFERENCES users(id) ON DELETE SET NULL,
    created_by INTEGER REFERENCES users(id) ON DELETE SET NULL,
    deadline DATE,
    time_spent INTEGER DEFAULT 0, -- in minutes
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- WORK LOGS TABLE
-- ============================================
CREATE TABLE IF NOT EXISTS work_logs (
    id SERIAL PRIMARY KEY,
    task_id INTEGER REFERENCES tasks(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    hours FLOAT NOT NULL,
    description TEXT,
    log_date DATE DEFAULT CURRENT_DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- INDEXES
-- ============================================
CREATE INDEX IF NOT EXISTS idx_projects_status ON projects(status);
CREATE INDEX IF NOT EXISTS idx_projects_created_by ON projects(created_by);
CREATE INDEX IF NOT EXISTS idx_tasks_project_id ON tasks(project_id);
CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks(status);
CREATE INDEX IF NOT EXISTS idx_tasks_assigned_to ON tasks(assigned_to);
CREATE INDEX IF NOT EXISTS idx_project_members_project_id ON project_members(project_id);
CREATE INDEX IF NOT EXISTS idx_project_members_user_id ON project_members(user_id);
CREATE INDEX IF NOT EXISTS idx_work_logs_task_id ON work_logs(task_id);
CREATE INDEX IF NOT EXISTS idx_work_logs_user_id ON work_logs(user_id);

-- ============================================
-- SEED DATA
-- ============================================

-- Password for all users: admin123 (bcrypt hash)
-- $2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy

INSERT INTO users (email, password_hash, first_name, last_name, role, avatar, phone, nationality, designation, skills) VALUES
('admin@erp.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'Yash', 'Ghori', 'admin', 'https://i.pravatar.cc/150?img=11', '+94 71234 5678', 'Sri Lanka', 'Project Manager', 'Project Management, Agile, Scrum'),
('sarah@erp.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'Sarah', 'Johnson', 'manager', 'https://i.pravatar.cc/150?img=5', '+94 71234 5679', 'Sri Lanka', 'Team Lead', 'Vue.js, React, Frontend Development'),
('mike@erp.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'Mike', 'Chen', 'member', 'https://i.pravatar.cc/150?img=12', '+94 71234 5680', 'Sri Lanka', 'Backend Developer', 'Go, Python, PostgreSQL'),
('emily@erp.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'Emily', 'Davis', 'member', 'https://i.pravatar.cc/150?img=9', '+94 71234 5681', 'Sri Lanka', 'UI/UX Designer', 'Figma, Adobe XD, CSS'),
('alex@erp.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'Alex', 'Kumar', 'member', 'https://i.pravatar.cc/150?img=15', '+94 71234 5682', 'Sri Lanka', 'Full Stack Developer', 'Vue.js, Go, Docker')
ON CONFLICT (email) DO NOTHING;

INSERT INTO projects (title, project_type, description, status, start_date, end_date, created_by) VALUES
('ERP System Module', 'Enterprise', 'Enterprise Resource Planning system with integrated modules for project management, HR, finance, and administration.', 'active', '2024-01-15', '2024-06-30', 1),
('Mobile App Redesign', 'Mobile', 'Complete redesign of the mobile application with modern UI/UX patterns and improved performance.', 'active', '2024-02-01', '2024-05-15', 1),
('API Gateway Setup', 'Infrastructure', 'Setting up centralized API gateway for microservices architecture with rate limiting and authentication.', 'on_hold', '2024-03-01', '2024-04-30', 2),
('Customer Portal', 'Web', 'Self-service customer portal for order tracking, support tickets, and account management.', 'completed', '2023-10-01', '2024-01-31', 1),
('Data Analytics Dashboard', 'Analytics', 'Real-time analytics dashboard for business intelligence and reporting across all ERP modules.', 'active', '2024-03-15', '2024-08-30', 2),
('DevOps Pipeline', 'Infrastructure', 'CI/CD pipeline setup with automated testing, building, and deployment using Docker and Kubernetes.', 'active', '2024-02-15', '2024-04-30', 3)
ON CONFLICT DO NOTHING;

INSERT INTO project_members (project_id, user_id, role) VALUES
(1, 1, 'team_lead'),
(1, 2, 'developer'),
(1, 3, 'developer'),
(1, 4, 'designer'),
(1, 5, 'developer'),
(2, 2, 'team_lead'),
(2, 4, 'designer'),
(2, 5, 'developer'),
(3, 3, 'team_lead'),
(3, 1, 'manager'),
(4, 1, 'team_lead'),
(4, 2, 'developer'),
(4, 3, 'developer'),
(5, 2, 'team_lead'),
(5, 3, 'developer'),
(5, 5, 'developer'),
(6, 3, 'team_lead'),
(6, 5, 'developer')
ON CONFLICT DO NOTHING;

INSERT INTO tasks (project_id, title, description, status, priority, assigned_to, created_by, deadline, time_spent) VALUES
(1, 'Design database schema for Project Management', 'Create comprehensive database schema with all required tables, indexes, and relationships for the project management module.', 'completed', 'high', 3, 1, '2024-02-15', 480),
(1, 'Implement user authentication with JWT', 'Set up JWT-based authentication system with login, register, and token refresh endpoints.', 'completed', 'critical', 3, 1, '2024-02-20', 360),
(1, 'Create REST API for projects CRUD', 'Implement full CRUD operations for projects including create, read, update, and delete endpoints.', 'in_progress', 'high', 3, 2, '2024-03-01', 240),
(1, 'Build Vue.js dashboard component', 'Create the main dashboard view with project overview cards, task statistics pie chart, and team members grid.', 'in_progress', 'high', 2, 1, '2024-03-10', 300),
(1, 'Implement task management frontend', 'Build the task list view with filtering, status updates, time tracking, and task assignment features.', 'pending', 'medium', 5, 1, '2024-03-15', 0),
(1, 'Setup Docker containerization', 'Dockerize frontend, backend, and database services with docker-compose for easy deployment.', 'pending', 'medium', 5, 1, '2024-03-20', 0),
(1, 'Write API documentation with Swagger', 'Document all API endpoints with request/response schemas using Swagger/OpenAPI specification.', 'pending', 'low', 2, 1, '2024-03-25', 0),
(2, 'Create wireframes for mobile redesign', 'Design wireframes for all major screens including home, profile, notifications, and settings.', 'completed', 'high', 4, 2, '2024-02-15', 600),
(2, 'Implement new navigation system', 'Build the new bottom tab navigation with smooth animations and gesture support.', 'in_progress', 'high', 5, 2, '2024-03-01', 180),
(2, 'Build notification center', 'Create in-app notification system with real-time updates and push notification integration.', 'on_hold', 'medium', 5, 2, '2024-03-15', 60),
(3, 'Configure API Gateway routing', 'Set up routing rules for all microservices through the API gateway with path-based routing.', 'on_hold', 'critical', 3, 3, '2024-03-15', 120),
(4, 'Build order tracking module', 'Implement real-time order tracking with status updates and delivery estimation.', 'completed', 'high', 2, 1, '2024-01-15', 720),
(4, 'Create support ticket system', 'Build customer support ticket creation, tracking, and resolution workflow.', 'completed', 'high', 3, 1, '2024-01-20', 540),
(5, 'Design analytics data models', 'Create data models and aggregation pipelines for business intelligence reporting.', 'in_progress', 'high', 3, 2, '2024-04-15', 200),
(5, 'Build chart components', 'Implement reusable chart components for various data visualizations using Chart.js.', 'pending', 'medium', 5, 2, '2024-04-30', 0)
ON CONFLICT DO NOTHING;

INSERT INTO work_logs (task_id, user_id, hours, description, log_date) VALUES
(1, 3, 4.0, 'Designed initial schema with users and projects tables', '2024-02-10'),
(1, 3, 4.0, 'Added tasks, work_logs tables and indexes', '2024-02-11'),
(2, 3, 3.0, 'Implemented JWT token generation and validation', '2024-02-15'),
(2, 3, 3.0, 'Added login and register endpoints with bcrypt', '2024-02-16'),
(3, 3, 4.0, 'Created project CRUD handlers with GORM', '2024-02-25'),
(4, 2, 5.0, 'Built dashboard layout with sidebar and topbar', '2024-03-01'),
(4, 2, 5.0, 'Implemented pie chart and project cards', '2024-03-02'),
(8, 4, 6.0, 'Completed all wireframe designs in Figma', '2024-02-12'),
(8, 4, 4.0, 'Revised wireframes based on feedback', '2024-02-14'),
(9, 5, 3.0, 'Started implementing bottom tab navigation', '2024-02-28'),
(12, 2, 8.0, 'Built complete order tracking with real-time updates', '2024-01-10'),
(12, 2, 4.0, 'Added delivery estimation algorithm', '2024-01-12'),
(13, 3, 6.0, 'Implemented ticket creation and listing', '2024-01-15'),
(13, 3, 3.0, 'Added ticket resolution workflow', '2024-01-18')
ON CONFLICT DO NOTHING;

-- ============================================
-- FUNCTIONS / TRIGGERS (optional enhancements)
-- ============================================

-- Auto-update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_projects_updated_at BEFORE UPDATE ON projects
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_tasks_updated_at BEFORE UPDATE ON tasks
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
