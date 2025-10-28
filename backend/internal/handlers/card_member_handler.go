package handlers

import (
	"net/http"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/gin-gonic/gin"
)

// AssignMemberToCard assigns a user to a card
func AssignMemberToCard(c *gin.Context) {
	cardID := c.Param("card_id")
	userID := c.GetUint("user_id")

	var req struct {
		MemberID uint `json:"member_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find card
	var card models.Card
	if err := database.DB.Preload("List.Board").First(&card, cardID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	// Verify user owns the board
	if card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Verify member exists
	var member models.User
	if err := database.DB.First(&member, req.MemberID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Check if member is already assigned
	var existingAssignment models.CardMember
	if err := database.DB.Where("card_id = ? AND user_id = ?", cardID, req.MemberID).First(&existingAssignment).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User is already assigned to this card"})
		return
	}

	// Assign member to card (many-to-many)
	if err := database.DB.Model(&card).Association("Members").Append(&member); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign member to card"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Member assigned to card successfully",
		"card_id":   cardID,
		"member_id": req.MemberID,
		"member": UserResponse{
			ID:        member.ID,
			Username:  member.Username,
			Email:     member.Email,
			AvatarURL: member.AvatarURL,
		},
	})
}

// GetCardMembers returns all members assigned to a card
func GetCardMembers(c *gin.Context) {
	cardID := c.Param("card_id")
	userID := c.GetUint("user_id")

	// Find card
	var card models.Card
	if err := database.DB.Preload("List.Board").Preload("Members").First(&card, cardID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	// Verify user owns the board
	if card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Convert to response
	members := make([]UserResponse, len(card.Members))
	for i, member := range card.Members {
		members[i] = UserResponse{
			ID:        member.ID,
			Username:  member.Username,
			Email:     member.Email,
			AvatarURL: member.AvatarURL,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"members": members,
		"count":   len(members),
	})
}

// UnassignMemberFromCard removes a user from a card
func UnassignMemberFromCard(c *gin.Context) {
	cardID := c.Param("card_id")
	memberID := c.Param("member_id")
	userID := c.GetUint("user_id")

	// Find card
	var card models.Card
	if err := database.DB.Preload("List.Board").First(&card, cardID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	// Verify user owns the board
	if card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Find member
	var member models.User
	if err := database.DB.First(&member, memberID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Remove member from card
	if err := database.DB.Model(&card).Association("Members").Delete(&member); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unassign member from card"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Member unassigned from card successfully",
		"card_id":   cardID,
		"member_id": memberID,
	})
}
