package models

// PostActivityRequest represents the payload for an Activity
type PostActivityRequest struct {
	ID             int     `json:"id" db:"id"`
	Name           string  `json:"name" db:"name"`
	ActivityDate   string  `json:"activity_date" db:"activity_date"`
	ActivityTypeID int     `json:"activity_type_id" db:"activity_type_id"`
	OwnerID        int     `json:"owner_id" db:"owner_id"`
	Duration       float64 `json:"duration" db:"duration"`
	Distance       float64 `json:"distance" db:"distance"`
	Unit           string  `json:"unit" db:"unit"`
	HeartRate      int     `json:"heart_rate" db:"heart_rate"`
	CreatedAt      string  `json:"created_at" db:"created_at"`
	UpdatedAt      string  `json:"updated_at" db:"updated_at"`
}
