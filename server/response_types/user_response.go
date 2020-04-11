package responsetypes

// UserResponse represents the database user
type UserResponse struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	GivenName   string `json:"given_name" db:"given_name"`
	Email       string `json:"email" db:"email"`
	StravaToken bool   `json:"strava_token" db:"strava_token"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}
