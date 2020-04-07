package controllers

import (
	"net/http"
	utils "server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	// Now that we have activities, let's cronch the numbies
	// ^ I have no idea who wrote that
	totalWorkouts := len(a)
	totalHours := utils.CalculateTotalHours(a)
	totalMiles := utils.CalculateMileage(a)
	lastTen := utils.CalculateLastTenDays(a, utcOffset)
	typeBreakdown := utils.CalculateTypeBreakdown(a)
	dayBreakdowns := utils.CalculateDayBreakdown(a)

	response := struct {
		Workouts      int                  `json:"workouts"`
		Hours         float64              `json:"hours"`
		Miles         float64              `json:"miles"`
		LastTen       []float64            `json:"last_ten"`
		TypeBreakdown []utils.TypePortion  `json:"type_breakdown"`
		DayBreakdown  []utils.DayBreakdown `json:"day_breakdown"`
	}{
		totalWorkouts, totalHours, totalMiles, lastTen, typeBreakdown, dayBreakdowns,
	}

	c.JSON(http.StatusOK, response)
}
