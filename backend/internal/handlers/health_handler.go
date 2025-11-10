package handlers

import (
	"net/http"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/gin-gonic/gin"
)

// HealthCheck checks if server and database are healthy
func HealthCheck(c *gin.Context) {
	// Check database connection
	sqlDB, err := database.DB.DB()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "unhealthy",
			"message": "Database connection failed",
			"error":   err.Error(),
		})
		return
	}

	// Ping database
	if err := sqlDB.Ping(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "unhealthy",
			"message": "Database ping failed",
			"error":   err.Error(),
		})
		return
	}

	// Everything is healthy
	c.JSON(http.StatusOK, gin.H{
		"status":   "healthy",
		"message":  "Server and database are operational",
		"database": "connected",
	})
}
