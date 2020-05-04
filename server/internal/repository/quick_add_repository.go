package repository

import "server/internal/models"

const (
	quickAddInsertSQL = `
		INSERT INTO quick_adds 
		(name, activity_type_id, owner_id, duration, distance, unit)
		VALUES (:name, :activity_type_id, :owner_id, :duration, :distance, :unit)
	`
)

// InsertQuickAdd adds a new activity to the database
func (db *db) InsertQuickAdd(a *models.QuickAdd) (*models.QuickAdd, error) {
	res, err := db.conn.NamedExec(quickAddInsertSQL, a)
	if err != nil {
		return nil, err
	}

	inserted, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Return the recently inserted record back to the user
	qa, err := db.GetQuickAddByID(int(inserted))
	return &qa, nil
}

// GetQuickAddByID returns a single quick add by ID
func (db *db) GetQuickAddByID(id int) (qa models.QuickAdd, err error) {
	err = db.conn.Get(&qa, `SELECT * FROM quick_adds WHERE id=?`, id)
	return
}

// GetQuickAddsByUser returns all activies from auser
func (db *db) GetQuickAddsByUser(uid int) (quickAdds []models.QuickAdd, err error) {
	err = db.conn.Select(&quickAdds, `SELECT * FROM quick_adds WHERE owner_id=? ORDER BY id DESC`, uid)
	return
}

// DeleteQuickAddByID deletes a quick add by ID, verified with userID
func (db *db) DeleteQuickAddByID(uid int, quickAddID int) (err error) {
	_, err = db.conn.Exec(`DELETE FROM quick_adds WHERE id=? AND owner_id=?`, quickAddID, uid)
	return
}
