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
			}

			// List routes
			lists := protected.Group("/lists")
			{
				lists.POST("", handlers.CreateList)
				lists.GET("/board/:board_id", handlers.GetLists)
				lists.GET("/:id", handlers.GetList)
				lists.PUT("/:id", handlers.UpdateList)
				lists.PATCH("/:id", handlers.UpdateList)
				lists.POST("/:id/move", handlers.MoveList)
				lists.DELETE("/:id", handlers.DeleteList)
			}

			// Card routes
			cards := protected.Group("/cards")
			{
				cards.POST("", handlers.CreateCard)            // Create card
				cards.GET("/list/:list_id", handlers.GetCards) // Get all cards in list
				cards.GET("/:id", handlers.GetCard)            // Get single card
				cards.PUT("/:id", handlers.UpdateCard)         // Update card
				cards.PATCH("/:id", handlers.UpdateCard)       // Update card
				cards.POST("/:id/move", handlers.MoveCard)     // Move card (drag-and-drop!)
				cards.DELETE("/:id", handlers.DeleteCard)      // Delete card
			}

			// We'll add comment, label, attachment routes later
		}
	}
}
