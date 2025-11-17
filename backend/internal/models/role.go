package models

import "time"

// Role represents a user role in the system (owner, admin, member, viewer)
type Role struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"unique;not null" json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// Permission represents a specific action that can be performed
type Permission struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"unique;not null" json:"name"`
	Resource    string    `gorm:"not null" json:"resource"`
	Action      string    `gorm:"not null" json:"action"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// RolePermission maps which permissions each role has
type RolePermission struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	RoleID       uint      `gorm:"not null" json:"role_id"`
	PermissionID uint      `gorm:"not null" json:"permission_id"`
	CreatedAt    time.Time `json:"created_at"`

	Role       Role       `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	Permission Permission `gorm:"foreignKey:PermissionID" json:"permission,omitempty"`
}

// BoardMember tracks which users have access to which boards and with what role
type BoardMember struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	BoardID    uint       `gorm:"not null" json:"board_id"`
	UserID     uint       `gorm:"not null" json:"user_id"`
	RoleID     uint       `gorm:"not null" json:"role_id"`
	InvitedBy  *uint      `json:"invited_by,omitempty"`
	InvitedAt  time.Time  `json:"invited_at"`
	AcceptedAt *time.Time `json:"accepted_at,omitempty"`
	Status     string     `gorm:"default:'active'" json:"status"` // active, pending, removed
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`

	Board Board `gorm:"foreignKey:BoardID" json:"board,omitempty"`
	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Role  Role  `gorm:"foreignKey:RoleID" json:"role,omitempty"`
}
