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
		// Public routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		// Protected routes
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
				boards.POST("", handlers.CreateBoard)
				boards.GET("", handlers.GetBoards)
				boards.GET("/:id", handlers.GetBoard)
				boards.PUT("/:id", handlers.UpdateBoard)
				boards.PATCH("/:id", handlers.UpdateBoard)
				boards.DELETE("/:id", handlers.DeleteBoard)

				// List routes (nested under boards)
				boards.GET("/:board_id/lists", handlers.GetLists) // Get all lists in a board
			}

			// List routes (direct access)
			lists := protected.Group("/lists")
			{
				lists.POST("", handlers.CreateList)        // Create list
				lists.GET("/:id", handlers.GetList)        // Get single list
				lists.PUT("/:id", handlers.UpdateList)     // Update list
				lists.PATCH("/:id", handlers.UpdateList)   // Update list
				lists.POST("/:id/move", handlers.MoveList) // Move list (reorder)
				lists.DELETE("/:id", handlers.DeleteList)  // Delete list
			}

			// We'll add card routes here later
		}
	}
}
