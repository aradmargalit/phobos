package repository

import (
	"server/internal/models"
	"server/internal/responsetypes"
)

// InsertUser inserts a User into the database, if possible
func (db *db) InsertUser(u models.User) (user responsetypes.User, err error) {
	result, err := db.conn.NamedExec(`INSERT INTO users (name, given_name, email) VALUES (:name, :given_name, :email)`, u)
	if err != nil {
		panic(err)
	}

	// Get the ID for the inserted user
	insertedID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// Grab the user from the database
	user, err = db.GetUserByID(int(insertedID))
	if err != nil {
		panic(err)
	}

	return
}

// GetAllUsers gets all users from the database
func (db *db) GetAllUsers() []models.User {
	people := []models.User{}
	err := db.conn.Select(&people, "SELECT * FROM `users` ORDER BY id ASC")
	if err != nil {
		panic(err)
	}

	return people
}

// GetUserByEmail gets a database user by their email
func (db *db) GetUserByEmail(email string) (u models.User, err error) {
	err = db.conn.Get(&u, "SELECT * FROM users WHERE email=?", email)
	return
}

// GetUserByID gets a database user by their ID
func (db *db) GetUserByID(id int) (u responsetypes.User, err error) {
	// This warrants an explanation!
	// I want to deserialize this query response to a responsetypes.UserResponse object
	// which expects "strava_token" to be a boolean, so I check for existance and convert
	// to a boolean in the SQL itself.
	err = db.conn.Get(&u, `
	SELECT users.*, 
		IF(strava_tokens.expiry IS NULL, FALSE, TRUE) as strava_token
	FROM users
	LEFT OUTER JOIN 
		strava_tokens ON strava_tokens.user_id = users.id
	WHERE users.id = ?
	`, id)
	return
}
