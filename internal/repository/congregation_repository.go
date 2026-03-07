package repository

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/user/gocrud-api/internal/models"
)

// CongregationRepository is the data access layer for congregations
// In a real application, this would interface with a database
// Here we use an in-memory slice with thread-safe operations using mutex
type CongregationRepository struct {
	congregations []models.Congregation
	mu            sync.RWMutex
	nextID        int
}

// CongregationHandler handles HTTP requests for congregations
type CongregationHandler struct {
	repo *CongregationRepository
}

// NewCongregationRepository creates a new congregation repository with dummy data
func NewCongregationRepository() *CongregationRepository {
	return &CongregationRepository{
		congregations: []models.Congregation{
			{ID: "1", Name: "First Congregation", Location: "City A"},
			{ID: "2", Name: "Second Congregation", Location: "City B"},
		},
		nextID: 3,
	}
}

// NewCongregationHandler creates a new congregation handler
func NewCongregationHandler(repo *CongregationRepository) *CongregationHandler {
	return &CongregationHandler{repo: repo}
}

// GetAll returns all congregations
func (r *CongregationRepository) GetAll() []models.Congregation {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.congregations
}

// GetByID retrieves a congregation by ID
func (r *CongregationRepository) GetByID(id string) (*models.Congregation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, congregation := range r.congregations {
		if congregation.ID == id {
			return &congregation, nil
		}
	}
	return nil, fmt.Errorf("congregation not found")
}

// Create adds a new congregation
func (r *CongregationRepository) Create(req models.CreateCongregationRequest) (*models.Congregation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	congregation := models.Congregation{
		ID:       fmt.Sprintf("%d", r.nextID),
		Name:     req.Name,
		Location: req.Location,
	}
	r.nextID++
	r.congregations = append(r.congregations, congregation)

	return &congregation, nil
}

// Update modifies an existing congregation
func (r *CongregationRepository) Update(id string, req models.UpdateCongregationRequest) (*models.Congregation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, congregation := range r.congregations {
		if congregation.ID == id {
			if req.Name != "" {
				r.congregations[i].Name = req.Name
			}
			if req.Location != "" {
				r.congregations[i].Location = req.Location
			}
			return &r.congregations[i], nil
		}
	}
	return nil, fmt.Errorf("congregation not found")
}

// Delete removes a congregation by ID
func (r *CongregationRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, congregation := range r.congregations {
		if congregation.ID == id {
			r.congregations = append(r.congregations[:i], r.congregations[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("congregation not found")
}

// CreateCongregation handles POST /congregations
// Creates a new congregation
// @Summary Create a new congregation
// @Description Add a new congregation to the system
// @Tags congregations
// @Accept json
// @Produce json
// @Param congregation body models.CreateCongregationRequest true "Congregation data"
// @Success 201 {object} models.CongregationResponse
// @Failure 400 {object} models.CongregationErrorResponse
// @Router /congregations [post]
func (h *CongregationHandler) CreateCongregation(c *gin.Context) {
	var req models.CreateCongregationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.CongregationErrorResponse{
			Error: "Invalid request data",
			Code:  400,
		})
		return
	}

	congregation, err := h.repo.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to create congregation",
			Code:  500,
		})
		return
	}

	c.JSON(http.StatusCreated, models.CongregationResponse{
		Message: "Congregation created successfully",
		Data:    congregation,
	})
}

// UpdateCongregation handles PUT /congregations/:id
// Updates an existing congregation
// @Summary Update a congregation
// @Description Modify the details of an existing congregation
// @Tags congregations
// @Accept json
// @Produce json
// @Param id path string true "Congregation ID"
// @Param congregation body models.UpdateCongregationRequest true "Updated congregation data"
// @Success 200 {object} models.CongregationResponse
// @Failure 400 {object} models.CongregationErrorResponse
// @Failure 404 {object} models.CongregationErrorResponse
// @Router /congregations/{id} [put]
func (h *CongregationHandler) UpdateCongregation(c *gin.Context) {
	id := c.Param("id")
	var req models.UpdateCongregationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.CongregationErrorResponse{
			Error: "Invalid request data",
			Code:  400,
		})
		return
	}

	congregation, err := h.repo.Update(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.CongregationErrorResponse{
			Error: "Congregation not found",
			Code:  404,
		})
		return
	}

	c.JSON(http.StatusOK, models.CongregationResponse{
		Message: "Congregation updated successfully",
		Data:    congregation,
	})
}

// DeleteCongregation handles DELETE /congregations/:id
// Deletes a congregation by ID
// @Summary Delete a congregation
// @Description Remove a congregation from the system by its ID
// @Tags congregations
// @Accept json
// @Produce json
// @Param id path string true "Congregation ID"
// @Success 204 "No Content"
// @Failure 404 {object} models.CongregationErrorResponse
// @Router /congregations/{id} [delete]
func (h *CongregationHandler) DeleteCongregation(c *gin.Context) {
	id := c.Param("id")

	if err := h.repo.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, models.CongregationErrorResponse{
			Error: "Congregation not found",
			Code:  404,
		})
		return
	}

	c.Status(http.StatusNoContent)
}
