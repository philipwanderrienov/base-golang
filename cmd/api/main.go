package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// this import is required to register the generated swagger docs
	_ "github.com/user/gocrud-api/docs"

	"github.com/user/gocrud-api/internal/handlers"
	"github.com/user/gocrud-api/internal/repository"
	"github.com/user/gocrud-api/internal/routes"
)

// @title Inventory Management API
// @version 1.0
// @description This is a sample CRUD API for inventory management system
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @tag.name items
// @tag.description Operations for managing inventory items
func main() {
	// Initialize the repository with dummy data
	// This acts as our "database" layer in the microservices architecture
	repo := repository.NewItemRepository()

	// Initialize the handler (service layer)
	// This handles business logic and request processing
	handler := handlers.NewItemHandler(repo)

	// Create a new Gin router
	// Gin is a web framework that provides HTTP request handling
	r := gin.Default()

	// Create API v1 group
	// This allows for versioning of our API
	v1 := r.Group("/api/v1")
	{
		// register resource-specific route groups
		routes.RegisterItemRoutes(v1, handler)
	}

	// Swagger documentation route
	// Access at /swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server
	// Server runs on port 8080
	log.Println("Server starting on http://localhost:8080")
	log.Println("Swagger docs available at http://localhost:8080/swagger/index.html")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
