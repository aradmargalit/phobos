package service

import (
	"server/internal/models"
	"server/mocks"
	"server/testdata"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConvertStravaActivity(t *testing.T) {
	// Arrange
	mockDB := new(mocks.PhobosDB)
	mockDB.On("GetActivityTypeIDByStravaType", "Run").Return(1, nil)

	stravaActivity := models.StravaActivity{
		ID:          1,
		Name:        "Best Run Yet!",
		Distance:    10000, // 10K meters
		ElapsedTime: 22,
		MovingTime:  21,
		Type:        "Run",
		StartDate:   time.Now().Format("2006-01-02T15:04:05Z"),
		Timezone:    "(GMT-08:00) America/Los_Angeles",
	}

	// Act
	activity, err := convertStravaActivity(stravaActivity, 1, mockDB)

	// Assert
	expected := models.Activity{
		ID:             0,
		Name:           "Best Run Yet!",
		ActivityDate:   time.Now().Format("2006-01-02"),
		ActivityTypeID: 1,
		OwnerID:        1,
		Duration:       0.35,
		Distance:       6.21,
		Unit:           "miles",
		HeartRate: 			testdata.MakeIntPointer(0),
		StravaID:       testdata.MakeIntPointer(1),
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, *activity)
}
