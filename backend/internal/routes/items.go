package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/user/gocrud-api/internal/handlers"
)

// RegisterItemRoutes wires up the item-related endpoints under the
// provided versioned group.  Keeping route definitions in their own
// package makes it easy to add additional resource groups (auth, users,
// etc.) later without cluttering the main entry point.
func RegisterItemRoutes(v1 *gin.RouterGroup, h *handlers.ItemHandler) {
	items := v1.Group("/items")
	{
		// CRUD
		items.GET("", h.GetAllItems)
		items.GET("/:id", h.GetItemByID)
		items.POST("", h.CreateItem)
		items.PUT("/:id", h.UpdateItem)
		items.DELETE("/:id", h.DeleteItem)

		// extra operations
		items.GET("/search", h.SearchItems)
		items.GET("/category/:category", h.GetItemsByCategory)
	}
}
