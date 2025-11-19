package database

import (
	"log"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
)

// SeedRolesAndPermissions seeds the database with roles and permissions
func SeedRolesAndPermissions() {
	// Check if roles already exist
	var count int64
	DB.Model(&models.Role{}).Count(&count)
	if count > 0 {
		log.Println("✅ Roles already seeded, skipping...")
		return
	}

	log.Println("🌱 Seeding roles and permissions...")

	// Seed roles
	seedRoles()

	seedPermissions()

	assignPermissionsToRoles()

	log.Println(" Roles and permissions seeded successfully")
}

// seedRoles creates the four role types
func seedRoles() {
	roles := []models.Role{
		{Name: "owner", Description: "Board creator with full control"},
		{Name: "admin", Description: "Can manage members and all board content"},
		{Name: "member", Description: "Can create and edit cards and lists"},
		{Name: "viewer", Description: "Read-only access to board"},
	}

	for _, role := range roles {
		DB.Create(&role)
	}
	log.Println(" Roles created")
}

// seedPermissions creates all permission types
func seedPermissions() {
	permissions := []models.Permission{
		{Name: "view_board", Resource: "board", Action: "view", Description: "Can view board and its contents"},
		{Name: "edit_board", Resource: "board", Action: "edit", Description: "Can edit board settings"},
		{Name: "update_board", Resource: "board", Action: "update", Description: "Can update board properties"},
		{Name: "delete_board", Resource: "board", Action: "delete", Description: "Can delete the board"},
		{Name: "manage_members", Resource: "board", Action: "manage", Description: "Can add/remove members"},
		{Name: "invite_member", Resource: "board", Action: "invite", Description: "Can invite members to board"}, 
		{Name: "create_list", Resource: "list", Action: "create", Description: "Can create new lists"},
		{Name: "edit_list", Resource: "list", Action: "edit", Description: "Can edit list properties"},
		{Name: "delete_list", Resource: "list", Action: "delete", Description: "Can delete lists"},
		{Name: "create_card", Resource: "card", Action: "create", Description: "Can create new cards"},
		{Name: "edit_card", Resource: "card", Action: "edit", Description: "Can edit card properties"},
		{Name: "delete_card", Resource: "card", Action: "delete", Description: "Can delete cards"},
		{Name: "move_card", Resource: "card", Action: "move", Description: "Can move cards between lists"},
	}

	for _, perm := range permissions {
		DB.Create(&perm)
	}
	log.Println(" Permissions created")
}

// assignPermissionsToRoles assigns permissions to each role
func assignPermissionsToRoles() {

	var ownerRole, adminRole, memberRole, viewerRole models.Role
	DB.Where("name = ?", "owner").First(&ownerRole)
	DB.Where("name = ?", "admin").First(&adminRole)
	DB.Where("name = ?", "member").First(&memberRole)
	DB.Where("name = ?", "viewer").First(&viewerRole)

	// Get all permissions
	var allPermissions []models.Permission
	DB.Find(&allPermissions)

	// Owner: ALL permissions
	assignPermissionsToRole(ownerRole.ID, allPermissions, nil)

	// Admin: All except delete_board
	excludeAdmin := []string{"delete_board"}
	assignPermissionsToRole(adminRole.ID, allPermissions, excludeAdmin)

	// Member: Can create/edit/delete lists and cards
	memberAllowed := []string{
		"view_board", "create_list", "edit_list", "delete_list",
		"create_card", "edit_card", "delete_card", "move_card",
	}
	assignSpecificPermissions(memberRole.ID, allPermissions, memberAllowed)

	// Viewer: Only view
	viewerAllowed := []string{"view_board"}
	assignSpecificPermissions(viewerRole.ID, allPermissions, viewerAllowed)

	log.Println(" Permissions assigned to roles")
}

// assignPermissionsToRole assigns permissions to a role (with optional exclusions)
func assignPermissionsToRole(roleID uint, permissions []models.Permission, exclude []string) {
	for _, perm := range permissions {

		shouldExclude := false
		for _, excluded := range exclude {
			if perm.Name == excluded {
				shouldExclude = true
				break
			}
		}

		if !shouldExclude {
			DB.Create(&models.RolePermission{
				RoleID:       roleID,
				PermissionID: perm.ID,
			})
		}
	}
}

// assignSpecificPermissions assigns only specific permissions to a role
func assignSpecificPermissions(roleID uint, allPermissions []models.Permission, allowed []string) {
	for _, perm := range allPermissions {
		for _, allowedName := range allowed {
			if perm.Name == allowedName {
				DB.Create(&models.RolePermission{
					RoleID:       roleID,
					PermissionID: perm.ID,
				})
				break
			}
		}
	}
}
