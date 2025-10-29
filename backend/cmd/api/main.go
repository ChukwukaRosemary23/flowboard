package main

import (
	"log"
	"os"
	"time"

	"github.com/ChukwukaRosemary23/flowboard-backend/config"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/handlers"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/routes"
	ws "github.com/ChukwukaRosemary23/flowboard-backend/internal/websocket"
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

	// Create WebSocket hub
	hub := ws.NewHub()
	go hub.Run()
	log.Println("🔌 WebSocket hub started")

	// Set global hub for handlers
	handlers.WSHub = hub

	// Set Gin mode based on environment
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Gin router
	router := gin.Default()

	// Serve uploaded files
	router.Static("/uploads", "./uploads")

	// Setup CORS (allow frontend to connect)
	allowedOrigins := []string{
		"http://localhost:3000",
		"http://localhost:3001",
		"http://localhost:5173",
		"http://127.0.0.1:3000",
		"http://127.0.0.1:3001",
		"http://127.0.0.1:5173",
	}

	// Add production frontend URL from environment variable
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL != "" {
		allowedOrigins = append(allowedOrigins, frontendURL)
		log.Printf("✅ Added frontend URL to CORS: %s", frontendURL)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  "FlowBoard API is running!",
		})
	})

	// Setup API routes (pass hub)
	routes.SetupRoutes(router, hub)

	// Get port from environment variable (Render uses this)
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port // Fallback to config (local development)
	}

	serverAddr := ":" + port
	log.Printf("🚀 Server starting on port %s\n", port)
	log.Printf("📊 Environment: %s\n", cfg.Env)
	log.Println("📚 API Endpoints:")
	log.Println("   Health Check:")
	log.Println("     GET    /ping                         - Health check")
	log.Println("   WebSocket:")
	log.Println("     GET    /api/v1/ws?board_id=X         - WebSocket connection")
	log.Println("   Auth:")
	log.Println("     POST   /api/v1/auth/register         - Register new user")
	log.Println("     POST   /api/v1/auth/login            - Login user")
	log.Println("   Boards, Lists, Cards, Comments, Labels, etc...")

	if err := router.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
