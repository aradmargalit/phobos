package models

// ActivityType represents a type for a workout
type ActivityType struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	StravaName string `json:"strava_name" db:"strava_name"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
}
