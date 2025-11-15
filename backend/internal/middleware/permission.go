package middleware

import (
	"net/http"
	"strconv"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/services"
	"github.com/gin-gonic/gin"
)

// RequirePermission middleware checks if user has specific permission on a board
func RequirePermission(permissionName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

	
		boardIDStr := c.Param("board_id")
		if boardIDStr == "" {
			boardIDStr = c.Param("id") 
		}

		boardID, err := strconv.ParseUint(boardIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
			c.Abort()
			return
		}

		// Check permission
		permService := &services.PermissionService{}
		hasPermission := permService.CheckPermission(userID.(uint), uint(boardID), permissionName)

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "You do not have permission to perform this action",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireOwner ensures only board owner can access
func RequireOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		boardIDStr := c.Param("id")
		if boardIDStr == "" {
			boardIDStr = c.Param("board_id")
		}

		boardID, err := strconv.ParseUint(boardIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
			c.Abort()
			return
		}

		permService := &services.PermissionService{}
		if !permService.IsOwner(userID.(uint), uint(boardID)) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "You do not have permission to perform this action",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireAdmin ensures user is admin or owner
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		boardIDStr := c.Param("id")
		if boardIDStr == "" {
			boardIDStr = c.Param("board_id")
		}

		boardID, err := strconv.ParseUint(boardIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
			c.Abort()
			return
		}

		permService := &services.PermissionService{}
		if !permService.IsAdmin(userID.(uint), uint(boardID)) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "You do not have permission to perform this action",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireBoardAccess checks if user has any access to the board
func RequireBoardAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		boardIDStr := c.Param("id")
		if boardIDStr == "" {
			boardIDStr = c.Param("board_id")
		}

		boardID, err := strconv.ParseUint(boardIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
			c.Abort()
			return
		}

		permService := &services.PermissionService{}
		if !permService.HasBoardAccess(userID.(uint), uint(boardID)) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "You do not have access to this board",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
