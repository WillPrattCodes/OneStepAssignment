package services

import (
	"backend/internal/models"
	"database/sql"
	"encoding/json"
	"fmt"
)

// Service to fetch user preferences
func FetchPreferences(db *sql.DB, userID int) (models.Preferences, error) {
	//create preferences object
	var preferences models.Preferences
	var hiddenDevicesJSON string // Temporary variable to hold the JSON string

	//define query to fetch preferences
	query := `SELECT id, user_id, sort_order, hidden_devices, custom_icons 
			  FROM preferences WHERE user_id = ?`
	//execute query
	err := db.QueryRow(query, userID).Scan(
		&preferences.ID,
		&preferences.UserID,
		&preferences.SortOrder,
		&hiddenDevicesJSON, // Scan hidden_devices as a JSON string
		&preferences.CustomIcons,
	)

	// Error handling
	if err != nil {
		if err == sql.ErrNoRows {
			return preferences, nil // Return empty preferences if not found
		}
		return preferences, fmt.Errorf("failed to fetch preferences: %v", err)
	}

	// Deserialize hiddenDevices from JSON string
	if hiddenDevicesJSON != "" {
		err = json.Unmarshal([]byte(hiddenDevicesJSON), &preferences.HiddenDevices)
		if err != nil {
			return preferences, fmt.Errorf("failed to unmarshal hidden devices: %v", err)
		}
	}

	//return preferences
	return preferences, nil
}

// Service to save or update user preferences
func SaveOrUpdatePreferences(db *sql.DB, userID int, prefs models.Preferences) error {
	// Convert hiddenDevices to a JSON string
	hiddenDevicesJSON, err := json.Marshal(prefs.HiddenDevices)
	if err != nil {
		return fmt.Errorf("failed to marshal hidden devices: %v", err)
	}

	// Define query
	query := `INSERT INTO preferences (user_id, sort_order, hidden_devices, custom_icons)
			  VALUES (?, ?, ?, ?)
			  ON CONFLICT(user_id) DO UPDATE SET 
			  sort_order = ?, hidden_devices = ?, custom_icons = ?`
	// Execute query
	_, err = db.Exec(query,
		userID, prefs.SortOrder, hiddenDevicesJSON, prefs.CustomIcons,
		prefs.SortOrder, hiddenDevicesJSON, prefs.CustomIcons,
	)
	// Error handling
	if err != nil {
		return fmt.Errorf("failed to save/update preferences: %v", err)
	}
	return nil
}
