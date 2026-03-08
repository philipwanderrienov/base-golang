package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/gocrud-api/internal/models"
	"github.com/user/gocrud-api/internal/repository"
)

type CongregationHandler struct {
	repo *repository.CongregationRepository
}

func NewCongregationHandler(repo *repository.CongregationRepository) *CongregationHandler {
	return &CongregationHandler{repo: repo}
}

// GetAllCongregations handles GET /congregations
// Returns a list of all congregations in the system
// @Summary Get all congregations
// @Description Retrieve all congregations from the system
// @Tags congregations
// @Accept json
// @Produce json
// @Success 200 {object} models.CongregationsListResponse
// @Router /congregations [get]
func (h *CongregationHandler) GetAllCongregations(c *gin.Context) {
	// This is a placeholder implementation. In a real application, you would retrieve this data from a repository.
	c.JSON(http.StatusOK, models.CongregationsListResponse{
		Message: "Congregations retrieved successfully",
		Data: []models.Congregation{
			{ID: "1", Name: "First Congregation", Location: "City A"},
			{ID: "2", Name: "Second Congregation", Location: "City B"},
		},
		Total: 2,
	})
}

// GetCongregationByID handles GET /congregations/:id
// Retrieves a single congregation by its ID
// @Summary Get congregation by ID
// @Description Retrieve a specific congregation by its unique identifier
// @Tags congregations
// @Accept json
// @Produce json
// @Param id path string true "Congregation ID"
// @Success 200 {object} models.CongregationResponse
// @Failure 404 {object} models.CongregationErrorResponse
// @Router /congregations/{id} [get]
func (h *CongregationHandler) GetCongregationByID(c *gin.Context) {
	id := c.Param("id")
	// This is a placeholder implementation. In a real application, you would retrieve this data from a repository.
	if id == "1" {
		c.JSON(http.StatusOK, models.CongregationResponse{
			Message: "Congregation retrieved successfully",
			Data:    &models.Congregation{ID: "1", Name: "First Congregation", Location: "City A"},
		})
	} else if id == "2" {
		c.JSON(http.StatusOK, models.CongregationResponse{
			Message: "Congregation retrieved successfully",
			Data:    &models.Congregation{ID: "2", Name: "Second Congregation", Location: "City B"},
		})
	} else {
		c.JSON(http.StatusNotFound, models.CongregationErrorResponse{
			Error: "Congregation not found",
			Code:  404,
		})
	}
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

	// This is a placeholder implementation. In a real application, you would save this data to a repository and generate a unique ID.
	newCongregation := models.Congregation{
		ID:       "3", // In a real application, this would be generated
		Name:     req.Name,
		Location: req.Location,
	}

	c.JSON(http.StatusCreated, models.CongregationResponse{
		Message: "Congregation created successfully",
		Data:    &newCongregation,
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

	// This is a placeholder implementation. In a real application, you would update this data in a repository.
	if id == "1" {
		updatedCongregation := models.Congregation{
			ID:       "1",
			Name:     req.Name,
			Location: req.Location,
		}
		c.JSON(http.StatusOK, models.CongregationResponse{
			Message: "Congregation updated successfully",
			Data:    &updatedCongregation,
		})
	} else if id == "2" {
		updatedCongregation := models.Congregation{
			ID:       "2",
			Name:     req.Name,
			Location: req.Location,
		}
		c.JSON(http.StatusOK, models.CongregationResponse{
			Message: "Congregation updated successfully",
			Data:    &updatedCongregation,
		})
	} else {
		c.JSON(http.StatusNotFound, models.CongregationErrorResponse{
			Error: "Congregation not found",
			Code:  404,
		})
	}
}

// DeleteCongregation handles DELETE /congregations/:id
// Deletes a congregation by its ID
// @Summary Delete a congregation
// @Description Remove a congregation from the system by its unique identifier
// @Tags congregations
// @Accept json
// @Produce json
// @Param id path string true "Congregation ID"
// @Success 204 "No Content"
// @Failure 404 {object} models.CongregationErrorResponse
// @Router /congregations/{id} [delete]
func (h *CongregationHandler) DeleteCongregation(c *gin.Context) {
	id := c.Param("id")
	// This is a placeholder implementation. In a real application, you would delete this data from a repository.
	if id == "1" || id == "2" {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusNotFound, models.CongregationErrorResponse{
			Error: "Congregation not found",
			Code:  404,
		})
	}
}
