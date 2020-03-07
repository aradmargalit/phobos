package controllers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GetUserStatistics returns some fun user statistics for the frontend
func (e *Env) GetUserStatistics(c *gin.Context) {
	// Pull user out of context to figure out which activities to grab
	uid, ok := c.Get("user")
	if !ok {
		panic("No user id in cookie!")
	}

	a, err := e.DB.GetActivitiesByUser(uid.(int))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	// Now that we have activities, let's cronch the numbies
	totalWorkouts := len(a)
	totalHours := calculateTotalHours(a)
	totalMiles := calculateMileage(a)

	response := struct {
		Workouts int     `json:"workouts"`
		Hours    float64 `json:"hours"`
		Miles    float64 `json:"miles"`
	}{
		totalWorkouts, totalHours, totalMiles,
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
