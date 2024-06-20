# GO-API-REST

An API in GO, with CRUD operations for users and products, following the architecture proposed in Robert Cecil Martin's Clean Architecture. It utilizes GORM, Gin Gonic, and SQL.

## Table of Contents

- [Description](#description)
- [Installation](#installation)
- [Contact](#contact)

## Description

This project is a RESTful API developed in Go, implementing CRUD (Create, Read, Update, Delete) operations for users and products. The API uses Gin Gonic for routing, GORM for database interactions, and SQL as the underlying database.

## Installation

Instructions to install and set up the project.

```bash
# Clone the repository
git clone https://github.com/JPSBarbosa/GO-API-REST.git

# Navigate to the project directory
cd ./cmd/api

# Install dependencies
go mod tidy

# Configure your database (adjust the .env file as needed)
# Example .env configuration:
# DB_HOST=localhost
# DB_USER=your_user
# DB_PASSWORD=your_password
# DB_NAME=your_database
# DB_PORT=3306

# Start the application
go run main.go

```

## Endpoints

### Users

- `POST /users` - Create a new user
- `GET /users/:id` - Get a specific user
- `PUT /users/:id` - Update a specific user
- `DELETE /users/:id` - Delete a specific user

### Products

- `POST /products` - Create a new product
- `GET /products/:id` - Get details of a specific product
- `PUT /products/:id` - Update a specific product
- `DELETE /products/:id` - Delete a specific product

## Contact

João Pedro S. Barbosa - jpbarbosa.dev@gmail.com
