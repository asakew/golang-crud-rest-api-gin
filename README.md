# Golang CRUD REST API with Gin

This is a simple CRUD (Create, Read, Update, Delete) REST API in Go using the Gin web framework.

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine.

### Prerequisites

- Go (at least version 1.11)

### Cloning the Project

```bash
git clone https://github.com/Shikha-code36/golang-crud-rest-api-gin.git
cd golang-crud-rest-api-gin
```

### Initializing the Project

```bash
go mod init github.com/Shikha-code36/golang-crud-rest-api-gin
```

### Downloading Dependencies

```bash
go mod tidy
```

## Running the Application

```bash
go run main.go
```

The application will start running at `http://localhost:8080`.

## Explanation of Code

### main.go

The `main.go` file is the entry point of our application. Here's a breakdown of its structure and functionality:

```go
package main
```
This line defines the package that this file belongs to. In Go, the `main` package is special. It's the entry point of the application, and it must contain a `main` function.

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/Shikha-code36/golang-crud-rest-api-gin/handlers"
)
```
This block imports two packages. The first one, `github.com/gin-gonic/gin`, is the Gin web framework that we're using to build our API. The second one, `github.com/Shikha-code36/golang-crud-rest-api-gin/handlers`, is the package that contains our handler functions.

```go
func main() {
    r := gin.Default()
```
This is the `main` function where our application starts. `gin.Default()` creates a new Gin engine with some default middleware (logger and recovery). This engine is used to set up our routes and start the server.

```go
    r.GET("/users", handlers.GetUsers)
    r.GET("/user/:id", handlers.GetUser)
    r.POST("/create_user", handlers.CreateUser)
    r.PUT("/update_user/:id", handlers.UpdateUser)
    r.DELETE("/delete_user/:id", handlers.DeleteUser)
```
These lines define our API routes. For each route, we specify an HTTP method (GET, POST, PUT, DELETE), a path, and a handler function. When a request is received that matches a route's method and path, the corresponding handler function is called.

```go
    r.Run()
}
```
Finally, `r.Run()` starts the Gin server and begins listening for requests. By default, it listens on `localhost:8080`, but you can pass a different address as an argument to `Run` if you want.

That's the basic structure of our `main.go` file. It sets up our server, defines our routes, and starts the server.

### users.go in handlers

The `users.go` file in the `handlers` directory contains the handler functions for our routes. You can find the detailed explanation [here](handlers).

## API Endpoints

- **GET /users**: Fetch all users.
- **GET /user/:id**: Fetch a user by ID.
- **POST /create_user**: Create a new user. The request body should be a JSON object with a `name` field, like this: `{"name": "Alice"}`. The ID will be assigned automatically.
- **PUT /update_user/:id**: Update a user by ID. The request body should be a JSON object with a `name` field, like this: `{"name": "Bob"}`.
- **DELETE /delete_user/:id**: Delete a user by ID.

## Testing the API

You can use a tool like curl or Postman to test the API. Here are some examples:

- Fetch all users:

```bash
curl -X GET http://localhost:8080/users
```

- Fetch a user by ID:

```bash
curl -X GET http://localhost:8080/user/1
```

- Create a new user:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"Alice"}' http://localhost:8080/create_user
```

- Update a user by ID:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Bob"}' http://localhost:8080/update_user/1
```

- Delete a user by ID:

```bash
curl -X DELETE http://localhost:8080/delete_user/1
```
