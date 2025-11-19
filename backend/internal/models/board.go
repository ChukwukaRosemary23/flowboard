package models

import (
	"time"

	"gorm.io/gorm"
)

// Board represents a project board (like "Website Redesign")
type Board struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Title           string         `gorm:"not null" json:"title"`
	Description     string         `json:"description"`
	BackgroundColor string         `gorm:"default:#0079BF" json:"background_color"`
	OwnerID         uint           `gorm:"not null" json:"owner_id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Owner  User    `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
	Lists  []List  `gorm:"foreignKey:BoardID;constraint:OnDelete:CASCADE" json:"lists,omitempty"`
	Labels []Label `gorm:"foreignKey:BoardID;constraint:OnDelete:CASCADE" json:"labels,omitempty"`
}
