package controllers

import (
	"math"
	"net/http"
	"server/models"
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
	uid, ok := c.Get("user")
	if !ok {
		panic("No user id in cookie!")
	}

	// Pull the user's timezone out of the request
	utcOffset, err := strconv.Atoi(c.Query("utc_offset"))
	if err != nil {
		panic(err)
	}

	a, err := e.DB.GetActivitiesByUser(uid.(int))
	if err != nil {
		panic(err)
	}

	at, err := e.DB.GetActivityTypes()
	if err != nil {
		panic(err)
	}

	// Now that we have activities, let's cronch the numbies
	totalWorkouts := len(a)
	totalHours := calculateTotalHours(a)
	totalMiles := calculateMileage(a)
	lastTen := calculateLastTenDays(a, utcOffset)
	typeBreakdown := calculateTypeBreakdown(a, at)
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

func calculateTotalHours(activities []models.Activity) float64 {
	var running float64 = 0

	for _, activity := range activities {
		running += activity.Duration
	}

	return running / 60
}

func calculateMileage(activities []models.Activity) float64 {
	var running float64 = 0

	for _, activity := range activities {
		if activity.Unit == "miles" {
			running += activity.Distance
		}
	}

	return running
}

func calculateLastTenDays(activities []models.Activity, utcOffset int) (lastTen []float64) {
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

func calculateTypeBreakdown(activities []models.Activity, activityTypes []models.ActivityType) (typePortions []typePortion) {
	total := float64(len(activities))
	typeMap := make(map[int]int)
	for _, activity := range activities {
		running, ok := typeMap[activity.ActivityTypeID]
		if !ok {
			running = 0
		}

		running++
		typeMap[activity.ActivityTypeID] = running
	}

	// Build a quick map of activity type ids to names
	activityTypeMap := map[int]string{}
	for _, aT := range activityTypes {
		activityTypeMap[aT.ID] = aT.Name
	}

	var insignificantTally int
	// Now, calculate proportions
	for typeID, count := range typeMap {
		// Check to make sure it's significant enough to show
		if (float64(count) / total) < .05 {
			insignificantTally += count
			continue
		}
		typePortions = append(typePortions, typePortion{activityTypeMap[typeID], count})
	}
	typePortions = append(typePortions, typePortion{"Other", insignificantTally})
	return
}

func calculateDayBreakdown(activities []models.Activity) (dayBreakdowns []dayBreakdown) {
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
