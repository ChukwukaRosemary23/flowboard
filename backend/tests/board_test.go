package tests

import (
	"fmt"
	"testing"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/stretchr/testify/suite"
)

type BoardTestSuite struct {
	suite.Suite
}

func (suite *BoardTestSuite) SetupTest() {
	// Setup before each test
}

func (suite *BoardTestSuite) TearDownTest() {
	// Cleanup after test
}

// Test creating a board
func (suite *BoardTestSuite) TestCreateBoard_Success() {
	user := Factory.CreateUser()
	token := GenerateTestJWT(user.ID, user.Username, user.Email)

	requestBody := map[string]string{
		"title":            "My New Board",
		"description":      "Test board description",
		"background_color": "#FF5733",
	}

	response := POST("/boards", requestBody, token)

	suite.Equal(201, response.StatusCode)
	suite.NotNil(response.Body["board"])
}

// Test creating board without authentication
func (suite *BoardTestSuite) TestCreateBoard_Unauthorized() {
	requestBody := map[string]string{
		"title": "My Board",
	}

	response := POST("/boards", requestBody) // No token

	suite.Equal(401, response.StatusCode)
}

// Test getting boards shows only accessible boards
func (suite *BoardTestSuite) TestGetBoards_ShowsOnlyAccessibleBoards() {
	user1 := Factory.CreateUser()
	user2 := Factory.CreateUser()

	Factory.CreateBoard(user1.ID)
	Factory.CreateBoard(user2.ID)

	token := GenerateTestJWT(user1.ID, user1.Username, user1.Email)
	response := GET("/boards", token)

	suite.Equal(200, response.StatusCode)

	boards := response.Body["boards"].([]interface{})
	suite.Equal(1, len(boards), "User1 should only see their own board")
}

// Test access control for deleting board
func (suite *BoardTestSuite) TestDeleteBoard_AccessControl() {
	owner := Factory.CreateUser()
	board := Factory.CreateBoard(owner.ID)

	testCases := []struct {
		role           string
		expectedStatus int
	}{
		{"owner", 200},
		{"admin", 403},
		{"member", 403},
		{"viewer", 403},
	}

	for _, tc := range testCases {
		suite.Run(tc.role, func() {
			var user *models.User

			if tc.role == "owner" {
				user = owner
			} else {
				user = Factory.CreateUser()
				Factory.CreateBoardMember(board.ID, user.ID, tc.role)
			}

			token := GenerateTestJWT(user.ID, user.Username, user.Email)
			response := DELETE(fmt.Sprintf("/boards/%d", board.ID), token)

			suite.Equal(tc.expectedStatus, response.StatusCode,
				fmt.Sprintf("%s should get %d status", tc.role, tc.expectedStatus))
		})
	}
}

// Test updating board access control
func (suite *BoardTestSuite) TestUpdateBoard_AccessControl() {
	owner := Factory.CreateUser()
	board := Factory.CreateBoard(owner.ID)

	testCases := []struct {
		role           string
		expectedStatus int
	}{
		{"owner", 200},
		{"admin", 200},
		{"member", 403},
		{"viewer", 403},
	}

	for _, tc := range testCases {
		suite.Run(tc.role, func() {
			var user *models.User

			if tc.role == "owner" {
				user = owner
			} else {
				user = Factory.CreateUser()
				Factory.CreateBoardMember(board.ID, user.ID, tc.role)
			}

			token := GenerateTestJWT(user.ID, user.Username, user.Email)

			requestBody := map[string]string{
				"title": "Updated Title",
			}

			response := PUT(fmt.Sprintf("/boards/%d", board.ID), requestBody, token)

			suite.Equal(tc.expectedStatus, response.StatusCode)
		})
	}
}

func TestBoardTestSuite(t *testing.T) {
	suite.Run(t, new(BoardTestSuite))
}
