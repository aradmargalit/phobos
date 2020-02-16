package models

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"

	// Driver for MySQL
	_ "github.com/go-sql-driver/mysql"
)

// DB allows for creating a database connection
type DB struct {
	conn *sqlx.DB
}

// Connect connects to the specified database
// The database may not yet be ready, so we're going to retry every second for 30 seconds
func (db *DB) Connect() {
	var err error
	var conn *sqlx.DB

	oneSecond, _ := time.ParseDuration("1s")

	for i := 0; i < 30; i++ {
		conn, err = sqlx.Connect("mysql", os.Getenv("API_DB_STRING"))
		if err != nil {
			fmt.Println("Error! ", err, " retrying another ", (30 - i), " times...")
			time.Sleep(oneSecond)
		} else {
			err = nil
			break
		}
	}

	if err != nil {
		panic(err)
	}

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
	err = db.conn.Get(&u, "SELECT * FROM users WHERE email=?", email)
	return
}
