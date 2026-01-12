**Task Service API**

A simple Go-based task processing service with:

- REST API

- Background worker pool

- SQLite persistence

- Docker support

- Graceful shutdown

This project demonstrates clean architecture, concurrency, and real-world backend patterns.

**📌 Features**

- Create tasks via HTTP API

- Asynchronous task processing using worker goroutines

- Task lifecycle:

    - pending → processing → completed

- SQLite database for persistence

- Graceful shutdown with context & signals

- Dockerized for easy deployment

**🏗 Project Structure**

<img width="499" height="402" alt="image" src="https://github.com/user-attachments/assets/fa871d72-8f23-4a2a-ba38-14da6d532faa" />


**🚀 Getting Started (Local)**

**1️⃣ Install dependencies**
go mod tidy

**2️⃣ Run the service**
go run ./cmd/api


Server will start on:

http://localhost:8080

**📡 API Endpoints**
➕ Create Task
POST /tasks
Content-Type: application/json


Request body

{
  "title": "My first task"
}


Response

{
  "id": "uuid",
  "title": "My first task",
  "status": "pending",
  "created_at": "2026-01-12T01:37:19Z"
}

**🔍 Get Task**
GET /task?id=<task_id>


Response

{
  "id": "uuid",
  "title": "My first task",
  "status": "completed",
  "created_at": "2026-01-12T01:37:19Z"
}

**⚙ Background Workers**

Workers are started on boot:

taskService.StartWorkers(3)


Tasks are processed asynchronously

Status updates are persisted in SQLite

**🗄 Database**

SQLite database file: tasks.db

Auto-created on startup

Table schema:

CREATE TABLE tasks (
  id TEXT PRIMARY KEY,
  title TEXT,
  status TEXT,
  created_at DATETIME
);

**🐳 Docker**
Build image
docker build -t task-service .

Run container
docker run -p 8080:8080 task-service

**🛑 Graceful Shutdown**

- Handles SIGINT / SIGTERM

- Stops HTTP server

- Waits for workers to finish

- Safely closes resources

**🧠 Design Principles**

- Clean separation of concerns

- Repository pattern

- Dependency injection

- Concurrent worker pool

- Production-ready structure

**📈 Possible Improvements**

- Pagination for task listing

- Task retry / failure handling

- Authentication

- Metrics & health checks

- PostgreSQL / Redis support

- Docker Compose

**👨‍💻 Author**

Aiman Danish
Backend Engineer (Go)
