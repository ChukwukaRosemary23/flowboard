package routes

import (
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/handlers"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/middleware"
	ws "github.com/ChukwukaRosemary23/flowboard-backend/internal/websocket"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, hub *ws.Hub) {

	api := router.Group("/api/v1")
	{

		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		api.GET("/ws", handlers.HandleWebSocket(hub))

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

			boards := protected.Group("/boards")
			{
				boards.POST("", handlers.CreateBoard)
				boards.GET("", handlers.GetBoards)

				// Board detail routes - require board access
				boards.GET("/:id", middleware.RequireBoardAccess(), handlers.GetBoard)
				boards.PUT("/:id", middleware.RequirePermission("update_board"), handlers.UpdateBoard)
				boards.DELETE("/:id", middleware.RequireOwner(), handlers.DeleteBoard)

				// Board member management routes
				boards.GET("/:id/members", middleware.RequireBoardAccess(), handlers.GetBoardMembers)
				boards.POST("/:id/members", middleware.RequirePermission("invite_member"), handlers.InviteMember)
				boards.DELETE("/:id/members/:member_id", middleware.RequirePermission("manage_members"), handlers.RemoveMember)
				boards.PUT("/:id/members/:user_id/role", middleware.RequirePermission("manage_members"), handlers.UpdateMemberRole)
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
				cards.POST("", handlers.CreateCard)
				cards.GET("/list/:list_id", handlers.GetCards)
				cards.GET("/:id", handlers.GetCard)
				cards.PUT("/:id", handlers.UpdateCard)
				cards.PATCH("/:id", handlers.UpdateCard)
				cards.POST("/:id/move", handlers.MoveCard)
				cards.DELETE("/:id", handlers.DeleteCard)
			}

			// Comment routes
			comments := protected.Group("/comments")
			{
				comments.POST("/card/:card_id", handlers.CreateComment)
				comments.GET("/card/:card_id", handlers.GetComments)
				comments.PUT("/:id", handlers.UpdateComment)
				comments.PATCH("/:id", handlers.UpdateComment)
				comments.DELETE("/:id", handlers.DeleteComment)
			}

			// Label routes
			labels := protected.Group("/labels")
			{
				labels.POST("/board/:board_id", handlers.CreateLabel)
				labels.GET("/board/:board_id", handlers.GetLabels)
				labels.PUT("/:id", handlers.UpdateLabel)
				labels.PATCH("/:id", handlers.UpdateLabel)
				labels.DELETE("/:id", handlers.DeleteLabel)
				labels.POST("/card/:card_id", handlers.AddLabelToCard)
				labels.DELETE("/card/:card_id/:label_id", handlers.RemoveLabelFromCard)
			}

			// Card Member routes
			cardMembers := protected.Group("/card-members")
			{
				cardMembers.POST("/card/:card_id", handlers.AssignMemberToCard)
				cardMembers.GET("/card/:card_id", handlers.GetCardMembers)
				cardMembers.DELETE("/card/:card_id/member/:member_id", handlers.UnassignMemberFromCard)
			}

			// Search routes
			search := protected.Group("/search")
			{
				search.GET("/cards", handlers.SearchCards)
				search.GET("/overdue", handlers.GetOverdueCards)
				search.GET("/upcoming", handlers.GetUpcomingCards)
			}

			// Activity routes
			activities := protected.Group("/activities")
			{
				activities.GET("/board/:board_id", handlers.GetBoardActivities)
				activities.GET("/me", handlers.GetUserActivities)
			}

			// Attachment routes
			attachments := protected.Group("/attachments")
			{
				attachments.POST("/card/:card_id", handlers.UploadAttachment)
				attachments.GET("/card/:card_id", handlers.GetAttachments)
				attachments.GET("/:id/download", handlers.DownloadAttachment)
				attachments.DELETE("/:id", handlers.DeleteAttachment)
			}
		}
	}
}
