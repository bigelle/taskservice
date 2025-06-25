# IO Task Tracker Service

A high-performance Go service for tracking and monitoring I/O bound tasks through a RESTful HTTP API. Built with scalability and extensibility in mind, this service provides real-time insights into your I/O operations with minimal overhead.

## 🚀 Features

- **Real-time I/O Task Monitoring** - Track database queries, file operations, network requests, and more
- **RESTful HTTP API** - Clean, intuitive endpoints for task management and monitoring
- **Extensible Architecture** - Easy drop-in replacements for databases, web frameworks, and other dependencies
- **Docker Support** - Production-ready containerization with multi-stage builds

## 📋 API Endpoints

```
POST /create - Create a new task
GET  /view   - Get specific task details
POST /update - Update task status
POST /delete - Remove task from tracking
```

## 🐳 Docker Usage

### Building the Docker Image

```bash
# Build the image
docker build -t io-task-tracker .
```

### Running the Container

```bash
# Run with default configuration
docker run -p 8080:8080 io-task-tracker
```

## 🏗️ Architecture & Extensibility

The service is built with a modular architecture supporting easy extension and drop-in replacements:

### Database Layer
```go
type TaskDB interface {
    Create(name, desc string) (uint, error)
	View(taskID uint) (Task, error)
	UpdateStatus(taskID uint, status string) (Task, error)
	UpdateResult(taskID uint, result string) (Task, error)
	Delete(taskID uint) error
}
```

**Supported Implementations:**
- In-memory (default)
- PostgreSQL 
- MySQL
- MongoDB
- Redis

### Web Framework Layer
```go
type Server interface {
    http.Handler
    GET(pattern string, handler HandlerFunc)
    POST(pattern string, handler HandlerFunc)
}
```

**Supported Frameworks:**
- Standard library (default)
- Gin 
- Gorilla Mux
- and more... 

> **NOTE:** The modular design allows you to swap implementations without changing your business logic. Simply implement the required interfaces and inject your preferred dependencies.

### API Testing
For comprehensive API testing, I recommend using [ghostman](https://github.com/bigelle/ghostman)- a tool made by me (BTW).

```bash
# Example API tests
# For now there's no binary :(
# From the ghostman git directory:
go run main.go localhost:8080/create -M POST --print-out --data '{"name": "foo", "description": "literally foo"}'
```

## 🚀 Quick Start

1. **Clone the repository:**
   ```bash
   git clone https://github.com/bigelle/taskservice
   cd taskservice
   ```

2. **Run with Docker:**
   ```bash
   docker run -p 8080:8080 io-task-tracker
   ```

3. **Test the API:**
   ```bash
   # Using curl (or you can use whatever tool you prefer *blink*)
   curl http://localhost:8080/create --data '{"name": "foo", "description": "literally foo"}'    
   ```


## 📝 Development

### Prerequisites
- Go 1.24+
- Docker 

