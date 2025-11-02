package services

import (
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
)

// PermissionService handles all permission-checking logic
type PermissionService struct{}

// CheckPermission verifies if a user has a specific permission on a board
func (ps *PermissionService) CheckPermission(userID, boardID uint, permissionName string) bool {
	var count int64

	database.DB.Table("board_members").
		Joins("JOIN roles ON board_members.role_id = roles.id").
		Joins("JOIN role_permissions ON roles.id = role_permissions.role_id").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("board_members.user_id = ?", userID).
		Where("board_members.board_id = ?", boardID).
		Where("board_members.status = ?", "active").
		Where("permissions.name = ?", permissionName).
		Count(&count)

	return count > 0
}

// GetUserRole returns the user's role on a specific board
func (ps *PermissionService) GetUserRole(userID, boardID uint) (string, error) {
	var boardMember models.BoardMember

	err := database.DB.Preload("Role").
		Where("user_id = ? AND board_id = ? AND status = ?", userID, boardID, "active").
		First(&boardMember).Error

	if err != nil {
		return "", err
	}

	return boardMember.Role.Name, nil
}

// IsOwner checks if user is the owner of the board
func (ps *PermissionService) IsOwner(userID, boardID uint) bool {
	role, err := ps.GetUserRole(userID, boardID)
	return err == nil && role == "owner"
}

// IsAdmin checks if user is admin or owner
func (ps *PermissionService) IsAdmin(userID, boardID uint) bool {
	role, err := ps.GetUserRole(userID, boardID)
	return err == nil && (role == "owner" || role == "admin")
}

// HasBoardAccess checks if a user has any access to a board
func (ps *PermissionService) HasBoardAccess(userID, boardID uint) bool {
	var count int64

	database.DB.Model(&models.BoardMember{}).
		Where("user_id = ? AND board_id = ? AND status = ?", userID, boardID, "active").
		Count(&count)

	return count > 0
}
