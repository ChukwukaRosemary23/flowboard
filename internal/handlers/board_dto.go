package handlers

import "time"

// CreateBoardRequest represents input for creating a board
type CreateBoardRequest struct {
	Title           string `json:"title" binding:"required,min=1,max=100"`
	Description     string `json:"description" binding:"max=500"`
	BackgroundColor string `json:"background_color" binding:"omitempty,hexcolor"`
}

// UpdateBoardRequest represents input for updating a board
type UpdateBoardRequest struct {
	Title           string `json:"title" binding:"omitempty,min=1,max=100"`
	Description     string `json:"description" binding:"omitempty,max=500"`
	BackgroundColor string `json:"background_color" binding:"omitempty,hexcolor"`
}

// BoardResponse represents board data returned to client
type BoardResponse struct {
	ID              uint      `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	BackgroundColor string    `json:"background_color"`
	OwnerID         uint      `json:"owner_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// BoardDetailResponse includes lists (for single board view)
type BoardDetailResponse struct {
	BoardResponse
	Lists []ListResponse `json:"lists"`
}

// ListResponse represents list data (we'll use this later)
type ListResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	BoardID  uint   `json:"board_id"`
	Position int    `json:"position"`
}
