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

// InsertUser inserts a User into the database, if possible
func (db *DB) InsertUser(u User) (err error) {
	_, err = db.conn.NamedExec(`INSERT INTO users (name, given_name, email) VALUES (:name, :given_name, :email)`, u)

	return
}

// GetAllUsers gets all users from the database
func (db *DB) GetAllUsers() []User {
	people := []User{}
	err := db.conn.Select(&people, "SELECT * FROM `users` ORDER BY id ASC")
	if err != nil {
		panic(err)
	}

	return people
}

// GetUserByEmail gets a database user by their email
func (db *DB) GetUserByEmail(email string) (u User, err error) {
	err = db.conn.Get(&u, "SELECT * FROM users WHERE email=?", email)
	return
}
