package models

import (
	"errors"
	responsetypes "server/response_types"
	"strconv"
)

// Activity represents a workout session
type Activity struct {
	ID             int     `json:"id" db:"id"`
	Name           string  `json:"name" db:"name"`
	ActivityDate   string  `json:"activity_date" db:"activity_date"`
	ActivityTypeID int     `json:"activity_type_id" db:"activity_type_id"`
	OwnerID        int     `json:"owner_id" db:"owner_id"`
	Duration       float64 `json:"duration" db:"duration"`
	Distance       float64 `json:"distance" db:"distance"`
	Unit           string  `json:"unit" db:"unit"`
	CreatedAt      string  `json:"created_at" db:"created_at"`
	UpdatedAt      string  `json:"updated_at" db:"updated_at"`
}

const (
	activityInsertSQL = `
		INSERT INTO activities 
		(name, activity_date, activity_type_id, owner_id, duration, distance, unit)
		VALUES (:name, :activity_date, :activity_type_id, :owner_id, :duration, :distance, :unit)
	`
	activityUpdateSQL = `
		UPDATE activities 
		SET 
			name=:name, 
			activity_date=:activity_date,
			activity_type_id=:activity_type_id,
			duration=:duration,
			distance=:distance,
			unit=:unit
		WHERE id=:id
		`
)

// InsertActivity adds a new activity to the database
func (db *DB) InsertActivity(a Activity) (activity Activity, err error) {
	res, err := db.conn.NamedExec(activityInsertSQL, a)
	if err != nil {
		return
	}

	inserted, err := res.LastInsertId()
	if err != nil {
		return
	}

	// Return the recently inserted record back to the user
	activity, err = db.GetActivityByID(int(inserted))
	return
}

// UpdateActivity updates an existing activity in the database
func (db *DB) UpdateActivity(a Activity) (activity Activity, err error) {
	res, err := db.conn.NamedExec(activityUpdateSQL, a)
	if err != nil {
		return
	}

	updatedCount, err := res.RowsAffected()
	if err != nil {
		return
	}
	if updatedCount != 1 {
		err = errors.New("Should have updated one row, but updated: " + strconv.Itoa(int(updatedCount)))
	}

	// Return the recently inserted record back to the user
	activity, err = db.GetActivityByID(a.ID)
	return
}

// GetActivityByID returns a single activity by Id
func (db *DB) GetActivityByID(id int) (activity Activity, err error) {
	err = db.conn.Get(&activity, `SELECT * FROM activities WHERE id=?`, id)
	return
}

// GetActivitiesByUser returns all activies from auser
func (db *DB) GetActivitiesByUser(uid int) (activities []Activity, err error) {
	err = db.conn.Select(&activities, `SELECT * FROM activities WHERE owner_id=? ORDER BY id DESC`, uid)
	return
}

// DeleteActivityByID deletes an activity by ID, verified with userID
func (db *DB) DeleteActivityByID(uid int, activityID int) (err error) {
	_, err = db.conn.Exec(`DELETE FROM activities WHERE id=? AND owner_id=?`, activityID, uid)
	return
}

// ExperimentalGetActivitiesByUser returns all activies from auser
func (db *DB) ExperimentalGetActivitiesByUser(uid int) (activities []responsetypes.ActivityResponse, err error) {
	err = db.conn.Select(&activities, `SELECT a.*, at.id "activity_type.id", at.name "activity_type.name" FROM activities a JOIN activity_types at ON at.id = a.activity_type_id WHERE owner_id=? ORDER BY a.id DESC`, uid)
	return
}
