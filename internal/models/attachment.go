package models

import (
	"time"

	"gorm.io/gorm"
)

// Attachment represents a file attached to a card
type Attachment struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Filename   string         `gorm:"not null" json:"filename"`
	FileURL    string         `gorm:"not null" json:"file_url"` // Path where file is stored
	FileSize   int64          `json:"file_size"`                // Size in bytes
	FileType   string         `json:"file_type"`                // MIME type (image/png, application/pdf, etc.)
	CardID     uint           `gorm:"not null" json:"card_id"`
	UploadedBy uint           `gorm:"not null" json:"uploaded_by"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Card     Card `gorm:"foreignKey:CardID" json:"card,omitempty"`
	Uploader User `gorm:"foreignKey:UploadedBy" json:"uploader,omitempty"`
}
