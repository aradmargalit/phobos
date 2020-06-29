package repository

import (
	"errors"
	"server/internal/models"
	"strconv"
)

const (
	goalInsertSQL = `
		INSERT INTO goals 
		(user_id, metric, period, goal)
		VALUES (:user_id, :metric, :period, :goal)
	`
)

// InsertGoal adds a new user goal to the database
func (db *db) InsertGoal(a *models.Goal) (*models.Goal, error) {
	res, err := db.conn.NamedExec(goalInsertSQL, a)
	if err != nil {
		return nil, err
	}

	inserted, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Return the recently inserted record back to the user
	qa, err := db.GetGoalByID(int(inserted))
	return &qa, err
}

// UpdateGoal updates an existing goal in the database
func (db *db) UpdateGoal(g *models.Goal) (*models.Goal, error) {
	res, err := db.conn.NamedExec(
		`
		UPDATE goals 
		SET 
			goal=:goal
		WHERE id=:id
		`,
		*g)

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
	goal, err := db.GetGoalByID(g.ID)
	return &goal, err
}

// GetGoalByID returns a single goal by ID
func (db *db) GetGoalByID(id int) (g models.Goal, err error) {
	err = db.conn.Get(&g, `SELECT * FROM goals WHERE id=?`, id)
	return
}

// GetGoalsByUser returns all goals for a given user
func (db *db) GetGoalsByUser(uid int) (goals []models.Goal, err error) {
	err = db.conn.Select(&goals, `SELECT * FROM goals WHERE user_id=? ORDER BY id DESC`, uid)
	return
}

// DeleteGoalByID deletes a goal by ID, verified with userID
func (db *db) DeleteGoalByID(uid int, goalID int) (err error) {
	_, err = db.conn.Exec(`DELETE FROM goals WHERE id=? AND user_id=?`, goalID, uid)
	return
}
