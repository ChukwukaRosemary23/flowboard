package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"unique;not null" json:"username"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"` // "-" means don't include in JSON response
	AvatarURL string         `json:"avatar_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete

	// Relationships
	Boards      []Board      `gorm:"foreignKey:OwnerID" json:"boards,omitempty"`
	CardMembers []CardMember `gorm:"foreignKey:UserID" json:"-"`
	Comments    []Comment    `gorm:"foreignKey:UserID" json:"-"`
	Attachments []Attachment `gorm:"foreignKey:UploadedBy" json:"-"`
}
