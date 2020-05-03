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
	GetActivityByStravaID(int) (models.Activity, error)
	GetActivityByID(int) (models.Activity, error)
	GetActivitiesByUser(int) ([]responsetypes.Activity, error)
	UpdateActivity(*models.Activity) (*models.Activity, error)
	DeleteActivityByID(int, int) error

	// Users
	InsertUser(models.User) (responsetypes.User, error)
	GetAllUsers() []models.User
	GetUserByEmail(string) (models.User, error)
	GetUserByID(int) (responsetypes.User, error)
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
