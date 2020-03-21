package models

// QuickAdd represents a ready-to-add workout session
type QuickAdd struct {
	ID             int     `json:"id" db:"id"`
	Name           string  `json:"name" db:"name"`
	ActivityTypeID int     `json:"activity_type_id" db:"activity_type_id"`
	OwnerID        int     `json:"owner_id" db:"owner_id"`
	Duration       float64 `json:"duration" db:"duration"`
	Distance       float64 `json:"distance" db:"distance"`
	Unit           string  `json:"unit" db:"unit"`
	CreatedAt      string  `json:"created_at" db:"created_at"`
	UpdatedAt      string  `json:"updated_at" db:"updated_at"`
}

const (
	quickAddInsertSQL = `
		INSERT INTO quick_adds 
		(name, activity_type_id, owner_id, duration, distance, unit)
		VALUES (:name, :activity_type_id, :owner_id, :duration, :distance, :unit)
	`
)

// InsertQuickAdd adds a new activity to the database
func (db *DB) InsertQuickAdd(a QuickAdd) (qa QuickAdd, err error) {
	res, err := db.conn.NamedExec(quickAddInsertSQL, a)
	if err != nil {
		return
	}

	inserted, err := res.LastInsertId()
	if err != nil {
		return
	}

	// Return the recently inserted record back to the user
	qa, err = db.GetQuickAddByID(int(inserted))
	return
}

// GetQuickAddByID returns a single quick add by ID
func (db *DB) GetQuickAddByID(id int) (qa QuickAdd, err error) {
	err = db.conn.Get(&qa, `SELECT * FROM quick_adds WHERE id=?`, id)
	return
}

// GetQuickAddsByUser returns all activies from auser
func (db *DB) GetQuickAddsByUser(uid int) (quickAdds []QuickAdd, err error) {
	err = db.conn.Select(&quickAdds, `SELECT * FROM quick_adds WHERE owner_id=? ORDER BY id DESC`, uid)
	return
}

// DeleteQuickAddByID deletes a quick add by ID, verified with userID
func (db *DB) DeleteQuickAddByID(uid int, quickAddID int) (err error) {
	_, err = db.conn.Exec(`DELETE FROM quick_adds WHERE id=? AND owner_id=?`, quickAddID, uid)
	return
}
