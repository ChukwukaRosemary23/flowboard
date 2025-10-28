package handlers

import "time"

// CreateListRequest represents input for creating a list
type CreateListRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=100"`
	BoardID uint   `json:"board_id" binding:"required"`
}

// UpdateListRequest represents input for updating a list
type UpdateListRequest struct {
	Title    string `json:"title" binding:"omitempty,min=1,max=100"`
	Position *int   `json:"position" binding:"omitempty,min=0"` // Pointer = can be null
}

// MoveListRequest for reordering lists
type MoveListRequest struct {
	Position int `json:"position" binding:"required,min=0"`
}

// ListDetailResponse includes cards
type ListDetailResponse struct {
	ID        uint           `json:"id"`
	Title     string         `json:"title"`
	BoardID   uint           `json:"board_id"`
	Position  int            `json:"position"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Cards     []CardResponse `json:"cards"`
}

// CardResponse represents card data (we'll use this now, implement later)
type CardResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ListID      uint       `json:"list_id"`
	Position    int        `json:"position"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}
