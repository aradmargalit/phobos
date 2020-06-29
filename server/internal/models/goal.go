package models

// Goal represents a user's workout goal for any "metric" and "period"
// For example, 2 Miles per Month would be {goal: 2, metric: "miles", period: "Month"}
type Goal struct {
	ID        int     `json:"id" db:"id"`
	UserID    int     `json:"user_id" db:"user_id"`
	Metric    string  `json:"metric" db:"metric"`
	Period    string  `json:"period" db:"period"`
	Goal      float64 `json:"goal" db:"goal"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	UpdatedAt string  `json:"updated_at" db:"updated_at"`
}
