package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/stretchr/testify/suite"
)

type BoardMemberTestSuite struct {
	suite.Suite
}

func (suite *BoardMemberTestSuite) TearDownTest() {

}

// Test inviting a member successfully
func (suite *BoardMemberTestSuite) TestInviteMember_Success() {
	owner := Factory.CreateUser()
	board := Factory.CreateBoard(owner.ID)
	newMember := Factory.CreateUser()

	log.Printf("\n🔍 === DATABASE STATE CHECK ===")

	// Check if owner is in board_members
	var boardMemberCount int64
	database.DB.Table("board_members").
		Where("user_id = ? AND board_id = ?", owner.ID, board.ID).
		Count(&boardMemberCount)
	log.Printf("Owner in board_members: %d", boardMemberCount)

	// Get owner's role
	var ownerBoardMember models.BoardMember
	database.DB.Preload("Role").
		Where("user_id = ? AND board_id = ?", owner.ID, board.ID).
		First(&ownerBoardMember)
	log.Printf("Owner's role: %s (ID: %d)", ownerBoardMember.Role.Name, ownerBoardMember.RoleID)

	// Check if owner's role has invite_member permission
	var permCount int64
	database.DB.Table("role_permissions").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id = ? AND permissions.name = ?", ownerBoardMember.RoleID, "invite_member").
		Count(&permCount)
	log.Printf("Owner's role has 'invite_member' permission: %d", permCount)

	// List all permissions for owner's role
	var permissions []string
	database.DB.Table("role_permissions").
		Select("permissions.name").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id = ?", ownerBoardMember.RoleID).
		Scan(&permissions)
	log.Printf("All permissions for role '%s': %v", ownerBoardMember.Role.Name, permissions)
	log.Printf("=================================\n")

	token := GenerateTestJWT(owner.ID, owner.Username, owner.Email)

	requestBody := map[string]interface{}{
		"user_id": newMember.ID,
		"role":    "member",
	}

	response := POST(fmt.Sprintf("/boards/%d/members", board.ID), requestBody, token)
	LogResponse("TestInviteMember_Success", response)

	suite.Equal(201, response.StatusCode, "Should return 201 Created")
	suite.Equal("Member added successfully", response.Body["message"])

	// Only check member if response succeeded
	if response.StatusCode == 201 && response.Body["member"] != nil {
		member := response.Body["member"].(map[string]interface{})
		suite.Equal(float64(board.ID), member["board_id"])
		suite.Equal(float64(newMember.ID), member["user_id"])
		suite.Equal(float64(owner.ID), member["invited_by"])
		suite.NotNil(member["role_id"])
	}
}

// Test duplicate member invitation
func (suite *BoardMemberTestSuite) TestInviteMember_DuplicateMember() {
	owner := Factory.CreateUser()
	board := Factory.CreateBoard(owner.ID)
	member := Factory.CreateUser()

	// Add member first time
	Factory.CreateBoardMember(board.ID, member.ID, "member")

	token := GenerateTestJWT(owner.ID, owner.Username, owner.Email)

	requestBody := map[string]interface{}{
		"user_id": member.ID,
		"role":    "member",
	}

	// Try to add same member again
	response := POST(fmt.Sprintf("/boards/%d/members", board.ID), requestBody, token)
	LogResponse("TestInviteMember_DuplicateMember", response)

	suite.Equal(400, response.StatusCode)
	suite.Contains(response.Body["error"], "already a member")
}

// Test access control for inviting members
func (suite *BoardMemberTestSuite) TestInviteMember_AccessControl() {
	testCases := []struct {
		role           string
		expectedStatus int
	}{
		{"owner", 201},
		{"admin", 201},
		{"member", 403},
		{"viewer", 403},
	}

	for _, tc := range testCases {
		suite.Run(tc.role, func() {

			owner := Factory.CreateUser()
			board := Factory.CreateBoard(owner.ID)
			newMember := Factory.CreateUser()

			var user *models.User

			if tc.role == "owner" {
				user = owner
			} else {
				user = Factory.CreateUser()
				Factory.CreateBoardMember(board.ID, user.ID, tc.role)
			}

			token := GenerateTestJWT(user.ID, user.Username, user.Email)

			requestBody := map[string]interface{}{
				"user_id": newMember.ID,
				"role":    "member",
			}

			response := POST(fmt.Sprintf("/boards/%d/members", board.ID), requestBody, token)
			LogResponse(fmt.Sprintf("TestInviteMember_AccessControl/%s", tc.role), response)

			suite.Equal(tc.expectedStatus, response.StatusCode,
				fmt.Sprintf("%s should get %d status", tc.role, tc.expectedStatus))

			// Assertions for successful invitations
			if tc.expectedStatus == 201 {
				suite.Equal("Member added successfully", response.Body["message"])
				suite.NotNil(response.Body["member"], "Should return member object")
			}

			// Assertions for denied access
			if tc.expectedStatus == 403 {
				suite.Contains(response.Body["error"], "permission")
			}
		})
	}
}

// Test removing a member successfully
func (suite *BoardMemberTestSuite) TestRemoveMember_Success() {
	owner := Factory.CreateUser()
	board := Factory.CreateBoard(owner.ID)
	member := Factory.CreateUser()

	boardMember := Factory.CreateBoardMember(board.ID, member.ID, "member")

	token := GenerateTestJWT(owner.ID, owner.Username, owner.Email)

	response := DELETE(fmt.Sprintf("/boards/%d/members/%d", board.ID, boardMember.ID), token)
	LogResponse("TestRemoveMember_Success", response)

	suite.Equal(200, response.StatusCode)
	suite.Equal("Member removed successfully", response.Body["message"])

}

// Test that owner cannot be removed
func (suite *BoardMemberTestSuite) TestRemoveMember_CannotRemoveOwner() {
	owner := Factory.CreateUser()
	board := Factory.CreateBoard(owner.ID)

	// Owner is automatically added as a board member
	var ownerMember models.BoardMember
	database.DB.Where("board_id = ? AND user_id = ?", board.ID, owner.ID).First(&ownerMember)

	token := GenerateTestJWT(owner.ID, owner.Username, owner.Email)

	response := DELETE(fmt.Sprintf("/boards/%d/members/%d", board.ID, ownerMember.ID), token)
	LogResponse("TestRemoveMember_CannotRemoveOwner", response)

	suite.Equal(400, response.StatusCode)
	suite.Contains(response.Body["error"], "owner")
}

func TestBoardMemberTestSuite(t *testing.T) {
	suite.Run(t, new(BoardMemberTestSuite))
}
