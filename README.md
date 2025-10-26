
# 🧩 Task Manager RESTful API

A simple yet complete **Task Manager API** built with **Go (Gin Framework)** and **GORM**, featuring **JWT Authentication**, CRUD operations, and deployment using **ngrok**.

---

## 📚 Features

✅ User Authentication (Register / Login)  
✅ JWT-based route protection  
✅ CRUD operations for Tasks  
✅ SQLite database (easy setup)  
✅ Environment variables support (.env)  
✅ Local deployment using ngrok  

---

## 🗂️ Folder Structure

```

task-manager/
│
├── main.go
├── go.mod
├── .env
│
├── config/
│   └── db.go
│
├── models/
│   ├── task.go
│   └── user.go
│
├── controllers/
│   ├── authController.go
│   └── taskController.go
│
├── routes/
│   ├── authRoutes.go
│   └── taskRoutes.go
│
└── middleware/
└── authMiddleware.go

````

---

## ⚙️ Setup Instructions

### 1️⃣ Clone the repository
```bash
git clone https://github.com/shivamg7753/task-manager.git
cd task-manager
````

### 2️⃣ Initialize Go Modules

```bash
go mod tidy
```

### 3️⃣ Create `.env` file in root

```env
JWT_SECRET=supersecretkey
PORT=8080
```

### 4️⃣ Run the server

```bash
go run main.go
```

Server will start at:
👉 [http://localhost:8080](http://localhost:8080)

---

## 🔐 Authentication Endpoints

### ➕ Register User

**POST** `/auth/register`

**Body:**

```json
{
  "username": "shivam",
  "password": "12345"
}
```

**Response:**

```json
{
  "message": "User registered successfully!"
}
```

---

### 🔑 Login User

**POST** `/auth/login`

**Body:**

```json
{
  "username": "shivam",
  "password": "12345"
}
```

**Response:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR..."
}
```

Copy the token for authorization.

---

## 📋 Task Endpoints (Protected)

Add Header:

```
Authorization: Bearer <your_token>
```

### ➕ Create Task

**POST** `/tasks`

```json
{
  "title": "Learn Go",
  "description": "Finish Gin project",
  "completed": false
}
```

---

### 📄 Get All Tasks

**GET** `/tasks`

**Response:**

```json
[
  {
    "ID": 1,
    "Title": "Learn Go",
    "Description": "Finish Gin project",
    "Completed": false
  }
]
```

---

### ✏️ Update Task

**PUT** `/tasks/:id`

```json
{
  "title": "Learn Gin",
  "description": "Build REST API",
  "completed": true
}
```

---

### ❌ Delete Task

**DELETE** `/tasks/:id`

---

## 🌍 Deploy with ngrok

### 1️⃣ Install ngrok

[Download Here](https://ngrok.com/download)

### 2️⃣ Connect ngrok

```bash
ngrok config add-authtoken <YOUR_TOKEN>
```

### 3️⃣ Start Server

```bash
go run main.go
```

### 4️⃣ Run ngrok tunnel

```bash
ngrok http 8080
```

Now your API is publicly accessible:
`https://abcd1234.ngrok.io`

---

## 🧠 Tech Stack

| Component  | Tech Used         |
| ---------- | ----------------- |
| Language   | Go                |
| Framework  | Gin               |
| Database   | SQLite (via GORM) |
| Auth       | JWT               |
| Deployment | ngrok             |

---

## 🧾 Example `.env` File

```env
JWT_SECRET=supersecretkey
PORT=8080
```

---

## 🧰 Useful Commands

| Command           | Description                |
| ----------------- | -------------------------- |
| `go run main.go`  | Run project                |
| `go mod tidy`     | Sync dependencies          |
| `ngrok http 8080` | Expose local API to public |

---

## 🧑‍💻 Author

**Shivam Gupta**
🚀 Passionate Go Developer | MERN Enthusiast
---

## 📜 License

This project is open-source and available under the [MIT License](LICENSE).
