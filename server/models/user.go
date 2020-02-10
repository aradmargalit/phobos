package models

// User represents the user the comes back from the Google Response
type User struct {
	Name      string `json:"name"`
	GivenName string `json:"given_name"`
	Email     string `json:"email"`
}