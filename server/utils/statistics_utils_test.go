package utils

import (
	"fmt"
	"server/internal/responsetypes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getActivityResponses(n int, dateStepHours int) []responsetypes.Activity {
	activities := []responsetypes.Activity{}
	startDate, _ := time.Parse(dbLayout, "2001-01-02 03:04:05")

	for i := 1; i <= n; i++ {
		step, _ := time.ParseDuration(fmt.Sprintf("%vh", dateStepHours*i))
		activities = append(activities, responsetypes.Activity{
			ID:           i,
			Name:         fmt.Sprintf("Activity Number: %v", i),
			ActivityDate: startDate.Add(step).Format(dbLayout),
			Duration:     1,
			Unit:         "miles",
			Distance:     1,
		})
	}
	return activities
}

func TestCalculateTotalHours(t *testing.T) {
	// Test that the total hours add up correctly
	// Create 20 1 minute activities
	activities := getActivityResponses(20, 24)
	assert.Equal(t, (float64(20) / float64(60)), CalculateTotalHours(activities))
}

func TestCalculateMileage(t *testing.T) {
	// Create 20 1 mile runs
	activities := getActivityResponses(20, 24)
	assert.Equal(t, float64(20), CalculateMileage(activities))

	// Yardage doesn't get counted
	activities = append(activities, responsetypes.Activity{Unit: "yards", Distance: 1})
	assert.Equal(t, float64(20), CalculateMileage(activities))

	// All mileage conuts
	activities = append(activities, responsetypes.Activity{Unit: "miles", Distance: 1})
	assert.Equal(t, float64(21), CalculateMileage(activities))
}

func TestCalculateLastTenDays(t *testing.T) {
	activities := []responsetypes.Activity{{Duration: 10, ActivityDate: time.Now().UTC().Format(dbLayout)}}
	want := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 10}
	assert.Equal(t, want, CalculateLastTenDays(activities, 0))
}

func TestCalculateTypeBreakdown(t *testing.T) {
	activities := []responsetypes.Activity{
		{ActivityType: responsetypes.ActivityType{Name: "Run"}},
		{ActivityType: responsetypes.ActivityType{Name: "Run"}},
		{ActivityType: responsetypes.ActivityType{Name: "Run"}},
		{ActivityType: responsetypes.ActivityType{Name: "Swim"}},
	}

	portions := CalculateTypeBreakdown(activities)
	assert.Contains(t, portions, responsetypes.TypePortion{Name: "Run", Portion: 3})
	assert.Contains(t, portions, responsetypes.TypePortion{Name: "Other", Portion: 0})
}

func TestCalculateTypeBreakdownWithOtherActivities(t *testing.T) {
	activities := []responsetypes.Activity{}
	for i := 0; i < 100; i++ {
		activities = append(activities, responsetypes.Activity{ActivityType: responsetypes.ActivityType{Name: "Run"}})
	}
	activities = append(activities, responsetypes.Activity{ActivityType: responsetypes.ActivityType{Name: "Jog"}})
	activities = append(activities, responsetypes.Activity{ActivityType: responsetypes.ActivityType{Name: "Swim"}})

	portions := CalculateTypeBreakdown(activities)
	assert.Contains(t, portions, responsetypes.TypePortion{Name: "Run", Portion: 100})
	assert.Contains(t, portions, responsetypes.TypePortion{Name: "Other", Portion: 2})
	assert.NotContains(t, portions, responsetypes.TypePortion{Name: "Jog", Portion: 1})
}

func TestCalculateDayBreakdown(t *testing.T) {
	// Create 20 1 mile runs
	activities := getActivityResponses(20, 24)
	want := []responsetypes.DayBreakdown([]responsetypes.DayBreakdown{
		{DOW: "Monday", Count: 3},
		{DOW: "Tuesday", Count: 2},
		{DOW: "Wednesday", Count: 3},
		{DOW: "Thursday", Count: 3},
		{DOW: "Friday", Count: 3},
		{DOW: "Saturday", Count: 3},
		{DOW: "Sunday", Count: 3},
	})

	assert.Equal(t, want, CalculateDayBreakdown(activities))
}
