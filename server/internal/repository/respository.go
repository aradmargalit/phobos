package repository

import (
	"database/sql"
	"server/internal/models"
	"server/internal/responsetypes"

	"github.com/jmoiron/sqlx"
)

// PhobosDB defines the methods available to access data
type PhobosDB interface {
	// Activities
	InsertActivity(*models.Activity) (*models.Activity, error)
	GetActivityByStravaID(sql.NullInt64) (models.Activity, error)
	GetActivityByID(int) (models.Activity, error)
	GetActivitiesByUser(int) ([]responsetypes.Activity, error)
	UpdateActivity(*models.Activity) (*models.Activity, error)
	DeleteActivityByID(int, int) error

	// Users
	InsertUser(models.User) (responsetypes.User, error)
	GetAllUsers() []models.User
	GetUserByEmail(string) (models.User, error)
	GetUserByID(int) (responsetypes.User, error)

	// Quick Adds
	InsertQuickAdd(a *models.QuickAdd) (*models.QuickAdd, error)
	GetQuickAddByID(id int) (qa models.QuickAdd, err error)
	GetQuickAddsByUser(uid int) (quickAdds []models.QuickAdd, err error)
	DeleteQuickAddByID(uid int, quickAddID int) (err error)

	// Metadata
	GetActivityTypes() ([]responsetypes.ActivityType, error)
	DeleteAllActivityTypes() error
	InsertActivityType(models.ActivityType) error

	// Strava
	InsertStravaToken(tok models.StravaToken) (dbToken models.StravaToken, err error)
	GetStravaTokenByUserID(uid int) (token models.StravaToken, err error)
	UpdateStravaToken(tok models.StravaToken) (dbToken models.StravaToken, err error)
	GetUserIDByStravaID(stravaID int) (userID int, err error)
	DeleteStravaTokenByUserID(uid int) error
	GetActivityTypeIDByStravaType(string) (int, error)
}

// db will be our data access object and holds the connection
type db struct {
	conn *sqlx.DB
}

// New initializes a new PhobosDB, connects, and makes the DB available to the service
func New() PhobosDB {
	db := db{}
	db.Connect()

	return &db
}