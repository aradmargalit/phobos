package responsetypes

import "database/sql"

// ActivityResponse represents a workout session
type ActivityResponse struct {
	ID             int                  `json:"id" db:"id"`
	Name           string               `json:"name" db:"name"`
	ActivityDate   string               `json:"activity_date" db:"activity_date"`
	ActivityTypeID int                  `json:"activity_type_id" db:"activity_type_id"`
	ActivityType   activityTypeResponse `json:"activity_type" db:"activity_type"`
	OwnerID        int                  `json:"owner_id" db:"owner_id"`
	Duration       float64              `json:"duration" db:"duration"`
	Distance       float64              `json:"distance" db:"distance"`
	Unit           string               `json:"unit" db:"unit"`
	StravaID       sql.NullInt64        `json:"strava_id" db:"strava_id"`
	LogicalIndex   int                  `json:"logical_index"`
	Epoch          int64                `json:"epoch"`
	CreatedAt      string               `json:"created_at" db:"created_at"`
	UpdatedAt      string               `json:"updated_at" db:"updated_at"`
}

type activityTypeResponse struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
