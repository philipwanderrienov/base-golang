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
	congregations, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			// Error: "Failed to retrieve congregations",
			Error: err.Error(), // Include the actual error message for better debugging
			Code:  500,
		})
		return
	}
	c.JSON(http.StatusOK, models.CongregationsListResponse{
		Message: "Congregations retrieved successfully",
		Data:    congregations,
		Total:   len(congregations),
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
	congregation, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to retrieve congregation",
			Code:  500,
		})
		return
	}
	if congregation == nil {
		c.JSON(http.StatusNotFound, models.CongregationErrorResponse{
			Error: "Congregation not found",
			Code:  404,
		})
		return
	}
	c.JSON(http.StatusOK, models.CongregationResponse{
		Message: "Congregation retrieved successfully",
		Data:    congregation,
	})
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
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to update congregation",
			Code:  500,
		})
		return
	}

	c.JSON(http.StatusOK, models.CongregationResponse{
		Message: "Congregation updated successfully",
		Data:    congregation,
	})
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
	err := h.repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to delete congregation",
			Code:  500,
		})
		return
	}
	c.Status(http.StatusNoContent)
}
