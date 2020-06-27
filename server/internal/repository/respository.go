package repository

import (
	"server/internal/models"
	"server/internal/responsetypes"

	"github.com/jmoiron/sqlx"
)

// PhobosDB defines the methods available to access data
type PhobosDB interface {
	// Activities
	InsertActivity(*models.Activity) (*models.Activity, error)
	GetActivityByStravaID(*int) (models.Activity, error)
	GetActivityByID(int) (models.Activity, error)
	GetAllActivities() ([]models.Activity, error)
	GetActivitiesByUser(int) ([]models.ActivityResponse, error)
	UpdateActivity(*models.Activity) (*models.Activity, error)
	DeleteActivityByID(int, int) error

	// Users
	InsertUser(models.User) (*responsetypes.User, error)
	GetUserByEmail(string) (models.User, error)
	GetUserByID(int) (responsetypes.User, error)

	// Quick Adds
	InsertQuickAdd(a *models.QuickAdd) (*models.QuickAdd, error)
	GetQuickAddByID(id int) (qa models.QuickAdd, err error)
	GetQuickAddsByUser(uid int) (quickAdds []models.QuickAdd, err error)
	DeleteQuickAddByID(uid int, quickAddID int) (err error)

	// Metadata
	GetActivityTypes() ([]models.ActivityType, error)
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
	err := db.Connect()
	if err != nil {
		// No choice but to panic here, we cannot proceed without a database connection
		panic(err)
	}

	return &db
}
