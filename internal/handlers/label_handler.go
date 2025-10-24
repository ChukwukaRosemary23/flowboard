package handlers

import (
	"net/http"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateLabel creates a label for a board
func CreateLabel(c *gin.Context) {
	boardID := c.Param("board_id")
	userID := c.GetUint("user_id")

	var req struct {
		Name  string `json:"name" binding:"required,min=1,max=50"`
		Color string `json:"color" binding:"required,hexcolor"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify board exists and user owns it
	var board models.Board
	if err := database.DB.Where("id = ? AND owner_id = ?", boardID, userID).First(&board).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found or access denied"})
		return
	}

	// Create label
	label := models.Label{
		Name:    req.Name,
		Color:   req.Color,
		BoardID: board.ID,
	}

	if err := database.DB.Create(&label).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create label"})
		return
	}

	c.JSON(http.StatusCreated, LabelResponse{
		ID:      label.ID,
		Name:    label.Name,
		Color:   label.Color,
		BoardID: label.BoardID,
	})
}

// GetLabels returns all labels for a board
func GetLabels(c *gin.Context) {
	boardID := c.Param("board_id")
	userID := c.GetUint("user_id")

	// Verify board exists and user owns it
	var board models.Board
	if err := database.DB.Where("id = ? AND owner_id = ?", boardID, userID).First(&board).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found or access denied"})
		return
	}

	// Get all labels for board
	var labels []models.Label
	if err := database.DB.Where("board_id = ?", boardID).Find(&labels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch labels"})
		return
	}

	// Convert to response
	response := make([]LabelResponse, len(labels))
	for i, label := range labels {
		response[i] = LabelResponse{
			ID:      label.ID,
			Name:    label.Name,
			Color:   label.Color,
			BoardID: label.BoardID,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"labels": response,
		"count":  len(response),
	})
}

// UpdateLabel updates a label
func UpdateLabel(c *gin.Context) {
	labelID := c.Param("id")
	userID := c.GetUint("user_id")

	var req struct {
		Name  string `json:"name" binding:"omitempty,min=1,max=50"`
		Color string `json:"color" binding:"omitempty,hexcolor"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find label
	var label models.Label
	if err := database.DB.Preload("Board").First(&label, labelID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Label not found"})
		return
	}

	// Verify user owns the board
	if label.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Update fields
	if req.Name != "" {
		label.Name = req.Name
	}
	if req.Color != "" {
		label.Color = req.Color
	}

	if err := database.DB.Save(&label).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update label"})
		return
	}

	c.JSON(http.StatusOK, LabelResponse{
		ID:      label.ID,
		Name:    label.Name,
		Color:   label.Color,
		BoardID: label.BoardID,
	})
}

// DeleteLabel deletes a label
func DeleteLabel(c *gin.Context) {
	labelID := c.Param("id")
	userID := c.GetUint("user_id")

	// Find label
	var label models.Label
	if err := database.DB.Preload("Board").First(&label, labelID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Label not found"})
		return
	}

	// Verify user owns the board
	if label.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Delete label (will automatically remove from cards via many-to-many)
	if err := database.DB.Delete(&label).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete label"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Label deleted successfully",
		"id":      labelID,
	})
}

// AddLabelToCard attaches a label to a card
func AddLabelToCard(c *gin.Context) {
	cardID := c.Param("card_id")
	userID := c.GetUint("user_id")

	var req struct {
		LabelID uint `json:"label_id" binding:"required"`
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

	// Find label
	var label models.Label
	if err := database.DB.First(&label, req.LabelID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Label not found"})
		return
	}

	// Verify label belongs to same board
	if label.BoardID != card.List.Board.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Label must belong to the same board"})
		return
	}

	// Add label to card (many-to-many)
	if err := database.DB.Model(&card).Association("Labels").Append(&label); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add label to card"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Label added to card successfully",
		"card_id":  cardID,
		"label_id": req.LabelID,
	})
}

// RemoveLabelFromCard removes a label from a card
func RemoveLabelFromCard(c *gin.Context) {
	cardID := c.Param("card_id")
	labelID := c.Param("label_id")
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

	// Find label
	var label models.Label
	if err := database.DB.First(&label, labelID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Label not found"})
		return
	}

	// Remove label from card
	if err := database.DB.Model(&card).Association("Labels").Delete(&label); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove label from card"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Label removed from card successfully",
		"card_id":  cardID,
		"label_id": labelID,
	})
}
