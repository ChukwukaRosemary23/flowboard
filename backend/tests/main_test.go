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
	"github.com/joho/godotenv"
)

var serverCmd *exec.Cmd

func TestMain(m *testing.M) {
	log.Println("🚀 Setting up test environment...")

	// Load test environment
	os.Setenv("ENV", "test")

	// Change to backend directory to find .env.test
	os.Chdir("..")

	// Force load .env.test file
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatal("❌ Error loading .env.test file:", err)
	}

	cfg := config.LoadConfig()

	// Connect to test database
	log.Println("📊 Connecting to test database:", cfg.DBName)
	if err := database.ConnectDatabase(cfg); err != nil {
		log.Fatal("❌ Failed to connect to test database:", err)
	}

	// Auto-migrate tables (ALL models in correct order)
	log.Println("🔄 Running database migrations...")
	database.DB.AutoMigrate(
		&models.User{},
		&models.Board{},
		&models.List{},
		&models.Card{},
		&models.Comment{},    // ← ADDED
		&models.Label{},      // ← ADDED
		&models.CardLabel{},  // ← ADDED
		&models.CardMember{}, // ← ADDED
		&models.Attachment{}, // ← ADDED
		&models.Activity{},   // ← ADDED
		&models.Role{},
		&models.Permission{},
		&models.RolePermission{},
		&models.BoardMember{},
	)

	// Seed roles and permissions using shared function
	database.SeedRolesAndPermissions()

	// Start HTTP server as subprocess
	log.Println("🌐 Starting HTTP server...")
	serverCmd = exec.Command("go", "run", "cmd/api/main.go")
	serverCmd.Env = append(os.Environ(), "ENV=test")

	if err := serverCmd.Start(); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}

	// Wait for server to be ready with retry logic
	log.Println("⏳ Waiting for server to be ready...")
	if !waitForServer(cfg.Port, 5, 3*time.Second) {
		log.Fatal("❌ Server failed to start after 5 retries")
	}

	log.Println("✅ Server is ready!")
	log.Println("🧪 Running tests...")

	// Run tests
	code := m.Run()

	// Cleanup
	log.Println("🧹 Cleaning up...")
	if serverCmd != nil && serverCmd.Process != nil {
		serverCmd.Process.Kill()
		log.Println("🛑 Server stopped")
	}

	// Rollback migrations (drop tables in reverse order)
	log.Println("🔄 Rolling back migrations...")
	database.DB.Migrator().DropTable(
		&models.BoardMember{},
		&models.RolePermission{},
		&models.Permission{},
		&models.Role{},
		&models.Activity{},   // ← ADDED
		&models.Attachment{}, // ← ADDED
		&models.CardMember{}, // ← ADDED
		&models.CardLabel{},  // ← ADDED
		&models.Label{},      // ← ADDED
		&models.Comment{},    // ← ADDED
		&models.Card{},
		&models.List{},
		&models.Board{},
		&models.User{},
	)

	log.Println("✅ Test environment cleaned up")
	os.Exit(code)
}

// waitForServer checks if server is ready with retry logic
func waitForServer(port string, maxRetries int, waitTime time.Duration) bool {
	url := "http://localhost:" + port + "/health"

	for i := 0; i < maxRetries; i++ {
		log.Printf("⏳ Checking server health (attempt %d/%d)...", i+1, maxRetries)

		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			return true
		}

		if i < maxRetries-1 {
			log.Printf("⏳ Server not ready, waiting %v before retry...", waitTime)
			time.Sleep(waitTime)
		}
	}

	return false
}
