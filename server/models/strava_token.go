package models

import (
	"errors"
	"strconv"
)

// StravaToken represents a ready-to-add workout session
type StravaToken struct {
	UserID       int    `json:"user_id" db:"user_id"`
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	Expiry       string `json:"expiry" db:"expiry"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}

const (
	insertStravaToken = `
		INSERT INTO strava_tokens 
		(user_id, access_token, refresh_token, expiry)
		VALUES (:user_id, :access_token, :refresh_token, :expiry)
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
func (db *DB) InsertStravaToken(tok StravaToken) (dbToken StravaToken, err error) {
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
func (db *DB) GetStravaTokenByUserID(uid int) (token StravaToken, err error) {
	err = db.conn.Get(&token, `SELECT * FROM strava_tokens WHERE user_id=?`, uid)
	return
}

// UpdateStravaToken refreshes the user's strava token
func (db *DB) UpdateStravaToken(tok StravaToken) (dbToken StravaToken, err error) {
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
	}

	// Return the recently inserted record back to the user
	dbToken, err = db.GetStravaTokenByUserID(tok.UserID)
	return
}