package models

//struct for user pref
type Preferences struct {
	ID            int    `json:"id,omitempty"`
	UserID        int    `json:"user_id,omitempty"`
	SortOrder     string `json:"sortOrder"`
	HiddenDevices []string `json:"hiddenDevices"`
	CustomIcons   string `json:"customIcons"`
}
