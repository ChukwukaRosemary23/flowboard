package utils

import (
	"encoding/json"

	"github.com/ChukwukaRosemary23/flowboard-backend/internal/database"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
)

// LogActivity creates an activity log entry
func LogActivity(action, entityType string, entityID, boardID, userID uint, entityTitle string, metadata map[string]interface{}) error {
	// Convert metadata to JSON
	metadataJSON := ""
	if metadata != nil {
		bytes, err := json.Marshal(metadata)
		if err == nil {
			metadataJSON = string(bytes)
		}
	}

	activity := models.Activity{
		Action:      action,
		EntityType:  entityType,
		EntityID:    entityID,
		EntityTitle: entityTitle,
		BoardID:     boardID,
		UserID:      userID,
		Metadata:    metadataJSON,
	}

	return database.DB.Create(&activity).Error
}
