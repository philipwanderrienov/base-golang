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
// @tag.name accounts
// @tag.description Operations for managing accounts
func main() {
	// Initialize repositories with dummy data
	// These acts as our "database" layer in the microservices architecture
	itemRepo := repository.NewItemRepository()
	accountRepo := repository.NewAccountRepository()
	congregationRepo := repository.NewCongregationRepository()

	// Initialize handlers (service layer)
	// These handle business logic and request processing
	itemHandler := handlers.NewItemHandler(itemRepo)
	accountHandler := handlers.NewAccountHandler(accountRepo)
	congregationHandler := handlers.NewCongregationHandler(congregationRepo)

	// Create a new Gin router
	// Gin is a web framework that provides HTTP request handling
	r := gin.Default()

	// Create API v1 group
	// This allows for versioning of our API
	v1 := r.Group("/api/v1")
	{
		// register resource-specific route groups
		routes.RegisterItemRoutes(v1, itemHandler)
		routes.RegisterAccountRoutes(v1, accountHandler)
		routes.RegisterCongregationRoutes(v1, congregationHandler)
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
