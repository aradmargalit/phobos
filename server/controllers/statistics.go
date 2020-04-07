package controllers

import (
	"math"
	"net/http"
	responsetypes "server/response_types"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
)

type typePortion struct {
	Name    string `json:"name"`
	Portion int    `json:"portion"`
}

type dayBreakdown struct {
	DOW   string `json:"day_of_week"`
	Count int    `json:"count"`
}

// GetUserStatistics returns some fun user statistics for the frontend
func (e *Env) GetUserStatistics(c *gin.Context) {
	// Pull user out of context to figure out which activities to grab
	uid := c.GetInt("user")

	// Pull the user's timezone out of the request
	utcOffset, err := strconv.Atoi(c.Query("utc_offset"))
	if err != nil {
		panic(err)
	}

	a, err := e.DB.GetActivitiesByUser(uid)
	if err != nil {
		panic(err)
	}

	// Now that we have activities, let's cronch the numbies
	totalWorkouts := len(a)
	totalHours := calculateTotalHours(a)
	totalMiles := calculateMileage(a)
	lastTen := calculateLastTenDays(a, utcOffset)
	typeBreakdown := calculateTypeBreakdown(a)
	dayBreakdowns := calculateDayBreakdown(a)

	response := struct {
		Workouts      int            `json:"workouts"`
		Hours         float64        `json:"hours"`
		Miles         float64        `json:"miles"`
		LastTen       []float64      `json:"last_ten"`
		TypeBreakdown []typePortion  `json:"type_breakdown"`
		DayBreakdown  []dayBreakdown `json:"day_breakdown"`
	}{
		totalWorkouts, totalHours, totalMiles, lastTen, typeBreakdown, dayBreakdowns,
	}

	c.JSON(http.StatusOK, response)
}

func calculateTotalHours(activities []responsetypes.ActivityResponse) float64 {
	var running float64 = 0

	for _, activity := range activities {
		running += activity.Duration
	}

	return running / 60
}

func calculateMileage(activities []responsetypes.ActivityResponse) float64 {
	var running float64 = 0

	for _, activity := range activities {
		if activity.Unit == "miles" {
			running += activity.Distance
		}
	}

	return running
}

func calculateLastTenDays(activities []responsetypes.ActivityResponse, utcOffset int) (lastTen []float64) {
	// For each of the past 10 days, we need to sum up the durations from those days
	for i := 9; i >= 0; i-- {
		// Get the date for "i" days ago)
		// Ugly, but use the browser offset to find the correct offset
		adjustment := (int(math.Floor(float64(-utcOffset) / 24))) + i*-1
		date := time.Now().UTC().AddDate(0, 0, adjustment)

		// Start a running duration for that date
		var running float64
		for _, a := range activities {
			// Parse the DB date
			dbDate, _ := time.Parse("2006-01-02 15:04:05", a.ActivityDate)

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

func calculateTypeBreakdown(activities []responsetypes.ActivityResponse) (typePortions []typePortion) {
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
	// Now, calculate proportions
	for typeName, count := range typeMap {
		// Check to make sure it's significant enough to show
		if (float64(count) / total) < .05 {
			insignificantTally += count
			continue
		}

		typePortions = append(typePortions, typePortion{typeName, count})
	}
	typePortions = append(typePortions, typePortion{"Other", insignificantTally})
	return
}

func calculateDayBreakdown(activities []responsetypes.ActivityResponse) (dayBreakdowns []dayBreakdown) {
	dayMap := make(map[string]int)
	for _, activity := range activities {
		// Parse date from activity
		dbDate, _ := time.Parse("2006-01-02 15:04:05", activity.ActivityDate)
		dow := dbDate.Weekday().String()

		running, ok := dayMap[dow]
		if !ok {
			running = 0
		}

		running++
		dayMap[dow] = running
	}

	for _, day := range []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"} {
		dayBreakdowns = append(dayBreakdowns, dayBreakdown{day, dayMap[day]})
	}
	return
}
