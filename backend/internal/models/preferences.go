package models

//struct for user pref
type Preferences struct {
	ID            int    `json:"id,omitempty"`
	UserID        int    `json:"user_id,omitempty"`
	SortOrder     string `json:"sort_order"`
	HiddenDevices string `json:"hidden_devices"`
	CustomIcons   string `json:"custom_icons"`
}
