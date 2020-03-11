package controllers

import (
	"net/http"
	"server/models"

	"time"

	"github.com/gin-gonic/gin"
)

type typePortion struct {
	Name    string `json:"name"`
	Portion int    `json:"portion"`
}

// GetUserStatistics returns some fun user statistics for the frontend
func (e *Env) GetUserStatistics(c *gin.Context) {
	// Pull user out of context to figure out which activities to grab
	uid, ok := c.Get("user")
	if !ok {
		panic("No user id in cookie!")
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
	lastTen := calculateLastTenDays(a)
	typeBreakdown := calculateTypeBreakdown(a, at)

	response := struct {
		Workouts      int           `json:"workouts"`
		Hours         float64       `json:"hours"`
		Miles         float64       `json:"miles"`
		LastTen       []float64     `json:"last_ten"`
		TypeBreakdown []typePortion `json:"type_breakdown"`
	}{
		totalWorkouts, totalHours, totalMiles, lastTen, typeBreakdown,
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

func calculateLastTenDays(activities []models.Activity) (lastTen []float64) {
	// For each of the past 10 days, we need to sum up the durations from those days
	for i := 9; i >= 0; i-- {
		// Get the date for "i" days ago
		date := time.Now().AddDate(0, 0, -1*i)

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

	// Now, calculate proportions
	for typeID, count := range typeMap {
		typePortions = append(typePortions, typePortion{activityTypeMap[typeID], count})
	}
	return
}
