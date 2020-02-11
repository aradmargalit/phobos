package models

import (
	"github.com/jmoiron/sqlx"

	// Driver for SQLite3
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS
users (
	id INTEGER PRIMARY KEY,
  name text,
  given_name text,
	email text UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

// DB allows for creating a database connection
type DB struct {
	conn *sqlx.DB
}

// Connect connects to the specified database
func (db *DB) Connect() {
	conn, err := sqlx.Connect("sqlite3", "./phobos.db")
	if err != nil {
		panic(err)
	}

	// Bootstrap the schema to create tables if needed
	conn.MustExec(schema)

	// Store the session to the db object
	db.conn = conn
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
	err = db.conn.Get(&u, "SELECT * FROM users WHERE email=$1", email)
	return
}
