package models

// Activity represents a workout session
type Activity struct {
	ID             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	ActivityDate   string `json:"activity_date" db:"activity_date"`
	ActivityTypeID int    `json:"activity_type_id" db:"activity_type_id"`
	Duration       int    `json:"duration" db:"duration"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	UpdatedAt      string `json:"updated_at" db:"updated_at"`
}

const (
	activityInsertSQL = `INSERT INTO activities (name, activity_date, activity_type_id, duration) VALUES (:name, :activity_date, :activity_type_id, :duration)`
)

// InsertActivity adds a new activity to the database
func (db *DB) InsertActivity(a Activity) (activity Activity, err error) {
	res, err := db.conn.NamedExec(activityInsertSQL, a)
	inserted, err := res.LastInsertId()
	activity, err = db.GetActivityByID(int(inserted))
	return
}

// GetActivityByID returns a single activity by Id
func (db *DB) GetActivityByID(id int) (activity Activity, err error) {
	err = db.conn.Get(&activity, `SELECT * FROM activities WHERE id=?`, id)
	return
}
