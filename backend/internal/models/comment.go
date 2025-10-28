package models

import (
	"time"

	"gorm.io/gorm"
)

// Comment represents a comment on a card
type Comment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	CardID    uint           `gorm:"not null" json:"card_id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Card Card `gorm:"foreignKey:CardID" json:"card,omitempty"`
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
