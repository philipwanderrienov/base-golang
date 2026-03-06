package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/gocrud-api/internal/repository"
	"github.com/user/gocrud-api/internal/models"
)

type AccountHandler struct {
	repo *repository.AccountRepository
}

func NewAccountHandler(repo *repository.AccountRepository) *AccountHandler {
	return &AccountHandler{repo: repo}
}

// GetAllAccounts handles GET /accounts
// Returns a list of all accounts in the system
// @Summary Get all accounts
// @Description Retrieve all accounts from the system
// @Tags accounts
// @Accept json
// @Produce json
// @Success 200 {object} models.AccountsListResponse
// @Router /accounts [get]
func (h *AccountHandler) GetAllAccounts(c *gin.Context) {
	accounts := h.repo.GetAll()
	c.JSON(http.StatusOK, models.AccountsListResponse{
		Message: "Accounts retrieved successfully",
		Data:    accounts,
		Total:   len(accounts),
	})
}

// GetAccountByID handles GET /accounts/:id
// Retrieves a single account by its ID
// @Summary Get account by ID
// @Description Retrieve a specific account by its unique identifier
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} models.AccountResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /accounts/{id} [get]
func (h *AccountHandler) GetAccountByID(c *gin.Context) {
	id := c.Param("id")
	account, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Account not found",
			Code:  404,
		})
		return
	}

	c.JSON(http.StatusOK, models.AccountResponse{
		Message: "Account retrieved successfully",
		Data:    account,
	})
}

// CreateAccount handles POST /accounts
// Creates a new account
// @Summary Create a new account
// @Description Add a new account to the system
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body models.CreateAccountRequest true "Account data"
// @Success 201 {object} models.AccountResponse
// @Failure 400 {object} models.AccountErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /accounts [post]
func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req models.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.AccountErrorResponse{
			Error: "Invalid request payload",
			Code:  400,
		})
		return
	}

	account, err := h.repo.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to create account",
			Code:  500,
		})
		return
	}

	c.JSON(http.StatusCreated, models.AccountResponse{
		Message: "Account created successfully",
		Data:    account,
	})
}

// UpdateAccount handles PUT /accounts/:id
// Updates an existing account
// @Summary Update an account
// @Description Update an existing account by its ID
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param account body models.UpdateAccountRequest true "Updated account data"
// @Success 200 {object} models.AccountResponse
// @Failure 400 {object} models.AccountErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /accounts/{id} [put]
func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	var req models.UpdateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.AccountErrorResponse{
			Error: "Invalid request payload",
			Code:  400,
		})
		return
	}

	account, err := h.repo.Update(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Account not found",
			Code:  404,
		})
		return
	}

	c.JSON(http.StatusOK, models.AccountResponse{
		Message: "Account updated successfully",
		Data:    account,
	})
}

// DeleteAccount handles DELETE /accounts/:id
// Deletes an account by its ID
// @Summary Delete an account
// @Description Delete an account by its ID
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} models.AccountResponse
// @Failure 404 {object} models.AccountErrorResponse
// @Router /accounts/{id} [delete]
func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	err := h.repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.AccountErrorResponse{
			Error: "Account not found",
			Code:  404,
		})
		return
	}

	c.JSON(http.StatusOK, models.AccountResponse{
		Message: "Account deleted successfully",
	})
}
