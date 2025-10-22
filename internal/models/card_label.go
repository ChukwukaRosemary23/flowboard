package models

// CardLabel represents the many-to-many relationship between cards and labels
type CardLabel struct {
	CardID  uint `gorm:"primaryKey" json:"card_id"`
	LabelID uint `gorm:"primaryKey" json:"label_id"`

	// Relationships
	Card  Card  `gorm:"foreignKey:CardID" json:"-"`
	Label Label `gorm:"foreignKey:LabelID" json:"-"`
}
