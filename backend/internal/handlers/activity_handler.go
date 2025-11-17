package handlers

import (
	"net/http"
	"time"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/gin-gonic/gin"
)

// ActivityResponse represents activity data
type ActivityResponse struct {
	ID          uint         `json:"id"`
	Action      string       `json:"action"`
	EntityType  string       `json:"entity_type"`
	EntityID    uint         `json:"entity_id"`
	EntityTitle string       `json:"entity_title"`
	BoardID     uint         `json:"board_id"`
	User        UserResponse `json:"user"`
	Metadata    string       `json:"metadata"`
	CreatedAt   time.Time    `json:"created_at"`
}

// GetBoardActivities returns recent activities for a board
func GetBoardActivities(c *gin.Context) {
	boardID := c.Param("board_id")
	userID := c.GetUint("user_id")

	// Verify board exists and user owns it
	var board models.Board
	if err := database.DB.Where("id = ? AND owner_id = ?", boardID, userID).First(&board).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found or access denied"})
		return
	}

	// Get activities 
	var activities []models.Activity
	if err := database.DB.Where("board_id = ?", boardID).
		Preload("User").
		Order("created_at DESC").
		Limit(50).
		Find(&activities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activities"})
		return
	}

	// Convert to response
	response := make([]ActivityResponse, len(activities))
	for i, activity := range activities {
		response[i] = ActivityResponse{
			ID:          activity.ID,
			Action:      activity.Action,
			EntityType:  activity.EntityType,
			EntityID:    activity.EntityID,
			EntityTitle: activity.EntityTitle,
			BoardID:     activity.BoardID,
			User: UserResponse{
				ID:        activity.User.ID,
				Username:  activity.User.Username,
				Email:     activity.User.Email,
				AvatarURL: activity.User.AvatarURL,
			},
			Metadata:  activity.Metadata,
			CreatedAt: activity.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"activities": response,
		"count":      len(response),
	})
}

// GetUserActivities returns recent activities by a user
func GetUserActivities(c *gin.Context) {
	userID := c.GetUint("user_id")

	// Get user's activities across all their boards
	var activities []models.Activity
	if err := database.DB.
		Joins("JOIN boards ON boards.id = activities.board_id").
		Where("boards.owner_id = ?", userID).
		Preload("User").
		Order("activities.created_at DESC").
		Limit(50).
		Find(&activities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activities"})
		return
	}

	// Convert to response
	response := make([]ActivityResponse, len(activities))
	for i, activity := range activities {
		response[i] = ActivityResponse{
			ID:          activity.ID,
			Action:      activity.Action,
			EntityType:  activity.EntityType,
			EntityID:    activity.EntityID,
			EntityTitle: activity.EntityTitle,
			BoardID:     activity.BoardID,
			User: UserResponse{
				ID:        activity.User.ID,
				Username:  activity.User.Username,
				Email:     activity.User.Email,
				AvatarURL: activity.User.AvatarURL,
			},
			Metadata:  activity.Metadata,
			CreatedAt: activity.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"activities": response,
		"count":      len(response),
	})
}
