package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

// UploadAttachment uploads a file to a card
func UploadAttachment(c *gin.Context) {
	cardID := c.Param("card_id")
	userID := c.GetUint("user_id")

	// Verify card exists and user has access
	var card models.Card
	if err := database.DB.Preload("List.Board").First(&card, cardID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	// Verify user owns the board
	if card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Get uploaded file
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	// Validate file size (max 10MB)
	if header.Size > 10*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large (max 10MB)"})
		return
	}

	
	uploadsDir := "./uploads"
	if err := os.MkdirAll(uploadsDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create uploads directory"})
		return
	}

	// Generate unique filename
	timestamp := time.Now().Unix()
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d_%s%s", timestamp, cardID, ext)
	filepath := filepath.Join(uploadsDir, filename)

	// Save file to disk
	dst, err := os.Create(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Create attachment record
	attachment := models.Attachment{
		Filename:   header.Filename,
		FileURL:    "/uploads/" + filename,
		FileSize:   header.Size,
		FileType:   header.Header.Get("Content-Type"),
		CardID:     card.ID,
		UploadedBy: userID,
	}

	if err := database.DB.Create(&attachment).Error; err != nil {
		
		os.Remove(filepath)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create attachment record"})
		return
	}

	// Log activity
	utils.LogActivity("uploaded_file", "attachment", attachment.ID, card.List.Board.ID, userID, header.Filename, nil)

	// Load uploader info for response
	database.DB.Preload("Uploader").First(&attachment, attachment.ID)

	c.JSON(http.StatusCreated, AttachmentResponse{
		ID:       attachment.ID,
		Filename: attachment.Filename,
		FileURL:  attachment.FileURL,
		FileSize: attachment.FileSize,
		FileType: attachment.FileType,
		CardID:   attachment.CardID,
		UploadedBy: UserResponse{
			ID:        attachment.Uploader.ID,
			Username:  attachment.Uploader.Username,
			Email:     attachment.Uploader.Email,
			AvatarURL: attachment.Uploader.AvatarURL,
		},
		CreatedAt: attachment.CreatedAt,
	})
}

// GetAttachments returns all attachments for a card
func GetAttachments(c *gin.Context) {
	cardID := c.Param("card_id")
	userID := c.GetUint("user_id")

	// Verify card exists and user has access
	var card models.Card
	if err := database.DB.Preload("List.Board").First(&card, cardID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	if card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Get all attachments
	var attachments []models.Attachment
	if err := database.DB.Where("card_id = ?", cardID).
		Preload("Uploader").
		Order("created_at DESC").
		Find(&attachments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attachments"})
		return
	}

	// Convert to response
	response := make([]AttachmentResponse, len(attachments))
	for i, attachment := range attachments {
		response[i] = AttachmentResponse{
			ID:       attachment.ID,
			Filename: attachment.Filename,
			FileURL:  attachment.FileURL,
			FileSize: attachment.FileSize,
			FileType: attachment.FileType,
			CardID:   attachment.CardID,
			UploadedBy: UserResponse{
				ID:        attachment.Uploader.ID,
				Username:  attachment.Uploader.Username,
				Email:     attachment.Uploader.Email,
				AvatarURL: attachment.Uploader.AvatarURL,
			},
			CreatedAt: attachment.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"attachments": response,
		"count":       len(response),
	})
}

// DownloadAttachment serves a file for download
func DownloadAttachment(c *gin.Context) {
	attachmentID := c.Param("id")
	userID := c.GetUint("user_id")

	// Find attachment
	var attachment models.Attachment
	if err := database.DB.Preload("Card.List.Board").First(&attachment, attachmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attachment not found"})
		return
	}

	// Verify user owns the board
	if attachment.Card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Get file path
	filePath := "." + attachment.FileURL

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found on server"})
		return
	}

	// Serve file
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+attachment.Filename)
	c.Header("Content-Type", attachment.FileType)
	c.File(filePath)
}

// DeleteAttachment deletes an attachment
func DeleteAttachment(c *gin.Context) {
	attachmentID := c.Param("id")
	userID := c.GetUint("user_id")

	// Find attachment
	var attachment models.Attachment
	if err := database.DB.Preload("Card.List.Board").First(&attachment, attachmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attachment not found"})
		return
	}

	// Verify user owns the board
	if attachment.Card.List.Board.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Delete file from disk
	filePath := "." + attachment.FileURL
	if err := os.Remove(filePath); err != nil {
		
		fmt.Println("Warning: Could not delete file:", err)
	}

	// Delete database record
	if err := database.DB.Delete(&attachment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete attachment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Attachment deleted successfully",
		"id":      attachmentID,
	})
}
