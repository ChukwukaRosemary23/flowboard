package handlers

import (
	"net/http"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/utils"
	ws "github.com/ChukwukaRosemary23/flowboard-backend/internal/websocket"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var WSHub *ws.Hub

// CreateCard creates a new card in a list
func CreateCard(c *gin.Context) {
	var req CreateCardRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	// Verify list exists and user owns the board
	var list models.List
	if err := database.DB.Preload("Board").First(&list, req.ListID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}

	// Verify user owns the board
	if list.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Get the highest position in this list
	var maxPosition int
	database.DB.Model(&models.Card{}).
		Where("list_id = ?", req.ListID).
		Select("COALESCE(MAX(position), -1)").
		Scan(&maxPosition)

	// Create card at the end
	card := models.Card{
		Title:       req.Title,
		Description: req.Description,
		ListID:      req.ListID,
		Position:    maxPosition + 1,
		DueDate:     req.DueDate,
	}

	if err := database.DB.Create(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create card"})
		return
	}

	// Log activity
	utils.LogActivity("created_card", "card", card.ID, list.Board.ID, userID, card.Title, nil)

	// Broadcast to WebSocket clients
	if WSHub != nil {
		WSHub.BroadcastToBoard(list.Board.ID, "card_created", CardResponse{
			ID:          card.ID,
			Title:       card.Title,
			Description: card.Description,
			ListID:      card.ListID,
			Position:    card.Position,
			DueDate:     card.DueDate,
			CreatedAt:   card.CreatedAt,
		})
	}

	c.JSON(http.StatusCreated, CardResponse{
		ID:          card.ID,
		Title:       card.Title,
		Description: card.Description,
		ListID:      card.ListID,
		Position:    card.Position,
		DueDate:     card.DueDate,
		CreatedAt:   card.CreatedAt,
	})
}

// GetCards returns all cards in a list
func GetCards(c *gin.Context) {
	listID := c.Param("list_id")
	userID := c.GetUint("user_id")

	// Verify list exists and user owns the board
	var list models.List
	if err := database.DB.Preload("Board").First(&list, listID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}

	if list.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Get all cards ordered by position
	var cards []models.Card
	if err := database.DB.Where("list_id = ?", listID).
		Order("position ASC").
		Find(&cards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cards"})
		return
	}

	// Convert to response
	response := make([]CardResponse, len(cards))
	for i, card := range cards {
		response[i] = CardResponse{
			ID:          card.ID,
			Title:       card.Title,
			Description: card.Description,
			ListID:      card.ListID,
			Position:    card.Position,
			DueDate:     card.DueDate,
			CreatedAt:   card.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"cards": response,
		"count": len(response),
	})
}

// GetCard returns a single card with all details
func GetCard(c *gin.Context) {
	cardID := c.Param("id")
	userID := c.GetUint("user_id")

	// Find card with all relationships
	var card models.Card
	if err := database.DB.
		Preload("List.Board").
		Preload("Members").
		Preload("Labels").
		Preload("Comments.User").
		Preload("Attachments.Uploader").
		First(&card, cardID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	// Verify user owns the board
	if card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Convert members to response
	members := make([]UserResponse, len(card.Members))
	for i, member := range card.Members {
		members[i] = UserResponse{
			ID:        member.ID,
			Username:  member.Username,
			Email:     member.Email,
			AvatarURL: member.AvatarURL,
		}
	}

	// Convert labels to response
	labels := make([]LabelResponse, len(card.Labels))
	for i, label := range card.Labels {
		labels[i] = LabelResponse{
			ID:      label.ID,
			Name:    label.Name,
			Color:   label.Color,
			BoardID: label.BoardID,
		}
	}

	// Convert comments to response
	comments := make([]CommentResponse, len(card.Comments))
	for i, comment := range card.Comments {
		comments[i] = CommentResponse{
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

	// Convert attachments to response
	attachments := make([]AttachmentResponse, len(card.Attachments))
	for i, attachment := range card.Attachments {
		attachments[i] = AttachmentResponse{
			ID:       attachment.ID,
			Filename: attachment.Filename,
			FileURL:  attachment.FileURL,
			FileSize: attachment.FileSize,
			FileType: attachment.FileType,
			CardID:   attachment.CardID,
			UploadedBy: UserResponse{
				ID:        attachment.Uploader.ID,
				Username:  attachment.Uploader.Username,
				Email:     attachment.Uploader.Email,
				AvatarURL: attachment.Uploader.AvatarURL,
			},
			CreatedAt: attachment.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, CardDetailResponse{
		ID:          card.ID,
		Title:       card.Title,
		Description: card.Description,
		ListID:      card.ListID,
		Position:    card.Position,
		DueDate:     card.DueDate,
		CreatedAt:   card.CreatedAt,
		UpdatedAt:   card.UpdatedAt,
		Members:     members,
		Labels:      labels,
		Comments:    comments,
		Attachments: attachments,
	})
}

// UpdateCard updates a card
func UpdateCard(c *gin.Context) {
	cardID := c.Param("id")
	userID := c.GetUint("user_id")

	var req UpdateCardRequest
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

	// Verify ownership
	if card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Update fields
	if req.Title != "" {
		card.Title = req.Title
	}
	if req.Description != "" {
		card.Description = req.Description
	}
	if req.Position != nil {
		card.Position = *req.Position
	}
	// DueDate can be set or cleared
	card.DueDate = req.DueDate

	if err := database.DB.Save(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update card"})
		return
	}

	c.JSON(http.StatusOK, CardResponse{
		ID:          card.ID,
		Title:       card.Title,
		Description: card.Description,
		ListID:      card.ListID,
		Position:    card.Position,
		DueDate:     card.DueDate,
		CreatedAt:   card.CreatedAt,
	})
}

// MoveCard moves a card to a different list or position
func MoveCard(c *gin.Context) {
	cardID := c.Param("id")
	userID := c.GetUint("user_id")

	var req MoveCardRequest
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

	// Verify ownership of source board
	if card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Verify destination list exists and user owns that board too
	var destList models.List
	if err := database.DB.Preload("Board").First(&destList, req.ListID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Destination list not found"})
		return
	}

	if destList.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied to destination list"})
		return
	}

	oldListID := card.ListID
	oldPosition := card.Position
	newListID := req.ListID
	newPosition := req.Position

	// Moving to different list
	if oldListID != newListID {
		// Remove from old list: decrease position of cards after this one
		database.DB.Model(&models.Card{}).
			Where("list_id = ? AND position > ?", oldListID, oldPosition).
			UpdateColumn("position", gorm.Expr("position - 1"))

		// Make space in new list: increase position of cards at/after new position
		database.DB.Model(&models.Card{}).
			Where("list_id = ? AND position >= ?", newListID, newPosition).
			UpdateColumn("position", gorm.Expr("position + 1"))

		// Update card
		card.ListID = newListID
		card.Position = newPosition
	} else {
		// Moving within same list
		if newPosition > oldPosition {
			// Moving down: decrease position of cards in between
			database.DB.Model(&models.Card{}).
				Where("list_id = ? AND position > ? AND position <= ?", oldListID, oldPosition, newPosition).
				UpdateColumn("position", gorm.Expr("position - 1"))
		} else if newPosition < oldPosition {
			// Moving up: increase position of cards in between
			database.DB.Model(&models.Card{}).
				Where("list_id = ? AND position >= ? AND position < ?", oldListID, newPosition, oldPosition).
				UpdateColumn("position", gorm.Expr("position + 1"))
		}

		card.Position = newPosition
	}

	database.DB.Save(&card)

	// Broadcast to WebSocket clients
	if WSHub != nil {
		WSHub.BroadcastToBoard(card.List.Board.ID, "card_moved", gin.H{
			"card_id":      card.ID,
			"old_list_id":  oldListID,
			"new_list_id":  card.ListID,
			"new_position": card.Position,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Card moved successfully",
		"id":           card.ID,
		"new_list_id":  card.ListID,
		"new_position": card.Position,
	})
}

// DeleteCard deletes a card
func DeleteCard(c *gin.Context) {
	cardID := c.Param("id")
	userID := c.GetUint("user_id")

	// Find card
	var card models.Card
	if err := database.DB.Preload("List.Board").First(&card, cardID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	// Verify ownership
	if card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Store board ID before deleting
	boardID := card.List.Board.ID
	listID := card.ListID

	// Soft delete
	if err := database.DB.Delete(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete card"})
		return
	}

	// Broadcast to WebSocket clients
	if WSHub != nil {
		WSHub.BroadcastToBoard(boardID, "card_deleted", gin.H{
			"card_id": cardID,
			"list_id": listID,
		})
	}

	// Adjust positions of remaining cards
	database.DB.Model(&models.Card{}).
		Where("list_id = ? AND position > ?", listID, card.Position).
		UpdateColumn("position", gorm.Expr("position - 1"))

	c.JSON(http.StatusOK, gin.H{
		"message": "Card deleted successfully",
		"id":      cardID,
	})
}
