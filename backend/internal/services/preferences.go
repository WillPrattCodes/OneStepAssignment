package services

import (
	"backend/internal/models"
	"database/sql"
	"fmt"
)

//service to fetch user preferences
func FetchPreferences(db *sql.DB, userID int) (models.Preferences, error) {
	//create preferences object
	var preferences models.Preferences
	//define query to fetch preferences
	query := `SELECT id, user_id, sort_order, hidden_devices, custom_icons 
			  FROM preferences WHERE user_id = ?`
	//execute query
	err := db.QueryRow(query, userID).Scan(
		&preferences.ID,
		&preferences.UserID,
		&preferences.SortOrder,
		&preferences.HiddenDevices,
		&preferences.CustomIcons,
	)

	//error handling
	if err != nil {
		if err == sql.ErrNoRows {
			return preferences, nil // Return empty preferences if not found
		}
		return preferences, fmt.Errorf("failed to fetch preferences: %v", err)
	}
	//return preferences
	return preferences, nil
}

//service to save or update user preferences
func SaveOrUpdatePreferences(db *sql.DB, userID int, prefs models.Preferences) error {
	//define query
	query := `INSERT INTO preferences (user_id, sort_order, hidden_devices, custom_icons)
			  VALUES (?, ?, ?, ?)
			  ON CONFLICT(user_id) DO UPDATE SET 
			  sort_order = ?, hidden_devices = ?, custom_icons = ?`
	//execute query
	_, err := db.Exec(query,
		userID, prefs.SortOrder, prefs.HiddenDevices, prefs.CustomIcons,
		prefs.SortOrder, prefs.HiddenDevices, prefs.CustomIcons,
	)
	//error handling
	if err != nil {
		return fmt.Errorf("failed to save/update preferences: %v", err)
	}
	return nil
}
