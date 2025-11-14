package handlers

import (
	"net/http"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateBoard handles creating a new board
func CreateBoard(c *gin.Context) {
	var req CreateBoardRequest

	// Validate input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context (set by auth middleware)
	userID := c.GetUint("user_id")

	// Set default background color if not provided
	backgroundColor := req.BackgroundColor
	if backgroundColor == "" {
		backgroundColor = "#0079BF" // Trello blue
	}

	// Wrap all database operations in a transaction
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// Create board
		board := models.Board{
			Title:           req.Title,
			Description:     req.Description,
			BackgroundColor: backgroundColor,
			OwnerID:         userID,
		}

		if err := tx.Create(&board).Error; err != nil {
			return err
		}

		// Get the owner role
		var ownerRole models.Role
		if err := tx.Where("name = ?", "owner").First(&ownerRole).Error; err != nil {
			return err // Critical error - role missing means migrations not run
		}

		// Add creator as owner in board_members table
		boardMember := models.BoardMember{
			BoardID: board.ID,
			UserID:  userID,
			RoleID:  ownerRole.ID,
			Status:  "active",
		}

		if err := tx.Create(&boardMember).Error; err != nil {
			return err
		}

		// // Success! Return the board
		// c.JSON(http.StatusCreated, BoardResponse{
		// 	ID:              board.ID,
		// 	Title:           board.Title,
		// 	Description:     board.Description,
		// 	BackgroundColor: board.BackgroundColor,
		// 	OwnerID:         board.OwnerID,
		// 	CreatedAt:       board.CreatedAt,
		// 	UpdatedAt:       board.UpdatedAt,
		// })

		// Success! Return the board
		c.JSON(http.StatusCreated, gin.H{
			"board": BoardResponse{
				ID:              board.ID,
				Title:           board.Title,
				Description:     board.Description,
				BackgroundColor: board.BackgroundColor,
				OwnerID:         board.OwnerID,
				CreatedAt:       board.CreatedAt,
				UpdatedAt:       board.UpdatedAt,
			},
			"message": "Board created successfully",
		})

		return nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create board"})
		return
	}
}

// GetBoards returns all boards the current user has access to
func GetBoards(c *gin.Context) {
	userID := c.GetUint("user_id")

	var boards []models.Board

	// Get all boards where user is a member (including owned boards)
	err := database.DB.
		Joins("JOIN board_members ON boards.id = board_members.board_id").
		Where("board_members.user_id = ? AND board_members.status = ?", userID, "active").
		Preload("Owner").
		Order("boards.created_at DESC").
		Find(&boards).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch boards"})
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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Board not found"})
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

// // UpdateBoard updates a board
// func UpdateBoard(c *gin.Context) {
// 	userID := c.GetUint("user_id")
// 	boardID := c.Param("id")

// 	var req UpdateBoardRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Find board and verify ownership
// 	var board models.Board
// 	if err := database.DB.Where("id = ? AND owner_id = ?", boardID, userID).First(&board).Error; err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Board not found"})
// 		return
// 	}

// UpdateBoard updates a board
func UpdateBoard(c *gin.Context) {
	// userID := c.GetUint("user_id")
	boardID := c.Param("id")

	var req UpdateBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find board (permission already checked by middleware)
	var board models.Board
	if err := database.DB.Where("id = ?", boardID).First(&board).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Board not found"})
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update board"})
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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Board not found"})
		return
	}

	// Soft delete (sets deleted_at timestamp)
	if err := database.DB.Delete(&board).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete board"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Board deleted successfully",
		"id":      boardID,
	})
}
