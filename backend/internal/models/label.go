package models

import (
	"time"

	"gorm.io/gorm"
)

// Label represents a colored tag for categorizing cards
type Label struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Color     string         `gorm:"not null" json:"color"` //color like #FF5733
	BoardID   uint           `gorm:"not null" json:"board_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Board Board  `gorm:"foreignKey:BoardID" json:"board,omitempty"`
	Cards []Card `gorm:"many2many:card_labels" json:"-"`
}
