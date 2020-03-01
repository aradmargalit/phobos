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
func (db *DB) InsertUser(u User) (user User, err error) {
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

// GetUserByID gets a database user by their ID
func (db *DB) GetUserByID(id int) (u User, err error) {
	err = db.conn.Get(&u, "SELECT * FROM users WHERE id=?", id)
	return
}