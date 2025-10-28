package handlers

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	ws "github.com/ChukwukaRosemary23/flowboard-backend/internal/websocket"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins in development
		// In production, check origin properly
		return true
	},
}

// HandleWebSocket handles WebSocket connections
func HandleWebSocket(hub *ws.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		boardIDStr := c.Query("board_id")
		if boardIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "board_id required"})
			return
		}

		boardID, err := strconv.ParseUint(boardIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board_id"})
			return
		}

		// Get token from query parameter (for WebSocket)
		tokenString := c.Query("token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
			return
		}

		// Remove "Bearer " prefix if present
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Get JWT secret from environment
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			log.Println("WARNING: JWT_SECRET not set, using default")
			jwtSecret = "change-me-in-production"
		}

		// Parse and validate JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Make sure token method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			log.Printf("Invalid token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Extract user ID from token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user_id in token"})
			return
		}
		userID := uint(userIDFloat)

		// Verify user has access to this board
		var board models.Board
		if err := database.DB.Where("id = ? AND owner_id = ?", boardID, userID).First(&board).Error; err != nil {
			log.Printf("Board access denied: User %d -> Board %d, Error: %v", userID, boardID, err)
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied to this board"})
			return
		}

		// Upgrade HTTP connection to WebSocket
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("WebSocket upgrade error: %v", err)
			return
		}

		// Create new client with all fields
		client := ws.NewClient(hub, conn, uint(boardID), userID)

		// Register client
		hub.Register(client)

		// Start goroutines
		go client.WritePump()
		go client.ReadPump()

		log.Printf("🔌 WebSocket connected: User %d → Board %d", userID, boardID)
	}
}
