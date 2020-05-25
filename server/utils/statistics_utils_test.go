package utils

import (
	"fmt"
	"server/internal/models"
	"server/internal/responsetypes"
	"server/testdata"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTotalHours(t *testing.T) {
	// Test that the total hours add up correctly
	// Create 20 1 minute activities
	activities := testdata.GetActivityResponses(20, 24)
	assert.Equal(t, (float64(20) / float64(60)), CalculateTotalHours(activities))
}

func TestCalculateMileage(t *testing.T) {
	// Create 20 1 mile runs
	activities := testdata.GetActivityResponses(20, 24)
	assert.Equal(t, float64(20), CalculateMileage(activities))

	// Yardage doesn't get counted
	activities = append(activities, models.ActivityResponse{Activity: models.Activity{Unit: "yards", Distance: 1}})
	assert.Equal(t, float64(20), CalculateMileage(activities))

	// All mileage conuts
	activities = append(activities, models.ActivityResponse{Activity: models.Activity{Unit: "miles", Distance: 1}})
	assert.Equal(t, float64(21), CalculateMileage(activities))
}

func TestCalculateTypeBreakdown(t *testing.T) {
	activities := []models.ActivityResponse{
		{ActivityType: models.ActivityType{Name: "Run"}},
		{ActivityType: models.ActivityType{Name: "Run"}},
		{ActivityType: models.ActivityType{Name: "Run"}},
		{ActivityType: models.ActivityType{Name: "Swim"}},
	}

	portions := CalculateTypeBreakdown(activities)
	assert.Contains(t, portions, responsetypes.TypePortion{Name: "Run", Portion: 3})
	assert.Contains(t, portions, responsetypes.TypePortion{Name: "Other", Portion: 0})
}

func TestCalculateTypeBreakdownWithOtherActivities(t *testing.T) {
	activities := []models.ActivityResponse{}
	for i := 0; i < 100; i++ {
		activities = append(activities, models.ActivityResponse{ActivityType: models.ActivityType{Name: "Run"}})
	}
	activities = append(activities, models.ActivityResponse{ActivityType: models.ActivityType{Name: "Jog"}})
	activities = append(activities, models.ActivityResponse{ActivityType: models.ActivityType{Name: "Swim"}})

	portions := CalculateTypeBreakdown(activities)
	assert.Contains(t, portions, responsetypes.TypePortion{Name: "Run", Portion: 100})
	assert.Contains(t, portions, responsetypes.TypePortion{Name: "Other", Portion: 2})
	assert.NotContains(t, portions, responsetypes.TypePortion{Name: "Jog", Portion: 1})
}

func TestCalculateDayBreakdown(t *testing.T) {
	// Create 20 1 mile runs
	activities := testdata.GetActivityResponses(20, 24)
	want := []responsetypes.DayBreakdown([]responsetypes.DayBreakdown{
		{DOW: "Monday", Count: 2},
		{DOW: "Tuesday", Count: 3},
		{DOW: "Wednesday", Count: 3},
		{DOW: "Thursday", Count: 3},
		{DOW: "Friday", Count: 3},
		{DOW: "Saturday", Count: 3},
		{DOW: "Sunday", Count: 3},
	})

	assert.Equal(t, want, CalculateDayBreakdown(activities))
}

// Tests that this works as expected with contiguous days
func TestCalculateLastNDaysDense(t *testing.T) {
	// Arrange
	// Set up 3 activities
	var activities []models.ActivityResponse
	for i := 0; i < 3; i++ {
		activities = append(activities, models.ActivityResponse{
			Activity: models.Activity{
				Duration:     10,
				ActivityDate: time.Now().UTC().Add(time.Hour * time.Duration((-24 * i))).Format(dbLayout),
			},
		})
	}

	// Act
	scenarios := []struct{
		n int
		want []float64
	} {
		{
			n: 0,
			want: []float64(nil),
		},
		{
			n: 1,
			want: []float64{10},
		},
		{
			n: 2,
			want: []float64{10, 10},
		},
		{
			n: 3,
			want: []float64{10, 10, 10},
		},
		{
			n: 9,
			want: []float64{0, 0, 0, 0, 0, 0, 10, 10, 10},
		},
	}

	// Assert
	for _, scenario := range scenarios {
		got := CalculateLastNDays(&activities, 0, scenario.n)
		assert.Equal(t, scenario.want, *got)
	}
}

func TestCalculateLastTenDays(t *testing.T) {
	activities := []models.ActivityResponse{
		{
			Activity: models.Activity{
				Duration:     10,
				ActivityDate: time.Now().UTC().Format(dbLayout),
			},
		},
	}
	want := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 10}
	got := CalculateLastNDays(&activities, 0, 10)
	assert.Equal(t, want, *got)
}