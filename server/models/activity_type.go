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

// DeleteAllActivityTypes deletes all activity types, should only be used during seeding
func (db *DB) DeleteAllActivityTypes() (err error) {
	_, err = db.conn.Exec(`DELETE FROM activity_types`)
	if err != nil {
		return
	}

	_, err = db.conn.Exec(`ALTER TABLE activity_types AUTO_INCREMENT=1`)

	return
}

// GetActivityTypeIDByStravaType swaps a Strava activity string to an ID
func (db *DB) GetActivityTypeIDByStravaType(stravaType string) (typeID int, err error) {
	err = db.conn.Get(&typeID, `SELECT id FROM activity_types WHERE name=?`, stravaType)
	return
}
