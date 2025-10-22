package models

import (
	"time"

	"gorm.io/gorm"
)

// Card represents a task/card in a list
type Card struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	ListID      uint           `gorm:"not null" json:"list_id"`
	Position    int            `gorm:"not null;default:0" json:"position"` // Order within list
	DueDate     *time.Time     `json:"due_date,omitempty"`                 // Pointer = can be null
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	List        List         `gorm:"foreignKey:ListID" json:"list,omitempty"`
	Members     []User       `gorm:"many2many:card_members" json:"members,omitempty"`
	Labels      []Label      `gorm:"many2many:card_labels" json:"labels,omitempty"`
	Comments    []Comment    `gorm:"foreignKey:CardID;constraint:OnDelete:CASCADE" json:"comments,omitempty"`
	Attachments []Attachment `gorm:"foreignKey:CardID;constraint:OnDelete:CASCADE" json:"attachments,omitempty"`
}
