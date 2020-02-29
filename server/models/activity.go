package models

// Activity represents a workout session
type Activity struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	ActivityDate string `json:"activity_date" db:"activity_date"`
	ActivityType string `json:"activity_type" db:"activity_type"`
	Duraton      int    `json:"duration" db:"duration"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}
