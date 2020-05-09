package repository

import (
	"errors"
	"server/internal/models"
	"strconv"
)

const (
	insertStravaToken = `
		INSERT INTO strava_tokens 
		(user_id, strava_id, access_token, refresh_token, expiry)
		VALUES (:user_id, :strava_id, :access_token, :refresh_token, :expiry)
	`

	updateStravaTokenSQL = `
		UPDATE strava_tokens 
		SET 
			access_token=:access_token,
			refresh_token=:refresh_token,
			expiry=:expiry
		WHERE user_id=:user_id
		`
)

// InsertStravaToken registers a new set of tokens for a user's Strava access
func (db *db) InsertStravaToken(tok models.StravaToken) (dbToken models.StravaToken, err error) {
	res, err := db.conn.NamedExec(insertStravaToken, tok)
	if err != nil {
		return
	}

	_, err = res.LastInsertId()
	if err != nil {
		return
	}

	// Return the recently inserted record back to the user
	dbToken, err = db.GetStravaTokenByUserID(tok.UserID)
	return
}

// GetStravaTokenByUserID returns all activies from auser
func (db *db) GetStravaTokenByUserID(uid int) (token models.StravaToken, err error) {
	err = db.conn.Get(&token, `SELECT * FROM strava_tokens WHERE user_id=?`, uid)
	return
}

// UpdateStravaToken refreshes the user's strava token
func (db *db) UpdateStravaToken(tok models.StravaToken) (dbToken models.StravaToken, err error) {
	res, err := db.conn.NamedExec(updateStravaTokenSQL, tok)
	if err != nil {
		return
	}

	updatedCount, err := res.RowsAffected()
	if err != nil {
		return
	}
	if updatedCount != 1 {
		err = errors.New("Should have updated one row, but updated: " + strconv.Itoa(int(updatedCount)))
		if err != nil {
			return models.StravaToken{}, err
		}
	}

	// Return the recently inserted record back to the user
	dbToken, err = db.GetStravaTokenByUserID(tok.UserID)
	return
}

// GetUserIDByStravaID swaps a Strava ID for a User ID
func (db *db) GetUserIDByStravaID(stravaID int) (userID int, err error) {
	err = db.conn.Get(&userID, `SELECT user_id FROM strava_tokens WHERE strava_id=?`, stravaID)
	return
}

// DeleteStravaTokenByUserID clears out a user's token
func (db *db) DeleteStravaTokenByUserID(uid int) (err error) {
	_, err = db.conn.Exec("DELETE FROM strava_tokens WHERE user_id=?", uid)
	return
}
