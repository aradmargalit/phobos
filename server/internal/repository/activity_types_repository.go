package repository

import (
	"server/internal/models"
)

// GetActivityTypes gets all activity types from the database
func (db *db) GetActivityTypes() (at []models.ActivityType, err error) {
	err = db.conn.Select(&at, "SELECT * FROM `activity_types` ORDER BY name ASC;")
	if err != nil {
		return nil, err
	}

	return
}

// InsertActivityType adds a new activity type to the database
func (db *db) InsertActivityType(at models.ActivityType) (err error) {
	_, err = db.conn.NamedExec(`INSERT INTO activity_types (name, strava_name) VALUES (:name, :strava_name)`, at)

	return
}

// DeleteAllActivityTypes deletes all activity types, should only be used during seeding
func (db *db) DeleteAllActivityTypes() (err error) {
	_, err = db.conn.Exec(`DELETE FROM activity_types`)
	if err != nil {
		return
	}

	_, err = db.conn.Exec(`ALTER TABLE activity_types AUTO_INCREMENT=1`)

	return
}

// GetActivityTypeIDByStravaType swaps a Strava activity string to an ID
func (db *db) GetActivityTypeIDByStravaType(stravaType string) (typeID int, err error) {
	err = db.conn.Get(&typeID, `SELECT id FROM activity_types WHERE strava_name=?`, stravaType)
	return
}
