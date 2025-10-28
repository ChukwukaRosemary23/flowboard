package handlers

import (
	"net/http"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateList creates a new list in a board
func CreateList(c *gin.Context) {
	var req CreateListRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	// Verify board exists and user owns it
	var board models.Board
	if err := database.DB.Where("id = ? AND owner_id = ?", req.BoardID, userID).First(&board).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found or you don't have access"})
		return
	}

	// Get the highest position in this board
	var maxPosition int
	database.DB.Model(&models.List{}).
		Where("board_id = ?", req.BoardID).
		Select("COALESCE(MAX(position), -1)").
		Scan(&maxPosition)

	// Create list at the end (max position + 1)
	list := models.List{
		Title:    req.Title,
		BoardID:  req.BoardID,
		Position: maxPosition + 1,
	}

	if err := database.DB.Create(&list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create list"})
		return
	}

	c.JSON(http.StatusCreated, ListResponse{
		ID:       list.ID,
		Title:    list.Title,
		BoardID:  list.BoardID,
		Position: list.Position,
	})
}

// GetLists returns all lists for a board
func GetLists(c *gin.Context) {
	boardID := c.Param("board_id")
	userID := c.GetUint("user_id")

	// Verify board exists and user owns it
	var board models.Board
	if err := database.DB.Where("id = ? AND owner_id = ?", boardID, userID).First(&board).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found or you don't have access"})
		return
	}

	// Get all lists ordered by position
	var lists []models.List
	if err := database.DB.Where("board_id = ?", boardID).
		Order("position ASC").
		Find(&lists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch lists"})
		return
	}

	// Convert to response
	response := make([]ListResponse, len(lists))
	for i, list := range lists {
		response[i] = ListResponse{
			ID:       list.ID,
			Title:    list.Title,
			BoardID:  list.BoardID,
			Position: list.Position,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"lists": response,
		"count": len(response),
	})
}

// GetList returns a single list with its cards
func GetList(c *gin.Context) {
	listID := c.Param("id")
	userID := c.GetUint("user_id")

	// Find list and verify user owns the board
	var list models.List
	if err := database.DB.Preload("Cards", func(db *gorm.DB) *gorm.DB {
		return db.Order("position ASC")
	}).First(&list, listID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}

	// Verify user owns the board
	var board models.Board
	if err := database.DB.Where("id = ? AND owner_id = ?", list.BoardID, userID).First(&board).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Access denied"})
		return
	}

	// Convert cards to response
	cards := make([]CardResponse, len(list.Cards))
	for i, card := range list.Cards {
		cards[i] = CardResponse{
			ID:          card.ID,
			Title:       card.Title,
			Description: card.Description,
			ListID:      card.ListID,
			Position:    card.Position,
			DueDate:     card.DueDate,
			CreatedAt:   card.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, ListDetailResponse{
		ID:        list.ID,
		Title:     list.Title,
		BoardID:   list.BoardID,
		Position:  list.Position,
		CreatedAt: list.CreatedAt,
		UpdatedAt: list.UpdatedAt,
		Cards:     cards,
	})
}

// UpdateList updates a list
func UpdateList(c *gin.Context) {
	listID := c.Param("id")
	userID := c.GetUint("user_id")

	var req UpdateListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find list
	var list models.List
	if err := database.DB.First(&list, listID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}

	// Verify user owns the board
	var board models.Board
	if err := database.DB.Where("id = ? AND owner_id = ?", list.BoardID, userID).First(&board).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Access denied"})
		return
	}

	// Update fields
	if req.Title != "" {
		list.Title = req.Title
	}
	if req.Position != nil {
		list.Position = *req.Position
	}

	if err := database.DB.Save(&list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update list"})
		return
	}

	c.JSON(http.StatusOK, ListResponse{
		ID:       list.ID,
		Title:    list.Title,
		BoardID:  list.BoardID,
		Position: list.Position,
	})
}

// MoveList changes list position (for drag and drop)
func MoveList(c *gin.Context) {
	listID := c.Param("id")
	userID := c.GetUint("user_id")

	var req MoveListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find list
	var list models.List
	if err := database.DB.First(&list, listID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}

	// Verify ownership
	var board models.Board
	if err := database.DB.Where("id = ? AND owner_id = ?", list.BoardID, userID).First(&board).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Access denied"})
		return
	}

	oldPosition := list.Position
	newPosition := req.Position

	// Update positions of affected lists
	if newPosition > oldPosition {
		// Moving right: decrease position of lists in between
		database.DB.Model(&models.List{}).
			Where("board_id = ? AND position > ? AND position <= ?", list.BoardID, oldPosition, newPosition).
			UpdateColumn("position", gorm.Expr("position - 1"))
	} else if newPosition < oldPosition {
		// Moving left: increase position of lists in between
		database.DB.Model(&models.List{}).
			Where("board_id = ? AND position >= ? AND position < ?", list.BoardID, newPosition, oldPosition).
			UpdateColumn("position", gorm.Expr("position + 1"))
	}

	// Update list position
	list.Position = newPosition
	database.DB.Save(&list)

	c.JSON(http.StatusOK, gin.H{
		"message":      "List moved successfully",
		"id":           list.ID,
		"new_position": newPosition,
	})
}

// DeleteList deletes a list (soft delete)
func DeleteList(c *gin.Context) {
	listID := c.Param("id")
	userID := c.GetUint("user_id")

	// Find list
	var list models.List
	if err := database.DB.First(&list, listID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}

	// Verify ownership
	var board models.Board
	if err := database.DB.Where("id = ? AND owner_id = ?", list.BoardID, userID).First(&board).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Access denied"})
		return
	}

	// Soft delete (CASCADE will delete all cards)
	if err := database.DB.Delete(&list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete list"})
		return
	}

	// Adjust positions of remaining lists
	database.DB.Model(&models.List{}).
		Where("board_id = ? AND position > ?", list.BoardID, list.Position).
		UpdateColumn("position", gorm.Expr("position - 1"))

	c.JSON(http.StatusOK, gin.H{
		"message": "List deleted successfully",
		"id":      listID,
	})
}
