# 🏗️ ERP Project Management Module

## Group 13 | Enterprise Resource Planning System
### IT2305 - Enterprise System Development & Integration

---

## 📋 Project Description

This is the **Project Management Module** of the Enterprise Resource Planning (ERP) System. It provides comprehensive project and task management capabilities including:

- **Project Management**: Create, update, and track projects with team assignments
- **Task Management**: Assign tasks, track progress, set deadlines, and manage workloads
- **Team Collaboration**: Manage team members, roles, and permissions
- **Work Logging**: Track time spent on tasks and generate performance reports
- **Dashboard Analytics**: Visual overview of project status, task distribution, and team performance
- **User Profile Management**: Edit profiles, manage settings, and view work statistics

---

## 🛠️ Technology Stack (Alternative Stack - Bonus)

| Layer      | Technology         |
|------------|--------------------|
| Frontend   | Vue.js 3 (Vite)    |
| Backend    | Go (Gin Framework) |
| Database   | PostgreSQL 15      |
| ORM        | GORM               |
| Auth       | JWT                |
| Container  | Docker + Docker Compose |

---

## 🚀 How to Run

### Prerequisites
- Docker & Docker Compose installed on your machine

### Steps

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd "Erp system"
   ```

2. **Start the system**
   ```bash
   docker-compose up --build
   ```

3. **Access the application**
   - 🌐 Frontend: [http://localhost:3000](http://localhost:3000)
   - 🔌 Backend API: [http://localhost:8080](http://localhost:8080)
   - 🗄️ Database: `localhost:5432`

4. **Default Login Credentials**
   ```
   Email: admin@erp.com
   Password: admin123
   ```

---

## 📡 Port Details

| Service    | Port  |
|------------|-------|
| Frontend   | 3000  |
| Backend    | 8080  |
| PostgreSQL | 5432  |

---

## 📚 API Endpoints

### Authentication
| Method | Endpoint              | Description          |
|--------|-----------------------|----------------------|
| POST   | `/api/auth/register`  | Register new user    |
| POST   | `/api/auth/login`     | Login & get JWT      |
| GET    | `/api/auth/profile`   | Get current profile  |

### Projects
| Method | Endpoint                            | Description            |
|--------|-------------------------------------|------------------------|
| GET    | `/api/projects`                     | List all projects      |
| GET    | `/api/projects/:id`                 | Get project details    |
| POST   | `/api/projects`                     | Create project         |
| PUT    | `/api/projects/:id`                 | Update project         |
| DELETE | `/api/projects/:id`                 | Delete project         |
| POST   | `/api/projects/:id/members`         | Add project member     |
| DELETE | `/api/projects/:id/members/:userId` | Remove project member  |

### Tasks
| Method | Endpoint                     | Description          |
|--------|------------------------------|----------------------|
| GET    | `/api/projects/:id/tasks`    | List project tasks   |
| POST   | `/api/projects/:id/tasks`    | Create task          |
| PUT    | `/api/tasks/:id`             | Update task          |
| DELETE | `/api/tasks/:id`             | Delete task          |
| PUT    | `/api/tasks/:id/assign`      | Assign task to user  |

### Users
| Method | Endpoint         | Description      |
|--------|------------------|------------------|
| GET    | `/api/users`     | List all users   |
| GET    | `/api/users/:id` | Get user details |
| PUT    | `/api/users/:id` | Update profile   |

### Work Logs
| Method | Endpoint                    | Description        |
|--------|-----------------------------|--------------------|
| GET    | `/api/tasks/:id/worklogs`   | Get task work logs |
| POST   | `/api/tasks/:id/worklogs`   | Add work log       |

### Dashboard
| Method | Endpoint               | Description            |
|--------|------------------------|------------------------|
| GET    | `/api/dashboard/stats` | Get dashboard stats    |

### Integration (For ERP System)
| Method | Endpoint                      | Description              |
|--------|-------------------------------|--------------------------|
| GET    | `/api/integration/projects`   | Shared project data      |
| GET    | `/api/integration/resources`  | Resource allocation data |

---

## 🔗 Integration Details

### Shared Identifiers
- **Project ID**: Used across ERP modules to reference projects
- **User ID**: Shared user identity across modules
- **Task ID**: Unique task reference

### Data Format
All API responses use **JSON** format.

### Integration with Other ERP Modules
- **HR Module** (Group 10): Shares employee/user data for project assignments
- **Finance Module** (Group 9): Provides project budget and cost tracking data
- **Admin Module** (Group 14): Shares system administration and access control
- **Asset Management** (Group 11): Links project resources with asset tracking
- **Procurement Module** (Group 14): Connects project material requirements

### Example Integration Flow
```
HR Module → Provides employee list → Project Management assigns members
Project Management → Reports resource allocation → Finance Module tracks costs
Project Management → Lists required assets → Asset Management tracks usage
```

---

## 📁 Project Structure

```
Erp system/
├── docker-compose.yml          # Docker orchestration
├── .env.example                # Environment variables template
├── README.md                   # This file
├── database/
│   └── init.sql                # Database initialization
├── backend/
│   ├── Dockerfile
│   ├── main.go                 # Entry point
│   ├── go.mod / go.sum
│   ├── config/                 # Configuration
│   ├── models/                 # Database models
│   ├── handlers/               # API handlers
│   ├── middleware/              # JWT auth middleware
│   ├── database/               # DB connection & seeds
│   └── routes/                 # Route definitions
└── frontend/
    ├── Dockerfile
    ├── nginx.conf
    ├── src/
    │   ├── views/              # Page components
    │   ├── components/         # Reusable components
    │   ├── layouts/            # Layout wrappers
    │   ├── stores/             # Pinia stores
    │   ├── services/           # API service
    │   └── router/             # Vue Router config
    └── package.json
```



## 📝 License

This project is developed as part of IT2305 Continuous Assessment 1 at the University.
