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

	// CLEAN UP OLD DATA BEFORE STARTING - Comprehensive cleanup
	log.Println("🧹 Cleaning up old test data...")
	database.DB.Exec("TRUNCATE TABLE activities CASCADE")
	database.DB.Exec("TRUNCATE TABLE attachments CASCADE")
	database.DB.Exec("TRUNCATE TABLE card_members CASCADE")
	database.DB.Exec("TRUNCATE TABLE card_labels CASCADE")
	database.DB.Exec("TRUNCATE TABLE comments CASCADE")
	database.DB.Exec("TRUNCATE TABLE cards CASCADE")
	database.DB.Exec("TRUNCATE TABLE labels CASCADE")
	database.DB.Exec("TRUNCATE TABLE lists CASCADE")
	database.DB.Exec("TRUNCATE TABLE board_members CASCADE")
	database.DB.Exec("TRUNCATE TABLE boards CASCADE")
	database.DB.Exec("TRUNCATE TABLE role_permissions CASCADE")
	database.DB.Exec("TRUNCATE TABLE permissions CASCADE")
	database.DB.Exec("TRUNCATE TABLE roles CASCADE")
	database.DB.Exec("TRUNCATE TABLE users CASCADE")
	log.Println("✅ Old test data cleaned")

	// Auto-migrate tables (ALL models in correct order)
	log.Println("🔄 Running database migrations...")
	database.DB.AutoMigrate(
		&models.User{},
		&models.Board{},
		&models.List{},
		&models.Card{},
		&models.Comment{},
		&models.Label{},
		&models.CardLabel{},
		&models.CardMember{},
		&models.Attachment{},
		&models.Activity{},
		&models.Role{},
		&models.Permission{},
		&models.RolePermission{},
		&models.BoardMember{},
	)

	// Seed roles and permissions using shared function
	database.SeedRolesAndPermissions()

	// Start HTTP server as subprocess
	log.Println("🌐 Starting HTTP server...")

	// Capture server output to file for debugging
	serverLogFile, err := os.Create("server_test.log")
	if err != nil {
		log.Fatal("Failed to create log file:", err)
	}
	defer func() {
		time.Sleep(100 * time.Millisecond) // Give time for writes
		serverLogFile.Close()
	}()

	serverCmd = exec.Command("go", "run", "cmd/api/main.go")
	serverCmd.Env = append(os.Environ(),
		"ENV=test",
		"DB_NAME=flowboard_tests3",
		"DB_HOST=localhost",
		"DB_PORT=5432",
		"DB_USER=postgres",
		"DB_PASSWORD=Rose1234",
		"PORT=8083",
		"JWT_SECRET=68aea209f5a75004f288d289973933808d5adfd8184fb767ad3",
	)
	serverCmd.Stdout = serverLogFile
	serverCmd.Stderr = serverLogFile

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
		&models.Activity{},
		&models.Attachment{},
		&models.CardMember{},
		&models.CardLabel{},
		&models.Label{},
		&models.Comment{},
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
