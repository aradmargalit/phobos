package utils

import (
	"fmt"
	"time"

	responsetypes "server/response_types"
)

const (
	dbLayout = "2006-01-02 15:04:05"
)

// TypePortion represents an activity type and it's portion relative to others
type TypePortion struct {
	Name    string `json:"name"`
	Portion int    `json:"portion"`
}

// DayBreakdown represents the day of the week and how many activities have fallen on that day
type DayBreakdown struct {
	DOW   string `json:"day_of_week"`
	Count int    `json:"count"`
}

// CalculateTotalHours returns the total number of hours across all activities
func CalculateTotalHours(activities []responsetypes.ActivityResponse) float64 {
	var running float64 = 0

	for _, activity := range activities {
		running += activity.Duration
	}

	return running / 60
}

// CalculateMileage returns the total mileage across all activities
func CalculateMileage(activities []responsetypes.ActivityResponse) float64 {
	var running float64 = 0

	for _, activity := range activities {
		if activity.Unit == "miles" {
			running += activity.Distance
		}
	}

	return running
}

// CalculateLastTenDays returns an array of the sum of each day's workout duration from today to 10 days ago
func CalculateLastTenDays(activities []responsetypes.ActivityResponse, utcOffset int) (lastTen []float64) {
	// For each of the past 10 days, we need to sum up the durations from those days
	for i := 9; i >= 0; i-- {
		// Get the date for "i" days ago)
		// Ugly, but use the browser offset to find the correct offset
		dur, _ := time.ParseDuration(fmt.Sprintf("%vh", utcOffset))
		date := time.Now().UTC().Add(-dur).AddDate(0, 0, i*-1)

		// Start a running duration for that date
		var running float64
		for _, a := range activities {
			// Parse the DB date
			dbDate, _ := time.Parse(dbLayout, a.ActivityDate)

			// If it's the same date, add to the running total
			if dbDate.YearDay() == date.YearDay() && dbDate.Year() == date.Year() {
				running += a.Duration
			}
		}

		// Append the sum to the slice
		lastTen = append(lastTen, running)
	}
	return
}

// CalculateTypeBreakdown determines which activities contribute to which portions
func CalculateTypeBreakdown(activities []responsetypes.ActivityResponse) (typePortions []TypePortion) {
	total := float64(len(activities))
	typeMap := make(map[string]int)
	for _, activity := range activities {
		running, ok := typeMap[activity.ActivityType.Name]
		if !ok {
			running = 0
		}

		running++
		typeMap[activity.ActivityType.Name] = running
	}

	var insignificantTally int
	// Now, Calculate proportions
	for typeName, count := range typeMap {
		// Check to make sure it's significant enough to show
		if (float64(count) / total) < .05 {
			insignificantTally += count
			continue
		}

		typePortions = append(typePortions, TypePortion{typeName, count})
	}
	typePortions = append(typePortions, TypePortion{"Other", insignificantTally})
	return
}

// CalculateDayBreakdown buckets activities into the days of the week
func CalculateDayBreakdown(activities []responsetypes.ActivityResponse) (dayBreakdowns []DayBreakdown) {
	dayMap := make(map[string]int)
	for _, activity := range activities {
		// Parse date from activity
		dbDate, _ := time.Parse(dbLayout, activity.ActivityDate)
		dow := dbDate.Weekday().String()

		running, ok := dayMap[dow]
		if !ok {
			running = 0
		}

		running++
		dayMap[dow] = running
	}

	for _, day := range []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"} {
		dayBreakdowns = append(dayBreakdowns, DayBreakdown{day, dayMap[day]})
	}
	return
}
