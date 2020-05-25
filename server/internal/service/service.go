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
	AddActivity(*models.PostActivityRequest, int) (*models.Activity, error)
	GetActivities(uid int) (*[]models.ActivityResponse, error)
	UpdateActivity(*models.Activity) (*models.Activity, error)
	DeleteActivity(activityID int, uid int) error

	// Statistics
	GetIntervalSummary(uid int, interval string, offset int) (*[]responsetypes.IntervalSum, error)
	GetUserStatistics(uid int, offset int) (*responsetypes.Stats, error)

	// Quick Adds
	GetQuickAdds(uid int) (*[]models.QuickAdd, error)
	AddQuickAdd(int, *models.QuickAdd) (*models.QuickAdd, error)
	DeleteQuickAdd(uid int, quickAddID int) error

	// Metadata
	GetActivityTypes() (*[]models.ActivityType, error)

	// Seeds
	SeedActivityTypes() (err error)

	// Strava
	HandleStravaLogin(c *gin.Context)
	HandleStravaCallback(c *gin.Context)
	HandleStravaDeauthorization(uid int) error
	HandleStravaWebhookVerification(c *gin.Context)
	HandleWebhookEvent(event models.StravaWebhookEvent) error

	// Trendline Data
	GetTrendPoints(int, string, int) (*[]float64, error)
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
