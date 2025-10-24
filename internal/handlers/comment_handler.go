package handlers

import (
	"net/http"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateComment adds a comment to a card
func CreateComment(c *gin.Context) {
	cardID := c.Param("card_id")
	userID := c.GetUint("user_id")

	var req struct {
		Content string `json:"content" binding:"required,min=1,max=2000"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify card exists and user has access
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

	// Create comment
	comment := models.Comment{
		Content: req.Content,
		CardID:  card.ID,
		UserID:  userID,
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	// Load user info for response
	database.DB.Preload("User").First(&comment, comment.ID)

	c.JSON(http.StatusCreated, CommentResponse{
		ID:      comment.ID,
		Content: comment.Content,
		CardID:  comment.CardID,
		User: UserResponse{
			ID:        comment.User.ID,
			Username:  comment.User.Username,
			Email:     comment.User.Email,
			AvatarURL: comment.User.AvatarURL,
		},
		CreatedAt: comment.CreatedAt,
	})
}

// GetComments returns all comments for a card
func GetComments(c *gin.Context) {
	cardID := c.Param("card_id")
	userID := c.GetUint("user_id")

	// Verify card exists and user has access
	var card models.Card
	if err := database.DB.Preload("List.Board").First(&card, cardID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	if card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Get all comments ordered by creation date
	var comments []models.Comment
	if err := database.DB.Where("card_id = ?", cardID).
		Preload("User").
		Order("created_at ASC").
		Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	// Convert to response
	response := make([]CommentResponse, len(comments))
	for i, comment := range comments {
		response[i] = CommentResponse{
			ID:      comment.ID,
			Content: comment.Content,
			CardID:  comment.CardID,
			User: UserResponse{
				ID:        comment.User.ID,
				Username:  comment.User.Username,
				Email:     comment.User.Email,
				AvatarURL: comment.User.AvatarURL,
			},
			CreatedAt: comment.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": response,
		"count":    len(response),
	})
}

// UpdateComment updates a comment
func UpdateComment(c *gin.Context) {
	commentID := c.Param("id")
	userID := c.GetUint("user_id")

	var req struct {
		Content string `json:"content" binding:"required,min=1,max=2000"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find comment
	var comment models.Comment
	if err := database.DB.Preload("Card.List.Board").First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Verify user is comment author OR board owner
	if comment.UserID != userID && comment.Card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Update comment
	comment.Content = req.Content
	if err := database.DB.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	// Load user for response
	database.DB.Preload("User").First(&comment, comment.ID)

	c.JSON(http.StatusOK, CommentResponse{
		ID:      comment.ID,
		Content: comment.Content,
		CardID:  comment.CardID,
		User: UserResponse{
			ID:        comment.User.ID,
			Username:  comment.User.Username,
			Email:     comment.User.Email,
			AvatarURL: comment.User.AvatarURL,
		},
		CreatedAt: comment.CreatedAt,
	})
}

// DeleteComment deletes a comment
func DeleteComment(c *gin.Context) {
	commentID := c.Param("id")
	userID := c.GetUint("user_id")

	// Find comment
	var comment models.Comment
	if err := database.DB.Preload("Card.List.Board").First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Verify user is comment author OR board owner
	if comment.UserID != userID && comment.Card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Soft delete
	if err := database.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted successfully",
		"id":      commentID,
	})
}
