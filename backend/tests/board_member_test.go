package tests

import (
	"fmt"
	"testing"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/stretchr/testify/suite"
)

type BoardMemberTestSuite struct {
	suite.Suite
}

func (suite *BoardMemberTestSuite) SetupTest() {
	// Setup
}

// Test inviting a member to board
func (suite *BoardMemberTestSuite) TestInviteMember_Success() {
	owner := Factory.CreateUser()
	board := Factory.CreateBoard(owner.ID)
	newMember := Factory.CreateUser()

	token := GenerateTestJWT(owner.ID, owner.Username, owner.Email)
	requestBody := map[string]interface{}{
		"user_id": newMember.ID,
		"role":    "member",
	}

	response := POST(fmt.Sprintf("/boards/%d/members", board.ID), requestBody, token)

	suite.Equal(201, response.StatusCode)
	suite.Equal("Member added successfully", response.Body["message"])
}

// Test inviting duplicate member
func (suite *BoardMemberTestSuite) TestInviteMember_DuplicateMember() {
	owner := Factory.CreateUser()
	board := Factory.CreateBoard(owner.ID)
	member := Factory.CreateUser()

	Factory.CreateBoardMember(board.ID, member.ID, "member")

	token := GenerateTestJWT(owner.ID, owner.Username, owner.Email)
	requestBody := map[string]interface{}{
		"user_id": member.ID,
		"role":    "member",
	}

	response := POST(fmt.Sprintf("/boards/%d/members", board.ID), requestBody, token)

	suite.Equal(400, response.StatusCode)
}

// Test only admin/owner can invite members
func (suite *BoardMemberTestSuite) TestInviteMember_AccessControl() {
	owner := Factory.CreateUser()
	board := Factory.CreateBoard(owner.ID)
	newUser := Factory.CreateUser()

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
			var inviter *models.User

			if tc.role == "owner" {
				inviter = owner
			} else {
				inviter = Factory.CreateUser()
				Factory.CreateBoardMember(board.ID, inviter.ID, tc.role)
			}

			token := GenerateTestJWT(inviter.ID, inviter.Username, inviter.Email)
			requestBody := map[string]interface{}{
				"user_id": newUser.ID,
				"role":    "member",
			}

			response := POST(fmt.Sprintf("/boards/%d/members", board.ID), requestBody, token)

			suite.Equal(tc.expectedStatus, response.StatusCode)
		})
	}
}

// Test removing member
func (suite *BoardMemberTestSuite) TestRemoveMember_Success() {
	owner := Factory.CreateUser()
	board := Factory.CreateBoard(owner.ID)
	member := Factory.CreateUser()
	Factory.CreateBoardMember(board.ID, member.ID, "member")

	token := GenerateTestJWT(owner.ID, owner.Username, owner.Email)
	response := DELETE(fmt.Sprintf("/boards/%d/members/%d", board.ID, member.ID), token)

	suite.Equal(200, response.StatusCode)
}

// Test cannot remove owner
func (suite *BoardMemberTestSuite) TestRemoveMember_CannotRemoveOwner() {
	owner := Factory.CreateUser()
	board := Factory.CreateBoard(owner.ID)
	admin := Factory.CreateUser()
	Factory.CreateBoardMember(board.ID, admin.ID, "admin")

	token := GenerateTestJWT(admin.ID, admin.Username, admin.Email)
	response := DELETE(fmt.Sprintf("/boards/%d/members/%d", board.ID, owner.ID), token)

	suite.Equal(403, response.StatusCode)
}

func TestBoardMemberTestSuite(t *testing.T) {
	suite.Run(t, new(BoardMemberTestSuite))
}
