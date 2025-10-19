# 🐹 Yasin Todo API

**Prepared by:** Yasin Sahyar  
**Technology Stack:** Go (Golang), MongoDB, Chi Router, Swagger, Docker  

A simple and complete REST API project built with **Go** and **MongoDB**, perfectly matching the “Go REST API Assignment”.  
Supports CRUD operations, advanced features, Swagger documentation, and Docker deployment.

---

## 🚀 Features Overview

| Step | Requirement | Status | Description |
|------|------------|--------|-------------|
| Step 1 | Study Go REST tutorials | ✅ | Tutorials reviewed before implementation. |
| Step 2 | CRUD operations | ✅ | Full CRUD for tasks (GET, POST, PUT, DELETE). |
| Step 3 | Advanced features (3+) | ✅ | Filtering, Pagination, Validation implemented. |
| Step 4 | API documentation | ✅ | Swagger UI available at `/swagger/index.html`. |
| Step 5 | Docker deployment | ✅ | Dockerfile provided; containerized for cloud deployment. |

---

## ⚙️ Technologies Used

- **Language:** Go 1.22+  
- **Framework:** Chi Router  
- **Database:** MongoDB  
- **API Documentation:** Swagger (`swaggo/swag`, `http-swagger`)  
- **Containerization:** Docker  

---

## 🗂 Project Structure


yasin-todo-api/
├── main.go
├── go.mod
├── go.sum
├── README.md
├── models/
│ └── task.go
├── handlers/
│ └── taskHandler.go
├── routes/
│ └── taskRoutes.go
├── static/
│ └── home.tpl
└── docs/
└── swagger files (auto-generated)



### 1️⃣ Clone the Repository
```bash
git clone https://github.com/<your-username>/yasin-todo-api.git
cd yasin-todo-api
```

### 2️⃣ Start MongoDB Locally

Ensure MongoDB runs at:

mongodb://localhost:27017

### 3️⃣ Install Dependencies
go mod tidy

### 4️⃣ Generate Swagger Docs
swag init

### 5️⃣ Run the Application
go run main.go


Server runs at:

http://localhost:9000

### 6️⃣ Open Swagger UI
http://localhost:9000/swagger/index.html

🧩 API Endpoints
Method	Endpoint	Description
GET	/tasks	Retrieve all tasks
GET	/tasks/{id}	Retrieve a single task by ID
POST	/tasks	Create a new task
PUT	/tasks/{id}	Update an existing task
DELETE	/tasks/{id}	Delete a task
🔍 Advanced Features

✅ Filtering: Filter tasks by status (completed / pending).
✅ Pagination: Supports page and page_size query parameters.
✅ Validation: Input validation for POST and PUT requests.

🐳 Docker Deployment
Build Docker Image
docker build -t yasin-todo-api .

Run Container
docker run -p 9000:9000 yasin-todo-api


Visit Swagger UI:

http://localhost:9000/swagger/index.html

🧾 Example Task Object
{
  "title": "Finish assignment",
  "description": "Complete Go REST API project",
  "status": "pending"
}

🧠 Notes

MongoDB Database: yasin_todo

Collection: tasks

Default MongoDB URL: mongodb://localhost:27017 (can override with MONGO_URL environment variable)

🏁 Summary

This project covers all assignment steps:

✅ Step 1–3: Tutorials + CRUD + Advanced Features

✅ Step 4: Swagger Documentation

✅ Step 5: Docker Deployment

Perfect for demonstrating understanding of Go REST APIs, MongoDB integration, and best practices for API development.