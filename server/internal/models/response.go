package models

// ActivityResponse represents the Activity as the client sees it
type ActivityResponse struct {
	Activity
	ActivityType `json:"activity_type" db:"activity_type"`
	LogicalIndex int   `json:"logical_index"`
	Epoch        int64 `json:"epoch"`
}
