package handlers

import (
	"net/http"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateBoard handles creating a new board
func CreateBoard(c *gin.Context) {
	var req CreateBoardRequest

	// Validate input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context (set by auth middleware)
	userID := c.GetUint("user_id")

	// Set default background color if not provided
	backgroundColor := req.BackgroundColor
	if backgroundColor == "" {
		backgroundColor = "#0079BF" // Trello blue
	}

	// Create board
	board := models.Board{
		Title:           req.Title,
		Description:     req.Description,
		BackgroundColor: backgroundColor,
		OwnerID:         userID,
	}

	if err := database.DB.Create(&board).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create board"})
		return
	}

	// Return response
	c.JSON(http.StatusCreated, BoardResponse{
		ID:              board.ID,
		Title:           board.Title,
		Description:     board.Description,
		BackgroundColor: board.BackgroundColor,
		OwnerID:         board.OwnerID,
		CreatedAt:       board.CreatedAt,
		UpdatedAt:       board.UpdatedAt,
	})
}

// GetBoards returns all boards owned by the current user
func GetBoards(c *gin.Context) {
	userID := c.GetUint("user_id")

	var boards []models.Board
	if err := database.DB.Where("owner_id = ?", userID).Order("created_at DESC").Find(&boards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch boards"})
		return
	}

	// Convert to response format
	response := make([]BoardResponse, len(boards))
	for i, board := range boards {
		response[i] = BoardResponse{
			ID:              board.ID,
			Title:           board.Title,
			Description:     board.Description,
			BackgroundColor: board.BackgroundColor,
			OwnerID:         board.OwnerID,
			CreatedAt:       board.CreatedAt,
			UpdatedAt:       board.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"boards": response,
		"count":  len(response),
	})
}

// GetBoard returns a single board by ID with its lists
func GetBoard(c *gin.Context) {
	userID := c.GetUint("user_id")
	boardID := c.Param("id")

	var board models.Board
	// Find board and verify ownership
	if err := database.DB.Where("id = ? AND owner_id = ?", boardID, userID).
		Preload("Lists").
		First(&board).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found"})
		return
	}

	// Convert lists to response format
	lists := make([]ListResponse, len(board.Lists))
	for i, list := range board.Lists {
		lists[i] = ListResponse{
			ID:       list.ID,
			Title:    list.Title,
			BoardID:  list.BoardID,
			Position: list.Position,
		}
	}

	// Return response
	c.JSON(http.StatusOK, BoardDetailResponse{
		BoardResponse: BoardResponse{
			ID:              board.ID,
			Title:           board.Title,
			Description:     board.Description,
			BackgroundColor: board.BackgroundColor,
			OwnerID:         board.OwnerID,
			CreatedAt:       board.CreatedAt,
			UpdatedAt:       board.UpdatedAt,
		},
		Lists: lists,
	})
}

// UpdateBoard updates a board
func UpdateBoard(c *gin.Context) {
	userID := c.GetUint("user_id")
	boardID := c.Param("id")

	var req UpdateBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find board and verify ownership
	var board models.Board
	if err := database.DB.Where("id = ? AND owner_id = ?", boardID, userID).First(&board).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found"})
		return
	}

	// Update fields (only if provided)
	if req.Title != "" {
		board.Title = req.Title
	}
	if req.Description != "" {
		board.Description = req.Description
	}
	if req.BackgroundColor != "" {
		board.BackgroundColor = req.BackgroundColor
	}

	// Save changes
	if err := database.DB.Save(&board).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update board"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, BoardResponse{
		ID:              board.ID,
		Title:           board.Title,
		Description:     board.Description,
		BackgroundColor: board.BackgroundColor,
		OwnerID:         board.OwnerID,
		CreatedAt:       board.CreatedAt,
		UpdatedAt:       board.UpdatedAt,
	})
}

// DeleteBoard deletes a board (soft delete)
func DeleteBoard(c *gin.Context) {
	userID := c.GetUint("user_id")
	boardID := c.Param("id")

	// Find board and verify ownership
	var board models.Board
	if err := database.DB.Where("id = ? AND owner_id = ?", boardID, userID).First(&board).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found"})
		return
	}

	// Soft delete (sets deleted_at timestamp)
	if err := database.DB.Delete(&board).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete board"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Board deleted successfully",
		"id":      boardID,
	})
}
