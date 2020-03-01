package models

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
	activityInsertSQL = `INSERT INTO activities (name, activity_date, activity_type_id, owner_id, duration, distance, unit) VALUES (:name, :activity_date, :activity_type_id, :owner_id, :duration, :distance, :unit)`
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

// GetActivityByID returns a single activity by Id
func (db *DB) GetActivityByID(id int) (activity Activity, err error) {
	err = db.conn.Get(&activity, `SELECT * FROM activities WHERE id=?`, id)
	return
}

// GetActivitiesByUser returns all activies from auser
func (db *DB) GetActivitiesByUser(uid int) (activities []Activity, err error) {
	err = db.conn.Select(&activities, `SELECT * FROM activities WHERE owner_id=?`, uid)
	return
}
