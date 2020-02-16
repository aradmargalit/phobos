package models

// User represents the user the comes back from the Google Response
type User struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	GivenName string `json:"given_name" db:"given_name"`
	Email     string `json:"email" db:"email"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
