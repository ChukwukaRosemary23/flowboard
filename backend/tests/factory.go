package tests

import (
	"os"
	"time"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/go-faker/faker/v4"
	"golang.org/x/crypto/bcrypt"
)

// FactoryHelper provides methods to create test data
type FactoryHelper struct{}

var Factory = &FactoryHelper{}

// CreateUser creates a user with random data
func (f *FactoryHelper) CreateUser() *models.User {
	
	testPassword := os.Getenv("TEST_USER_PASSWORD")
	if testPassword == "" {
		testPassword = "password123"
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(testPassword), bcrypt.DefaultCost)

	user := &models.User{
		Username: faker.Username(),
		Email:    faker.Email(),
		Password: string(hashedPassword),
	}

	database.DB.Create(user)
	return user
}

// CreateUserWithCredentials creates a user with specific credentials
func (f *FactoryHelper) CreateUserWithCredentials(email, password string) *models.User {

	if password == "" {
		panic("Password cannot be empty")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &models.User{
		Username: faker.Username(),
		Email:    email,
		Password: string(hashedPassword),
	}

	database.DB.Create(user)
	return user
}

// CreateBoard creates a board with an owner
func (f *FactoryHelper) CreateBoard(ownerID uint) *models.Board {
	board := &models.Board{
		Title:           faker.Word() + " Project",
		Description:     faker.Sentence(),
		BackgroundColor: "#0079BF",
		OwnerID:         ownerID,
	}

	database.DB.Create(board)

	var ownerRole models.Role
	database.DB.Where("name = ?", "owner").First(&ownerRole)

	boardMember := &models.BoardMember{
		BoardID:   board.ID,
		UserID:    ownerID,
		RoleID:    ownerRole.ID,
		InvitedAt: time.Now(),
		Status:    "active",
	}
	database.DB.Create(boardMember)

	return board
}

// CreateBoardMember adds a user to a board with a specific role
func (f *FactoryHelper) CreateBoardMember(boardID, userID uint, roleName string) *models.BoardMember {
	var role models.Role
	database.DB.Where("name = ?", roleName).First(&role)

	boardMember := &models.BoardMember{
		BoardID:   boardID,
		UserID:    userID,
		RoleID:    role.ID,
		InvitedAt: time.Now(),
		Status:    "active",
	}

	database.DB.Create(boardMember)
	return boardMember
}

// CreateList creates a list in a board
func (f *FactoryHelper) CreateList(boardID uint) *models.List {
	list := &models.List{
		Title:    faker.Word() + " List",
		BoardID:  boardID,
		Position: 0,
	}

	database.DB.Create(list)
	return list
}

// CreateCard creates a card in a list
func (f *FactoryHelper) CreateCard(listID uint) *models.Card {
	card := &models.Card{
		Title:       faker.Sentence(),
		Description: faker.Paragraph(),
		ListID:      listID,
		Position:    0,
	}

	database.DB.Create(card)
	return card
}
