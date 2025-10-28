package models

import "time"

// CardMember represents the many-to-many relationship between cards and users
type CardMember struct {
	CardID     uint      `gorm:"primaryKey" json:"card_id"`
	UserID     uint      `gorm:"primaryKey" json:"user_id"`
	AssignedAt time.Time `gorm:"autoCreateTime" json:"assigned_at"`

	// Relationships
	Card Card `gorm:"foreignKey:CardID" json:"-"`
	User User `gorm:"foreignKey:UserID" json:"-"`
}
