package testdata

import (
	"fmt"
	"server/internal/constants"
	"server/internal/models"
	"time"
)

// GetActivityResponses provides any number of mock activity responses
func GetActivityResponses(n int, dateStepHours int) []models.ActivityResponse {
	activities := []models.ActivityResponse{}
	startDate, _ := time.Parse(constants.DBLayout, "2001-01-02 03:04:05")

	for i := 1; i <= n; i++ {
		step, _ := time.ParseDuration(fmt.Sprintf("%vh", dateStepHours*(i-1)))
		activities = append(activities, models.ActivityResponse{
			Activity: models.Activity{
				ID:           i,
				Name:         fmt.Sprintf("Activity Number: %v", i),
				ActivityDate: startDate.Add(step).Format(constants.DBLayout),
				Duration:     1,
				Unit:         "miles",
				Distance:     1,
			},
			ActivityType: models.ActivityType{ID: 1, Name: "Run"},
		})
	}
	return activities
}