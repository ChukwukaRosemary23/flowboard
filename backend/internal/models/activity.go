package models

import (
	"time"

	"gorm.io/gorm"
)

// Activity represents an action taken on the board
type Activity struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Action      string         `gorm:"not null" json:"action"`
	EntityType  string         `gorm:"not null" json:"entity_type"`
	EntityID    uint           `gorm:"not null" json:"entity_id"`
	EntityTitle string         `json:"entity_title"`
	BoardID     uint           `gorm:"not null" json:"board_id"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	Metadata    string         `gorm:"type:text" json:"metadata"` // JSON string for extra data
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Board Board `gorm:"foreignKey:BoardID" json:"-"`
	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
