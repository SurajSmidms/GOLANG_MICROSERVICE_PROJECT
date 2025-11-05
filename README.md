🚀 Golang Microservices Project – Smarket 🛒

Hey there 👋
This project is my hands-on journey into building microservices using Golang.
I wanted to understand how independent services work together in real-world backend systems — so I built this mini e-commerce-style project called Smarket.

💡 What I’ve Built

I created four microservices that work together through an API Gateway — just like in real production systems.

🧱 The Services
Service	Port	Description
🧑‍💻 Auth Service	8081	Handles user registration, login, JWT authentication & token refresh.
📦 Product Service	8083	Manages all product data — add, update, delete, and view products.
🧾 Order Service	8082	Handles customer orders and stores order info.
💳 Payment Service	8084	Records payments for placed orders.
🌐 API Gateway	8080	Acts as the single entry point and routes requests securely to each service.

Each service runs independently, connects to MySQL, and communicates through REST APIs.

🧠 What I Learned

How to design and build microservice architecture in Golang

Connecting services to MySQL using GORM

Using Gin framework for creating REST APIs

Implementing JWT authentication (Access + Refresh Tokens)

Organizing project structure for scalability and clarity

Managing configurations with a shared .env file

Using API Gateway to connect and secure all services

⚙️ Technologies Used

Golang 🦫

Gin Framework

MySQL + GORM

JWT Authentication

godotenv for environment setup

🧩 How It Works

Each microservice runs on a different port and performs its specific task.
The API Gateway connects all these services and routes incoming requests — just like a real-world backend system.

🧰 How to Run

Clone this repository

git clone https://github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT.git
cd GOLANG_MICROSERVICE_PROJECT


Create a .env file in the root folder:

DB_USER=root
DB_PASS=root
DB_NAME=smarket
DB_HOST=127.0.0.1
DB_PORT=3306
ACCESS_SECRET=youraccesssecret
REFRESH_SECRET=yourrefreshsecret


Start MySQL and create the database:

CREATE DATABASE smarket;


Run each service in a new terminal:

cd auth-service && go run main.go
cd product-service && go run main.go
cd order-service && go run main.go
cd payment-service && go run main.go
cd api-gateway && go run main.go

🌟 Why This Project Matters

This project helped me:
✅ Understand microservices in depth
✅ Learn how to connect services cleanly in Go
✅ Implement real-world authentication and routing
✅ Build confidence in scalable backend development

💬 Final Thoughts

This project is a big step in my journey toward becoming a backend engineer with microservice architecture expertise.
It gave me practical experience with Golang, API design, database integration, and service orchestration — all from scratch.
