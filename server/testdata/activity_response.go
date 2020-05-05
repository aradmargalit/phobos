package testdata

import (
	"fmt"
	"server/internal/constants"
	"server/internal/responsetypes"
	"time"
)

// GetActivityResponses provides any number of mock activity responses
func GetActivityResponses(n int, dateStepHours int) []responsetypes.Activity {
	activities := []responsetypes.Activity{}
	startDate, _ := time.Parse(constants.DBLayout, "2001-01-02 03:04:05")

	for i := 1; i <= n; i++ {
		step, _ := time.ParseDuration(fmt.Sprintf("%vh", dateStepHours*(i-1)))
		activities = append(activities, responsetypes.Activity{
			ID:           i,
			Name:         fmt.Sprintf("Activity Number: %v", i),
			ActivityDate: startDate.Add(step).Format(constants.DBLayout),
			ActivityType: responsetypes.ActivityType{ID: 1, Name: "Run"},
			Duration:     1,
			Unit:         "miles",
			Distance:     1,
		})
	}
	return activities
}