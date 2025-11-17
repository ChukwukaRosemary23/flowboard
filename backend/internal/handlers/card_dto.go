package handlers

import "time"

// CreateCardRequest represents input for creating a card
type CreateCardRequest struct {
	Title       string     `json:"title" binding:"required,min=1,max=200"`
	Description string     `json:"description" binding:"max=2000"`
	ListID      uint       `json:"list_id" binding:"required"`
	DueDate     *time.Time `json:"due_date"` // Pointer = optional
}

// UpdateCardRequest represents input for updating a card
type UpdateCardRequest struct {
	Title       string     `json:"title" binding:"omitempty,min=1,max=200"`
	Description string     `json:"description" binding:"omitempty,max=2000"`
	Position    *int       `json:"position" binding:"omitempty,min=0"`
	DueDate     *time.Time `json:"due_date"`
}

type MoveCardRequest struct {
	ListID   uint `json:"list_id" binding:"required"`
	Position int  `json:"position" binding:"min=0"`
}

type CardDetailResponse struct {
	ID          uint                 `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	ListID      uint                 `json:"list_id"`
	Position    int                  `json:"position"`
	DueDate     *time.Time           `json:"due_date,omitempty"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	Members     []UserResponse       `json:"members"`
	Labels      []LabelResponse      `json:"labels"`
	Comments    []CommentResponse    `json:"comments"`
	Attachments []AttachmentResponse `json:"attachments"`
}

// LabelResponse for card labels (we'll implement labels later)
type LabelResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	BoardID uint   `json:"board_id"`
}

// CommentResponse for card comments (we'll implement later)
type CommentResponse struct {
	ID        uint         `json:"id"`
	Content   string       `json:"content"`
	CardID    uint         `json:"card_id"`
	User      UserResponse `json:"user"`
	CreatedAt time.Time    `json:"created_at"`
}

// AttachmentResponse for card file attachments (we'll implement later)
type AttachmentResponse struct {
	ID         uint         `json:"id"`
	Filename   string       `json:"filename"`
	FileURL    string       `json:"file_url"`
	FileSize   int64        `json:"file_size"`
	FileType   string       `json:"file_type"`
	CardID     uint         `json:"card_id"`
	UploadedBy UserResponse `json:"uploaded_by"`
	CreatedAt  time.Time    `json:"created_at"`
}
