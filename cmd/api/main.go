package main

import (
	"log"

	"github.com/ChukwukaRosemary23/flowboard-backend/config"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration from .env
	cfg := config.LoadConfig()

	// Connect to database
	if err := database.ConnectDatabase(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Set Gin mode based on environment
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Gin router
	router := gin.Default()

	// Setup CORS (allow frontend to connect)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"}, // React dev servers
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Health check route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  "FlowBoard API is running!",
		})
	})

	// Setup API routes
	routes.SetupRoutes(router)

	// Start server
	serverAddr := ":" + cfg.Port
	log.Printf("🚀 Server starting on http://localhost%s\n", serverAddr)
	log.Printf("📊 Environment: %s\n", cfg.Env)
	log.Println("📚 API Endpoints:")
	log.Println("   Auth:")
	log.Println("     POST   /api/v1/auth/register         - Register new user")
	log.Println("     POST   /api/v1/auth/login            - Login user")
	log.Println("   User:")
	log.Println("     GET    /api/v1/me                    - Get current user")
	log.Println("   Boards:")
	log.Println("     POST   /api/v1/boards                - Create board")
	log.Println("     GET    /api/v1/boards                - Get all boards")
	log.Println("     GET    /api/v1/boards/:id            - Get single board")
	log.Println("     PUT    /api/v1/boards/:id            - Update board")
	log.Println("     DELETE /api/v1/boards/:id            - Delete board")
	log.Println("   Lists:")
	log.Println("     POST   /api/v1/lists                 - Create list")
	log.Println("     GET    /api/v1/boards/:id/lists      - Get all lists in board")
	log.Println("     GET    /api/v1/lists/:id             - Get single list")
	log.Println("     PUT    /api/v1/lists/:id             - Update list")
	log.Println("     POST   /api/v1/lists/:id/move        - Move/reorder list")
	log.Println("     DELETE /api/v1/lists/:id             - Delete list")

	log.Println("   Cards:")
	log.Println("     POST   /api/v1/cards                 - Create card")
	log.Println("     GET    /api/v1/cards/list/:id        - Get all cards in list")
	log.Println("     GET    /api/v1/cards/:id             - Get single card")
	log.Println("     PUT    /api/v1/cards/:id             - Update card")
	log.Println("     POST   /api/v1/cards/:id/move        - Move card (drag-and-drop)")
	log.Println("     DELETE /api/v1/cards/:id             - Delete card")

	if err := router.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
