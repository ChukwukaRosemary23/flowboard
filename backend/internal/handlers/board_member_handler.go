package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/gin-gonic/gin"
)

// InviteMember adds a user to a board with a specific role
func InviteMember(c *gin.Context) {
	boardID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	inviterID, _ := c.Get("user_id")

	var req struct {
		UserID uint   `json:"user_id" binding:"required"`
		Role   string `json:"role" binding:"required,oneof=admin member viewer"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists
	var user models.User
	if err := database.DB.First(&user, req.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Check if user is already a member
	var existingMember models.BoardMember
	err := database.DB.Where("board_id = ? AND user_id = ?", boardID, req.UserID).First(&existingMember).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User is already a member of this board"})
		return
	}

	// Get role ID
	var role models.Role
	if err := database.DB.Where("name = ?", req.Role).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	// Create board member
	inviterIDUint := inviterID.(uint)
	boardMember := models.BoardMember{
		BoardID:   uint(boardID),
		UserID:    req.UserID,
		RoleID:    role.ID,
		InvitedBy: &inviterIDUint,
		InvitedAt: time.Now(),
		Status:    "active",
	}

	if err := database.DB.Create(&boardMember).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add member"})
		return
	}

	// Preload relationships for response
	database.DB.Preload("User").Preload("Role").First(&boardMember, boardMember.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Member added successfully",
		"member":  boardMember,
	})
}

// RemoveMember removes a user from a board
func RemoveMember(c *gin.Context) {
	boardID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	userID, _ := strconv.ParseUint(c.Param("user_id"), 10, 32)

	// Find the board member
	var boardMember models.BoardMember
	err := database.DB.Preload("Role").
		Where("board_id = ? AND user_id = ?", boardID, userID).
		First(&boardMember).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	// Cannot remove owner
	if boardMember.Role.Name == "owner" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot remove board owner"})
		return
	}

	// Update status to removed instead of deleting
	boardMember.Status = "removed"
	boardMember.UpdatedAt = time.Now()
	database.DB.Save(&boardMember)

	c.JSON(http.StatusOK, gin.H{"message": "Member removed successfully"})
}

// UpdateMemberRole changes a member's role
func UpdateMemberRole(c *gin.Context) {
	boardID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	userID, _ := strconv.ParseUint(c.Param("user_id"), 10, 32)

	var req struct {
		Role string `json:"role" binding:"required,oneof=admin member viewer"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the board member
	var boardMember models.BoardMember
	err := database.DB.Preload("Role").
		Where("board_id = ? AND user_id = ?", boardID, userID).
		First(&boardMember).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	// Cannot change owner role
	if boardMember.Role.Name == "owner" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot change owner role"})
		return
	}

	// Get new role ID
	var role models.Role
	if err := database.DB.Where("name = ?", req.Role).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	// Update member role
	boardMember.RoleID = role.ID
	boardMember.UpdatedAt = time.Now()
	database.DB.Save(&boardMember)

	// Preload for response
	database.DB.Preload("User").Preload("Role").First(&boardMember, boardMember.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Role updated successfully",
		"member":  boardMember,
	})
}

// GetBoardMembers returns all members of a board
func GetBoardMembers(c *gin.Context) {
	boardID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var members []models.BoardMember
	database.DB.Preload("User").
		Preload("Role").
		Where("board_id = ? AND status = ?", boardID, "active").
		Order("role_id ASC, created_at ASC").
		Find(&members)

	c.JSON(http.StatusOK, gin.H{
		"count":   len(members),
		"members": members,
	})
}
