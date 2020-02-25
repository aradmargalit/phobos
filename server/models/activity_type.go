package models

// ActivityType represents a type for a workout
type ActivityType struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

// GetActivityTypes gets all activity types from the database
func (db *DB) GetActivityTypes() (at []ActivityType, err error) {
	at = []ActivityType{}
	err = db.conn.Select(&at, "SELECT * FROM `activity_types`;")
	if err != nil {
		return nil, err
	}

	return
}

// InsertActivityType adds a new activity type to the database
func (db *DB) InsertActivityType(at ActivityType) (err error) {
	_, err = db.conn.NamedExec(`INSERT INTO activity_types (name) VALUES (:name)`, at)

	return
}
