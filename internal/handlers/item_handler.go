package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/gocrud-api/internal/models"
	"github.com/user/gocrud-api/internal/repository"
)

// ItemHandler handles HTTP requests for item operations
// This is the service layer in our microservices architecture
// It acts as an intermediary between HTTP requests and the data repository
type ItemHandler struct {
	repo *repository.ItemRepository
}

// NewItemHandler creates a new item handler with the given repository
// This follows the dependency injection pattern for better testability
func NewItemHandler(repo *repository.ItemRepository) *ItemHandler {
	return &ItemHandler{repo: repo}
}

// GetAllItems handles GET /items
// Returns a list of all items in the system
// @Summary Get all items
// @Description Retrieve all items from the inventory
// @Tags items
// @Accept json
// @Produce json
// @Success 200 {object} models.ItemsListResponse
// @Router /items [get]
func (h *ItemHandler) GetAllItems(c *gin.Context) {
	items := h.repo.GetAll()
	c.JSON(http.StatusOK, models.ItemsListResponse{
		Message: "Items retrieved successfully",
		Data:    items,
		Total:   len(items),
	})
}

// GetItemByID handles GET /items/:id
// Retrieves a single item by its ID
// @Summary Get item by ID
// @Description Retrieve a specific item by its unique identifier
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} models.ItemResponse
// @Router /items/{id} [get]
func (h *ItemHandler) GetItemByID(c *gin.Context) {
	// Parse item ID from URL path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid item ID",
			Code:  400,
		})
		return
	}

	// Retrieve item from repository
	item := h.repo.GetByID(id)
	if item == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Item not found",
			Code:  404,
		})
		return
	}

	c.JSON(http.StatusOK, models.ItemResponse{
		Message: "Item retrieved successfully",
		Data:    item,
	})
}

// CreateItem handles POST /items
// Creates a new item in the inventory
// @Summary Create new item
// @Description Add a new item to the inventory
// @Tags items
// @Accept json
// @Produce json
// @Param item body models.ItemRequest true "Item data"
// @Success 201 {object} models.ItemResponse
// @Router /items [post]
func (h *ItemHandler) CreateItem(c *gin.Context) {
	// Bind incoming JSON request to ItemRequest struct
	var req models.ItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body: " + err.Error(),
			Code:  400,
		})
		return
	}

	// Convert request to Item model
	item := &models.Item{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Category:    req.Category,
	}

	// Create item in repository
	created := h.repo.Create(item)

	c.JSON(http.StatusCreated, models.ItemResponse{
		Message: "Item created successfully",
		Data:    created,
	})
}

// UpdateItem handles PUT /items/:id
// Updates an existing item by its ID
// @Summary Update item
// @Description Update an existing item's information
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param item body models.ItemRequest true "Updated item data"
// @Success 200 {object} models.ItemResponse
// @Router /items/{id} [put]
func (h *ItemHandler) UpdateItem(c *gin.Context) {
	// Parse item ID from URL path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid item ID",
			Code:  400,
		})
		return
	}

	// Bind incoming JSON request to ItemRequest struct
	var req models.ItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body: " + err.Error(),
			Code:  400,
		})
		return
	}

	// Convert request to Item model
	item := &models.Item{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Category:    req.Category,
	}

	// Update item in repository
	updated := h.repo.Update(id, item)
	if updated == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Item not found",
			Code:  404,
		})
		return
	}

	c.JSON(http.StatusOK, models.ItemResponse{
		Message: "Item updated successfully",
		Data:    updated,
	})
}

// DeleteItem handles DELETE /items/:id
// Removes an item from the inventory by its ID
// @Summary Delete item
// @Description Remove an item from the inventory
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} models.ItemResponse
// @Router /items/{id} [delete]
func (h *ItemHandler) DeleteItem(c *gin.Context) {
	// Parse item ID from URL path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid item ID",
			Code:  400,
		})
		return
	}

	// Delete item from repository
	deleted := h.repo.Delete(id)
	if !deleted {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Item not found",
			Code:  404,
		})
		return
	}

	c.JSON(http.StatusOK, models.ItemResponse{
		Message: "Item deleted successfully",
		Data:    nil,
	})
}

// SearchItems handles GET /items/search
// Searches for items by name (case-insensitive)
// @Summary Search items
// @Description Search items by name (case-insensitive)
// @Tags items
// @Accept json
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} models.ItemsListResponse
// @Router /items/search [get]
func (h *ItemHandler) SearchItems(c *gin.Context) {
	// Get query parameter from URL
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Search query is required",
			Code:  400,
		})
		return
	}

	// Search items in repository
	items := h.repo.Search(query)

	c.JSON(http.StatusOK, models.ItemsListResponse{
		Message: "Search completed successfully",
		Data:    items,
		Total:   len(items),
	})
}

// GetItemsByCategory handles GET /items/category/:category
// Retrieves all items belonging to a specific category
// @Summary Get items by category
// @Description Retrieve all items in a specific category
// @Tags items
// @Accept json
// @Produce json
// @Param category path string true "Category name"
// @Success 200 {object} models.ItemsListResponse
// @Router /items/category/{category} [get]
func (h *ItemHandler) GetItemsByCategory(c *gin.Context) {
	// Get category from URL path parameter
	category := c.Param("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Category is required",
			Code:  400,
		})
		return
	}

	// Get items by category from repository
	items := h.repo.GetByCategory(category)

	c.JSON(http.StatusOK, models.ItemsListResponse{
		Message: "Items retrieved successfully",
		Data:    items,
		Total:   len(items),
	})
}
