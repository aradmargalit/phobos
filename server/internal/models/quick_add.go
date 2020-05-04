package models

// QuickAdd represents a ready-to-add workout session
type QuickAdd struct {
	ID             int     `json:"id" db:"id"`
	Name           string  `json:"name" db:"name"`
	ActivityTypeID int     `json:"activity_type_id" db:"activity_type_id"`
	OwnerID        int     `json:"owner_id" db:"owner_id"`
	Duration       float64 `json:"duration" db:"duration"`
	Distance       float64 `json:"distance" db:"distance"`
	Unit           string  `json:"unit" db:"unit"`
	CreatedAt      string  `json:"created_at" db:"created_at"`
	UpdatedAt      string  `json:"updated_at" db:"updated_at"`
}
