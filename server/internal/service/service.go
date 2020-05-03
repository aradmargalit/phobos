package service

import (
	"server/internal/models"
	"server/internal/repository"
	"server/internal/responsetypes"

	"github.com/gin-gonic/gin"
)

// PhobosAPI defines the service methods availble to all handlers
type PhobosAPI interface {
	// OAuth Handlers for Core Authentication
	HandleLogin(*gin.Context)
	HandleCallback(*gin.Context)
	Logout(*gin.Context)

	// User Management
	GetCurrentUser(*gin.Context) responsetypes.User

	// Activities
	AddActivity(*models.Activity, int) (*models.Activity, error)
	GetActivities(uid int) (*[]responsetypes.Activity, error)
	UpdateActivity(*models.Activity) (*models.Activity, error)
	DeleteActivity(activityID int, uid int) error

	// Statistics
	GetIntervalSummary(uid int, interval string) (*[]responsetypes.IntervalSum, error)
	GetUserStatistics(uid int, offset int) (*responsetypes.Stats, error)
}

type service struct {
	db repository.PhobosDB
}

// New creates a new API service with a database connection
func New(db repository.PhobosDB) PhobosAPI {
	svc := service{}
	svc.db = db
	return &svc
}
