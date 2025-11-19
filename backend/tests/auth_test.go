package tests

import (
	"testing"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite
}

func (suite *AuthTestSuite) SetupTest() {

	database.DB.Exec("DELETE FROM users")
}

func (suite *AuthTestSuite) TearDownTest() {

}

// Test user registration with valid data
func (suite *AuthTestSuite) TestRegister_Success() {
	requestBody := map[string]string{
		"username": "newuser",
		"email":    "newuser@test.com",
		"password": "password123",
	}

	response := POST("/auth/register", requestBody)

	suite.Equal(201, response.StatusCode, "Should return 201 Created")
	suite.NotNil(response.Body["user"], "Response should contain user")
	suite.NotNil(response.Body["token"], "Response should contain JWT token")
}

// Test registration with duplicate email
func (suite *AuthTestSuite) TestRegister_DuplicateEmail() {

	Factory.CreateUserWithCredentials("existing@test.com", "password123")

	requestBody := map[string]string{
		"username": "newuser",
		"email":    "existing@test.com",
		"password": "password123",
	}

	response := POST("/auth/register", requestBody)

	suite.True(response.StatusCode == 400 || response.StatusCode == 409)
	suite.NotNil(response.Body["error"])
}

// Test login with valid credentials
func (suite *AuthTestSuite) TestLogin_Success() {
	Factory.CreateUserWithCredentials("user@test.com", "password123")

	requestBody := map[string]string{
		"email":    "user@test.com",
		"password": "password123",
	}

	response := POST("/auth/login", requestBody)

	suite.Equal(200, response.StatusCode)
	suite.NotNil(response.Body["token"])
}

// Test login with invalid credentials
func (suite *AuthTestSuite) TestLogin_InvalidCredentials() {
	Factory.CreateUserWithCredentials("user@test.com", "password123")

	requestBody := map[string]string{
		"email":    "user@test.com",
		"password": "wrongpassword",
	}

	response := POST("/auth/login", requestBody)

	suite.Equal(401, response.StatusCode)
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
