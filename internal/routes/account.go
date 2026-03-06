package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/user/gocrud-api/internal/handlers"
)

func RegisterAccountRoutes(v1 *gin.RouterGroup, h *handlers.AccountHandler) {
	accounts := v1.Group("/accounts")
	{
		// CRUD
		accounts.GET("", h.GetAllAccounts)
		accounts.GET("/:id", h.GetAccountByID)
		accounts.POST("", h.CreateAccount)
		accounts.PUT("/:id", h.UpdateAccount)
		accounts.DELETE("/:id", h.DeleteAccount)
	}
}
