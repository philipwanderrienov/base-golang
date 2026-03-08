package models

// Item represents an item in our inventory system
// This is the core data model for our CRUD operations
type Item struct {
	ID          int    `json:"id" example:"1"`                    // Unique identifier for the item
	Name        string `json:"name" example:"Laptop"`             // Name of the item
	Description string `json:"description" example:"MacBook Pro"` // Detailed description of the item
	Price       int    `json:"price" example:"1500"`              // Price of the item in dollars
	Quantity    int    `json:"quantity" example:"10"`             // Available quantity in stock
	Category    string `json:"category" example:"Electronics"`    // Category the item belongs to
}

// ItemRequest is used for creating or updating an item
// This is the request body structure for POST and PUT operations
type ItemRequest struct {
	Name        string `json:"name" binding:"required" example:"Laptop"`
	Description string `json:"description" example:"MacBook Pro"`
	Price       int    `json:"price" binding:"required" example:"1500"`
	Quantity    int    `json:"quantity" binding:"required" example:"10"`
	Category    string `json:"category" binding:"required" example:"Electronics"`
}

// ItemResponse is the standard response structure for item operations
// Wraps the item data with a message for better API communication
type ItemResponse struct {
	Message string `json:"message" example:"Item retrieved successfully"`
	Data    *Item  `json:"data"` // Pointer to item, can be nil for empty responses
}

// ItemsListResponse is the response structure for listing multiple items
type ItemsListResponse struct {
	Message string `json:"message" example:"Items retrieved successfully"`
	Data    []Item `json:"data"`              // Slice of items
	Total   int    `json:"total" example:"5"` // Total count of items
}

// ErrorResponse is the standard error response structure
type ErrorResponse struct {
	Error string `json:"error" example:"Item not found"`
	Code  int    `json:"code" example:"404"` // HTTP status code
}
