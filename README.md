# 🧪 Dynamic Reminder System with Audit Trail

A flexible notification engine built with **Go** that evaluates configurable reminder rules against scheduled tasks, triggers alerts, and logs audit trails.

## Technology Stack

- **Go (Golang)** with **Gin** web framework
- **GORM** Object-Relational Mapping
- **PostgreSQL** relational database
- **robfig/cron** for background task scheduling
- **godotenv** for environment management

## Workspace Layout

```text
reminder-app/
├── cmd/
│   └── main.go                 # Application bootstrap & entrypoint
├── internal/
│   ├── database/               # Database connection & seed script
│   ├── handlers/               # HTTP request handlers
│   ├── models/                 # GORM entities & schema definitions
│   ├── repository/             # Data access layer
│   ├── routes/                 # Endpoint routing configuration
│   ├── scheduler/              # Cron job & rule evaluation engine
│   └── services/               # Core business logic layer
├── .env                        # Local environment variables
├── go.mod                      # Module specifications
└── README.md                   # Project documentation
```

## Setup & Running

### 1. Requirements
- Installed Go (v1.20+)
- Running PostgreSQL database instance

### 2. Environment Configuration
Create a `.env` file in the project root:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=task_reminder_db
DB_PORT=5432
PORT=8080
```

### 3. Dependencies & Execution
```bash
# Download required modules
go mod download

# Start the application server
go run cmd/main.go
```
*Note: The server automatically migrates PostgreSQL tables and inserts sample tasks on launch.*

## API Specifications

### Reminder Rules Management (`/api/reminder-rules`)
- `POST   /api/reminder-rules` — Define a new reminder rule
- `GET    /api/reminder-rules` — Retrieve all configured rules
- `GET    /api/reminder-rules/:id` — Retrieve rule details by ID
- `PUT    /api/reminder-rules/:id` — Update an existing rule
- `DELETE /api/reminder-rules/:id` — Remove a rule
- `PATCH  /api/reminder-rules/:id/toggle` — Enable or disable a rule

### Audit Trail Inspection (`/api/activity-logs`)
- `GET    /api/activity-logs` — List all activity & execution audit logs
- `GET    /api/activity-logs/:id` — Inspect a specific log entry by ID

## Background Reminder Engine
An autonomous background worker runs every minute to scan pending tasks against active rules. When condition criteria (`before_due`, `overdue`, `on_due`) are satisfied, the engine prints a reminder notification to stdout and logs the execution in the audit trail.
