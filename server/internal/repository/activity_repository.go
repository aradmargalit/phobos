package repository

import (
	"errors"
	"server/internal/models"
	"strconv"
)

// InsertActivity adds a new activity to the database
func (db *db) InsertActivity(a *models.Activity) (*models.Activity, error) {
	res, err := db.conn.NamedExec(
		`
		INSERT INTO activities 
		(name, activity_date, activity_type_id, owner_id, duration, distance, unit, heart_rate, strava_id)
		VALUES (:name, :activity_date, :activity_type_id, :owner_id, :duration, :distance, :unit, :heart_rate, :strava_id)
		`,
		*a)
	if err != nil {
		return nil, err
	}

	inserted, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Return the recently inserted record back to the user
	activity, err := db.GetActivityByID(int(inserted))
	return &activity, err
}

// GetActivityByStravaID will trade a strava activity ID for an application ID
func (db *db) GetActivityByStravaID(stravaID *int) (activity models.Activity, err error) {
	err = db.conn.Get(&activity, "SELECT * FROM activities WHERE strava_id = ?", *stravaID)
	return
}

// GetActivityByID returns a single activity by Id
func (db *db) GetActivityByID(id int) (activity models.Activity, err error) {
	err = db.conn.Get(&activity, `SELECT * FROM activities WHERE id=?`, id)
	return
}

// GetActivitiesByUser returns all activies from user
func (db *db) GetActivitiesByUser(uid int) (activities []models.ActivityResponse, err error) {
	err = db.conn.Select(&activities, `
		SELECT a.*, at.id "activity_type.id", at.name "activity_type.name" 
		FROM activities a 
		JOIN activity_types at ON at.id = a.activity_type_id 
		WHERE owner_id=? ORDER BY a.activity_date DESC
		`, uid)
	return
}

// UpdateActivity updates an existing activity in the database
func (db *db) UpdateActivity(a *models.Activity) (*models.Activity, error) {
	res, err := db.conn.NamedExec(
		`
		UPDATE activities 
		SET 
			name=:name, 
			activity_date=:activity_date,
			activity_type_id=:activity_type_id,
			duration=:duration,
			distance=:distance,
			unit=:unit,
			heart_rate=:heart_rate
		WHERE id=:id
		`,
		*a)

	if err != nil {
		return nil, err
	}

	updatedCount, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if updatedCount != 1 {
		err = errors.New("Should have updated one row, but updated: " + strconv.Itoa(int(updatedCount)))
		if err != nil {
			return nil, err
		}
	}

	// Return the recently inserted record back to the user
	activity, err := db.GetActivityByID(a.ID)
	return &activity, err
}

// DeleteActivityByID deletes an activity by ID, verified with userID
func (db *db) DeleteActivityByID(uid int, activityID int) (err error) {
	_, err = db.conn.Exec(`DELETE FROM activities WHERE id=? AND owner_id=?`, activityID, uid)
	return
}
