package models

import (
	"time"

	"gorm.io/gorm"
)

// Activity represents an action taken on the board
type Activity struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Action      string         `gorm:"not null" json:"action"` // "created_card", "moved_card", "added_comment", etc.
	EntityType  string         `gorm:"not null" json:"entity_type"` // "card", "list", "board", "comment"
	EntityID    uint           `gorm:"not null" json:"entity_id"`
	EntityTitle string         `json:"entity_title"` // Title of the card/list/board
	BoardID     uint           `gorm:"not null" json:"board_id"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	Metadata    string         `gorm:"type:text" json:"metadata"` // JSON string for extra data
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Board Board `gorm:"foreignKey:BoardID" json:"-"`
	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
}