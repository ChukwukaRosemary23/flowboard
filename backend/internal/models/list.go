package models

import (
	"time"

	"gorm.io/gorm"
)

// List represents a column in a board (like "To Do", "In Progress", "Done")
type List struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	BoardID   uint           `gorm:"not null" json:"board_id"`
	Position  int            `gorm:"not null;default:0" json:"position"` // Order of lists (0, 1, 2...)
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Board Board  `gorm:"foreignKey:BoardID" json:"board,omitempty"`
	Cards []Card `gorm:"foreignKey:ListID;constraint:OnDelete:CASCADE" json:"cards,omitempty"`
}
