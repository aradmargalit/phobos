package utils

import (
	"fmt"
	"server/internal/models"
	"server/internal/responsetypes"
	"time"
)

const (
	dbLayout = "2006-01-02 15:04:05"
)

// CalculateTotalHours returns the total number of hours across all activities
func CalculateTotalHours(activities []models.ActivityResponse) float64 {
	var running float64 = 0

	for _, activity := range activities {
		running += activity.Duration
	}

	return running / 60
}

// CalculateMileage returns the total mileage across all activities
func CalculateMileage(activities []models.ActivityResponse) float64 {
	var running float64 = 0

	for _, activity := range activities {
		if activity.Unit == "miles" {
			running += activity.Distance
		}
	}

	return running
}

// CalculateLastNDays returns an array of the sum of each day's workout duration from today to 10 days ago
func CalculateLastNDays(activities *[]models.ActivityResponse, utcOffset int, n int) *[]float64 {
	var lastTen []float64

	// For each of the past N days, we need to sum up the durations from those days
	for i := n - 1; i >= 0; i-- {
		// Get the date for "i" days ago)
		// Ugly, but use the browser offset to find the correct offset
		dur, _ := time.ParseDuration(fmt.Sprintf("%vh", utcOffset))
		date := time.Now().UTC().Add(-dur).AddDate(0, 0, i*-1)

		// Start a running duration for that date
		var running float64
		for _, a := range *activities {
			// Parse the DB date
			dbDate, _ := time.Parse(dbLayout, a.ActivityDate)

			// If it's the same date, add to the running total
			if dbDate.YearDay() == date.YearDay() && dbDate.Year() == date.Year() {
				running += a.Duration
			}

			// Optimization
			// If we've hit an earlier date, break, because this is a descending sorted list
			// Give a daylong buffer to make sure we don't accidentally exclude activities that technically occur "before" the date
			// Because of time offsets
			if dbDate.Before(date.Add(time.Hour * -24)) {
				break
			}
		}

		// Append the sum to the slice
		lastTen = append(lastTen, running)
	}
	return &lastTen
}

// CalculateTypeBreakdown determines which activities contribute to which portions
func CalculateTypeBreakdown(activities []models.ActivityResponse) (typePortions []responsetypes.TypePortion) {
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

		typePortions = append(typePortions, responsetypes.TypePortion{Name: typeName, Portion: count})
	}
	typePortions = append(typePortions, responsetypes.TypePortion{Name: "Other", Portion: insignificantTally})
	return
}

// CalculateDayBreakdown buckets activities into the days of the week
func CalculateDayBreakdown(activities []models.ActivityResponse) (dayBreakdowns []responsetypes.DayBreakdown) {
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
		dayBreakdowns = append(dayBreakdowns, responsetypes.DayBreakdown{DOW: day, Count: dayMap[day]})
	}
	return
}

// CalculateThisWeek calculates the number of minutes every day for the current week
func CalculateThisWeek(activities *[]models.ActivityResponse, utcOffset int) *[]float64 {
	// Need to find what today is based on UTC offset
	dur, _ := time.ParseDuration(fmt.Sprintf("%vh", utcOffset))
	now := time.Now().UTC().Add(-dur)

	// Current day of week number
	// Casting "Sunday" to an int returns 0, subtract one to make Monday "0" and mod 7 for safety
	dow := (int(now.Weekday()) % 7) + 1

	return CalculateLastNDays(activities, utcOffset, dow)
}

// CalculateThisMonth calculates the number of minutes every day for the current month
func CalculateThisMonth(activities *[]models.ActivityResponse, utcOffset int) *[]float64 {
	// Need to find what today is based on UTC offset
	dur, _ := time.ParseDuration(fmt.Sprintf("%vh", utcOffset))
	now := time.Now().UTC().Add(-dur)

	dayOfMonth := now.Day()

	return CalculateLastNDays(activities, utcOffset, dayOfMonth)
}
