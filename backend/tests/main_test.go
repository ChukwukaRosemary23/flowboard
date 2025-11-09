package tests

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/ChukwukaRosemary23/flowboard-backend/config"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
)

var serverCmd *exec.Cmd

func TestMain(m *testing.M) {
	log.Println("🚀 Setting up test environment...")

	// Load test environment
	os.Setenv("ENV", "test")

	// Change to backend directory to find .env.test
	os.Chdir("..")

	cfg := config.LoadConfig()

	// Connect to test database
	log.Println("📊 Connecting to test database:", cfg.DBName)
	if err := database.ConnectDatabase(cfg); err != nil {
		log.Fatal("❌ Failed to connect to test database:", err)
	}

	// Auto-migrate tables
	log.Println("🔄 Running database migrations...")
	database.DB.AutoMigrate(
		&models.User{},
		&models.Board{},
		&models.List{},
		&models.Card{},
		&models.Role{},
		&models.Permission{},
		&models.RolePermission{},
		&models.BoardMember{},
	)

	// Seed roles and permissions if not already seeded
	seedRolesAndPermissions()

	// Start HTTP server as subprocess
	log.Println("🌐 Starting HTTP server on port 8083...")
	serverCmd = exec.Command("go", "run", "cmd/api/main.go")
	serverCmd.Env = append(os.Environ(), "ENV=test")

	if err := serverCmd.Start(); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}

	// Wait for server to be ready
	log.Println("⏳ Waiting for server to be ready...")
	time.Sleep(3 * time.Second)

	// Check if server is responding
	resp, err := http.Get("http://localhost:8083/ping")
	if err != nil || resp.StatusCode != 200 {
		log.Println("⚠️  Warning: Server may not be ready, but continuing...")
	} else {
		log.Println("✅ Server is ready!")
	}

	log.Println("🧪 Running tests...")

	// Run tests
	code := m.Run()

	// Cleanup
	log.Println("🧹 Cleaning up...")
	if serverCmd != nil && serverCmd.Process != nil {
		serverCmd.Process.Kill()
		log.Println("🛑 Server stopped")
	}

	log.Println("✅ Test environment cleaned up")
	os.Exit(code)
}

// seedRolesAndPermissions seeds the database with roles and permissions
func seedRolesAndPermissions() {
	// Check if roles already exist
	var count int64
	database.DB.Model(&models.Role{}).Count(&count)
	if count > 0 {
		log.Println("✅ Roles already seeded, skipping...")
		return
	}

	log.Println("🌱 Seeding roles and permissions...")

	// Create roles
	roles := []models.Role{
		{Name: "owner", Description: "Board creator with full control"},
		{Name: "admin", Description: "Can manage members and all board content"},
		{Name: "member", Description: "Can create and edit cards and lists"},
		{Name: "viewer", Description: "Read-only access to board"},
	}

	for _, role := range roles {
		database.DB.Create(&role)
	}

	// Create permissions
	permissions := []models.Permission{
		{Name: "view_board", Resource: "board", Action: "view", Description: "Can view board and its contents"},
		{Name: "edit_board", Resource: "board", Action: "edit", Description: "Can edit board settings"},
		{Name: "delete_board", Resource: "board", Action: "delete", Description: "Can delete the board"},
		{Name: "manage_members", Resource: "board", Action: "manage", Description: "Can add/remove members"},
		{Name: "create_list", Resource: "list", Action: "create", Description: "Can create new lists"},
		{Name: "edit_list", Resource: "list", Action: "edit", Description: "Can edit list properties"},
		{Name: "delete_list", Resource: "list", Action: "delete", Description: "Can delete lists"},
		{Name: "create_card", Resource: "card", Action: "create", Description: "Can create new cards"},
		{Name: "edit_card", Resource: "card", Action: "edit", Description: "Can edit card properties"},
		{Name: "delete_card", Resource: "card", Action: "delete", Description: "Can delete cards"},
		{Name: "move_card", Resource: "card", Action: "move", Description: "Can move cards between lists"},
	}

	for _, perm := range permissions {
		database.DB.Create(&perm)
	}

	// Assign permissions to roles
	var ownerRole, adminRole, memberRole, viewerRole models.Role
	database.DB.Where("name = ?", "owner").First(&ownerRole)
	database.DB.Where("name = ?", "admin").First(&adminRole)
	database.DB.Where("name = ?", "member").First(&memberRole)
	database.DB.Where("name = ?", "viewer").First(&viewerRole)

	var allPermissions []models.Permission
	database.DB.Find(&allPermissions)

	// Owner has ALL permissions
	for _, perm := range allPermissions {
		database.DB.Create(&models.RolePermission{
			RoleID:       ownerRole.ID,
			PermissionID: perm.ID,
		})
	}

	// Admin has all except delete_board
	for _, perm := range allPermissions {
		if perm.Name != "delete_board" {
			database.DB.Create(&models.RolePermission{
				RoleID:       adminRole.ID,
				PermissionID: perm.ID,
			})
		}
	}

	// Member can create/edit/delete lists and cards
	memberPermissions := []string{
		"view_board", "create_list", "edit_list", "delete_list",
		"create_card", "edit_card", "delete_card", "move_card",
	}
	for _, perm := range allPermissions {
		for _, allowed := range memberPermissions {
			if perm.Name == allowed {
				database.DB.Create(&models.RolePermission{
					RoleID:       memberRole.ID,
					PermissionID: perm.ID,
				})
			}
		}
	}

	// Viewer can only view
	for _, perm := range allPermissions {
		if perm.Name == "view_board" {
			database.DB.Create(&models.RolePermission{
				RoleID:       viewerRole.ID,
				PermissionID: perm.ID,
			})
		}
	}

	log.Println("✅ Roles and permissions seeded successfully")
}
