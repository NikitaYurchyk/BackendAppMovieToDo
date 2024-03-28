# Movie-API REST API

## Overview
The MovieList REST API Service is a Golang-based application designed for efficient and robust retrieval of movie data. Built with a focus on security and performance, this service utilizes Chi for routing, PostgreSQL for database operations, and JWT-based authorization for secure access.

## Getting Started

### Prerequisites
- Go (1.21 or later)
- Docker

### Setup

#### Database and Docker
1. **Start PostgreSQL Container**
   - Run the `create_docker` command to create and start a Docker container for PostgreSQL:
     ```
     make create_docker
     ```
   - Use `reuse_docker` to start the container if it already exists:
     ```
     make reuse_docker
     ```

2. **Environment Variables**
   - Ensure the following environment variables are set in your `.env` file or your environment:
     ```
     DB_HOST=127.0.0.1
     DB_PASSWORD=your_password
     DB_SSLMODE=disable
     DB_PORT=5432
     DB_USERNAME="postgres"
     DB_DNAME="postgres"
     DB_DRIVERNAME="postgres"
     ```

#### Application
1. **Build the Service**
   - Navigate to the project directory and build the application:
     ```
     go build -o movieListService
     ```

2. **Run the Service**
   - Start the service with:
     ```
     ./movieListService
     ```

### Usage

#### Endpoints
- **Get All Movies**: `GET /api`
- **Create New Movie**: `POST /api`
- **Get Movie by ID**: `GET /api/movies/{movieId}`
- **Update Movie by ID**: `PUT /api/movies/{movieId}`
- **Delete Movie by ID**: `DELETE /api/movies/{movieId}`



### Cleaning Up

- **Stop and Remove Docker Container**
  ```
  make clean
  ```

- **Remove Docker Image (Optional)**
  ```
  make clean-image
  ```


