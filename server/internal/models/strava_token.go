package models

// StravaToken represents a ready-to-add workout session
type StravaToken struct {
	UserID       int    `json:"user_id" db:"user_id"`
	StravaID     int    `json:"strava_id" db:"strava_id"`
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	Expiry       string `json:"expiry" db:"expiry"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}
