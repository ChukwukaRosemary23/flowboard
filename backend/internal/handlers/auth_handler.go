package handlers

import (
	"net/http"

	"github.com/ChukwukaRosemary23/flowboard-backend/config"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

// Register handles user registration
func Register(c *gin.Context) {
	var req RegisterRequest

	// Validate input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if username already exists
	var existingUser models.User
	if err := database.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	// Check if email already exists
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate JWT token
	cfg := config.LoadConfig()
	token, err := utils.GenerateJWT(user.ID, user.Username, user.Email, cfg.JWTSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return response
	c.JSON(http.StatusCreated, gin.H{
		"token": token,
		"user": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"avatar_url": user.AvatarURL,
		},
	})
}

// Login handles user login
func Login(c *gin.Context) {
	var req LoginRequest

	// Validate input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check password
	if !utils.CheckPassword(req.Password, user.Password) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	cfg := config.LoadConfig()
	token, err := utils.GenerateJWT(user.ID, user.Username, user.Email, cfg.JWTSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"avatar_url": user.AvatarURL,
		},
	})
}
