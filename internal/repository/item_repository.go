package repository

import (
	"sync"

	"github.com/user/gocrud-api/internal/models"
)

// ItemRepository is the data access layer for items
// In a real application, this would interface with a database
// Here we use an in-memory slice with thread-safe operations using mutex
type ItemRepository struct {
	// items is the in-memory storage for our items
	items []models.Item
	// mu provides thread-safe access to the items slice
	mu sync.RWMutex
	// nextID tracks the next available ID for new items
	nextID int
}

// NewItemRepository creates a new item repository with dummy data
// This initializes our "database" with sample data
func NewItemRepository() *ItemRepository {
	// Initialize with some dummy data for demonstration
	repo := &ItemRepository{
		items: []models.Item{
			{
				ID:          1,
				Name:        "Laptop",
				Description: "High-performance laptop for professionals",
				Price:       1500,
				Quantity:    10,
				Category:    "Electronics",
			},
			{
				ID:          2,
				Name:        "Wireless Mouse",
				Description: "Ergonomic wireless mouse with precision tracking",
				Price:       50,
				Quantity:    100,
				Category:    "Electronics",
			},
			{
				ID:          3,
				Name:        "Office Chair",
				Description: "Comfortable ergonomic office chair",
				Price:       300,
				Quantity:    25,
				Category:    "Furniture",
			},
			{
				ID:          4,
				Name:        "Desk Lamp",
				Description: "LED desk lamp with adjustable brightness",
				Price:       45,
				Quantity:    50,
				Category:    "Furniture",
			},
			{
				ID:          5,
				Name:        "Notebook",
				Description: "Premium quality ruled notebook",
				Price:       10,
				Quantity:    200,
				Category:    "Stationery",
			},
		},
		nextID: 6, // Next new item will get ID 6
	}
	return repo
}

// GetAll retrieves all items from the repository
// Returns a slice of all items
func (r *ItemRepository) GetAll() []models.Item {
	r.mu.RLock() // Read lock for concurrent reads
	defer r.mu.RUnlock()
	return r.items
}

// GetByID retrieves a single item by its ID
// Returns nil if item not found
func (r *ItemRepository) GetByID(id int) *models.Item {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, item := range r.items {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

// Create adds a new item to the repository
// Returns the created item with assigned ID
func (r *ItemRepository) Create(item *models.Item) *models.Item {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Assign new ID and add to slice
	item.ID = r.nextID
	r.nextID++
	r.items = append(r.items, *item)
	return item
}

// Update modifies an existing item by its ID
// Returns the updated item or nil if not found
func (r *ItemRepository) Update(id int, item *models.Item) *models.Item {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, existing := range r.items {
		if existing.ID == id {
			// Preserve the original ID
			item.ID = id
			r.items[i] = *item
			return item
		}
	}
	return nil
}

// Delete removes an item by its ID
// Returns true if item was deleted, false if not found
func (r *ItemRepository) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, item := range r.items {
		if item.ID == id {
			// Remove item by slicing
			r.items = append(r.items[:i], r.items[i+1:]...)
			return true
		}
	}
	return false
}

// GetByCategory retrieves all items belonging to a specific category
// Returns a slice of items in that category
func (r *ItemRepository) GetByCategory(category string) []models.Item {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []models.Item
	for _, item := range r.items {
		if item.Category == category {
			result = append(result, item)
		}
	}
	return result
}

// Search performs a case-insensitive search on item names
// Returns all items that contain the query string in their name
func (r *ItemRepository) Search(query string) []models.Item {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []models.Item
	for _, item := range r.items {
		// Simple case-insensitive substring search
		if containsIgnoreCase(item.Name, query) {
			result = append(result, item)
		}
	}
	return result
}

// containsIgnoreCase is a helper function for case-insensitive string matching
func containsIgnoreCase(s, substr string) bool {
	s = toLower(s)
	substr = toLower(substr)
	return len(s) >= len(substr) && (len(s) == len(substr) || containsSubstring(s, substr))
}

func toLower(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		result[i] = c
	}
	return string(result)
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
