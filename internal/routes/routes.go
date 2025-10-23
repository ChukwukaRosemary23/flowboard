package routes

import (
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/handlers"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine) {
	// API v1 group
	api := router.Group("/api/v1")
	{
		// Public routes (no authentication required)
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		// Protected routes (authentication required)
		protected := api.Group("")
		protected.Use(middleware.AuthRequired())
		{
			// User routes
			protected.GET("/me", func(c *gin.Context) {
				userID := c.GetUint("user_id")
				username := c.GetString("username")
				email := c.GetString("email")

				c.JSON(200, gin.H{
					"message":  "You are authenticated!",
					"user_id":  userID,
					"username": username,
					"email":    email,
				})
			})

			// Board routes
			boards := protected.Group("/boards")
			{
				boards.POST("", handlers.CreateBoard)       // Create board
				boards.GET("", handlers.GetBoards)          // Get all user's boards
				boards.GET("/:id", handlers.GetBoard)       // Get single board
				boards.PUT("/:id", handlers.UpdateBoard)    // Update board
				boards.PATCH("/:id", handlers.UpdateBoard)  // Update board (alternative)
				boards.DELETE("/:id", handlers.DeleteBoard) // Delete board
			}

			// List routes
			lists := protected.Group("/lists")
			{
				lists.POST("", handlers.CreateList)              // Create list
				lists.GET("/board/:board_id", handlers.GetLists) // Get all lists in a board (CHANGED!)
				lists.GET("/:id", handlers.GetList)              // Get single list
				lists.PUT("/:id", handlers.UpdateList)           // Update list
				lists.PATCH("/:id", handlers.UpdateList)         // Update list
				lists.POST("/:id/move", handlers.MoveList)       // Move list (reorder)
				lists.DELETE("/:id", handlers.DeleteList)        // Delete list
			}

			// We'll add card routes here later
		}
	}
}
