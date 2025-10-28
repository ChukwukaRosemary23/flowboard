package handlers

import (
	"net/http"
	"strings"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/gin-gonic/gin"
)

// SearchCards searches cards across all user's boards
func SearchCards(c *gin.Context) {
	userID := c.GetUint("user_id")
	query := c.Query("q")                 // Search query
	boardID := c.Query("board")           // Optional: filter by board
	listID := c.Query("list")             // Optional: filter by list
	labelID := c.Query("label")           // Optional: filter by label
	memberID := c.Query("member")         // Optional: filter by assigned member
	hasDueDate := c.Query("has_due_date") // Optional: filter cards with due dates

	// Build query
	db := database.DB.Model(&models.Card{}).
		Joins("JOIN lists ON lists.id = cards.list_id").
		Joins("JOIN boards ON boards.id = lists.board_id").
		Where("boards.owner_id = ?", userID)

	// Search by title or description
	if query != "" {
		searchTerm := "%" + strings.ToLower(query) + "%"
		db = db.Where("LOWER(cards.title) LIKE ? OR LOWER(cards.description) LIKE ?", searchTerm, searchTerm)
	}

	// Filter by board
	if boardID != "" {
		db = db.Where("boards.id = ?", boardID)
	}

	// Filter by list
	if listID != "" {
		db = db.Where("cards.list_id = ?", listID)
	}

	// Filter by label
	if labelID != "" {
		db = db.Joins("JOIN card_labels ON card_labels.card_id = cards.id").
			Where("card_labels.label_id = ?", labelID)
	}

	// Filter by assigned member
	if memberID != "" {
		db = db.Joins("JOIN card_members ON card_members.card_id = cards.id").
			Where("card_members.user_id = ?", memberID)
	}

	// Filter by due date presence
	if hasDueDate == "true" {
		db = db.Where("cards.due_date IS NOT NULL")
	} else if hasDueDate == "false" {
		db = db.Where("cards.due_date IS NULL")
	}

	// Get cards
	var cards []models.Card
	if err := db.Preload("List").
		Preload("Members").
		Preload("Labels").
		Order("cards.created_at DESC").
		Limit(50). // Limit to 50 results
		Find(&cards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search cards"})
		return
	}

	// Convert to response
	response := make([]CardDetailResponse, len(cards))
	for i, card := range cards {
		// Members
		members := make([]UserResponse, len(card.Members))
		for j, member := range card.Members {
			members[j] = UserResponse{
				ID:        member.ID,
				Username:  member.Username,
				Email:     member.Email,
				AvatarURL: member.AvatarURL,
			}
		}

		// Labels
		labels := make([]LabelResponse, len(card.Labels))
		for j, label := range card.Labels {
			labels[j] = LabelResponse{
				ID:      label.ID,
				Name:    label.Name,
				Color:   label.Color,
				BoardID: label.BoardID,
			}
		}

		response[i] = CardDetailResponse{
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
			Comments:    []CommentResponse{},    // Don't load comments in search results
			Attachments: []AttachmentResponse{}, // Don't load attachments in search results
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"cards": response,
		"count": len(response),
		"query": query,
	})
}

// GetOverdueCards returns cards with due dates in the past
func GetOverdueCards(c *gin.Context) {
	userID := c.GetUint("user_id")

	var cards []models.Card
	if err := database.DB.
		Joins("JOIN lists ON lists.id = cards.list_id").
		Joins("JOIN boards ON boards.id = lists.board_id").
		Where("boards.owner_id = ? AND cards.due_date < NOW() AND cards.due_date IS NOT NULL", userID).
		Preload("List").
		Preload("Members").
		Preload("Labels").
		Order("cards.due_date ASC").
		Find(&cards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch overdue cards"})
		return
	}

	// Convert to response
	response := make([]CardDetailResponse, len(cards))
	for i, card := range cards {
		// Members
		members := make([]UserResponse, len(card.Members))
		for j, member := range card.Members {
			members[j] = UserResponse{
				ID:        member.ID,
				Username:  member.Username,
				Email:     member.Email,
				AvatarURL: member.AvatarURL,
			}
		}

		// Labels
		labels := make([]LabelResponse, len(card.Labels))
		for j, label := range card.Labels {
			labels[j] = LabelResponse{
				ID:      label.ID,
				Name:    label.Name,
				Color:   label.Color,
				BoardID: label.BoardID,
			}
		}

		response[i] = CardDetailResponse{
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
			Comments:    []CommentResponse{},
			Attachments: []AttachmentResponse{},
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"cards": response,
		"count": len(response),
	})
}

// GetUpcomingCards returns cards with due dates in the next 7 days
func GetUpcomingCards(c *gin.Context) {
	userID := c.GetUint("user_id")

	var cards []models.Card
	if err := database.DB.
		Joins("JOIN lists ON lists.id = cards.list_id").
		Joins("JOIN boards ON boards.id = lists.board_id").
		Where("boards.owner_id = ? AND cards.due_date BETWEEN NOW() AND NOW() + INTERVAL '7 days'", userID).
		Preload("List").
		Preload("Members").
		Preload("Labels").
		Order("cards.due_date ASC").
		Find(&cards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch upcoming cards"})
		return
	}

	// Convert to response
	response := make([]CardDetailResponse, len(cards))
	for i, card := range cards {
		// Members
		members := make([]UserResponse, len(card.Members))
		for j, member := range card.Members {
			members[j] = UserResponse{
				ID:        member.ID,
				Username:  member.Username,
				Email:     member.Email,
				AvatarURL: member.AvatarURL,
			}
		}

		// Labels
		labels := make([]LabelResponse, len(card.Labels))
		for j, label := range card.Labels {
			labels[j] = LabelResponse{
				ID:      label.ID,
				Name:    label.Name,
				Color:   label.Color,
				BoardID: label.BoardID,
			}
		}

		response[i] = CardDetailResponse{
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
			Comments:    []CommentResponse{},
			Attachments: []AttachmentResponse{},
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"cards": response,
		"count": len(response),
	})
}
