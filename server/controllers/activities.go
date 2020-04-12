package controllers

import (
	"fmt"
	"net/http"
	models "server/models"
	responsetypes "server/response_types"
	"server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// AddActivityHandler adds a new activity to the database
func (e *Env) AddActivityHandler(c *gin.Context) {
	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// MySQL doesn't like RFC3339 times, so convert it to YYYY-MM-DD
	d, err := time.Parse(time.RFC3339, activity.ActivityDate)
	activity.ActivityDate = d.Format("2006-01-02")

	// Add the owner ID to the activituy
	uid := c.GetInt("user")

	activity.OwnerID = uid

	record, err := e.DB.InsertActivity(activity)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Currently not consumed by the UI, but echo back the record
	c.JSON(http.StatusOK, record)
}

// UpdateActivityHandler adds a new activity to the database
func (e *Env) UpdateActivityHandler(c *gin.Context) {
	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// MySQL doesn't like RFC3339 times, so convert it to YYYY-MM-DD
	d, err := time.Parse(time.RFC3339, activity.ActivityDate)
	activity.ActivityDate = d.Format("2006-01-02")

	record, err := e.DB.UpdateActivity(activity)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, record)
}

// GetActivitiesHandler returns all the user's activities
func (e *Env) GetActivitiesHandler(c *gin.Context) {
	// Pull user out of context to figure out which activities to grab
	uid := c.GetInt("user")

	a, err := e.DB.GetActivitiesByUser(uid)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	count := len(a)
	fmt.Printf("Found %v activities for user ID: %v...\n", count, uid)

	withIndices := make([]responsetypes.ActivityResponse, count)
	// No smart way to do this, add an decreasing logical index to each for the frontend's benefit
	// I also want to represent the date in an fast-to-sort fashion, so doing that here
	for idx, activity := range a {
		activity.LogicalIndex = count - idx

		// Convert date to seconds since epoch, much faster to sort ints in the UI than cast to Date objects
		t, _ := time.Parse("2006-01-02 15:04:05", activity.ActivityDate)
		activity.Epoch = t.Unix()
		withIndices[idx] = activity
	}

	c.JSON(http.StatusOK, gin.H{"activities": withIndices})
}

// DeleteActivityHandler returns all the user's activities
func (e *Env) DeleteActivityHandler(c *gin.Context) {
	// Pull user out of context to confirm it's safe to delete the activity
	uid := c.GetInt("user")

	activityID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = e.DB.DeleteActivityByID(uid, activityID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "Successfully deleted activity: "+c.Param("id"))
	return
}

// GetMonthlySums returns the user's monthly sum of workout hours and miles
func (e *Env) GetMonthlySums(c *gin.Context) {
	// Pull user out of context to figure out which activities to grab
	uid := c.GetInt("user")

	a, err := e.DB.GetActivitiesByUser(uid)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// In order to figure out which days have no activities, we start with an array with every day from
	// their first activity until now
	firstActivityDate, _ := time.Parse("2006-01-02 15:04:05", a[len(a)-1].ActivityDate)

	// Make a map of days to whether or not they were skipped
	skippedMap := map[time.Time]bool{}

	// Start by assuming each date is skipped, we'll invalidate that assumption as we go
	for d := firstActivityDate; !utils.DateEqual(d, time.Now().AddDate(0, 0, 1)); d = d.AddDate(0, 0, 1) {
		skippedMap[utils.RoundTimeToDay(d)] = true
	}

	/* This is in the format:
	{
		"January 2019": {
			"duration": 12.123,
			"distance": 12.231256,
			"daysSkipped": 12
		}
	}
	*/
	monthMap := map[string]map[string]float64{}

	for _, activity := range a {
		m, _ := time.Parse("2006-01-02 15:04:05", activity.ActivityDate)

		// Go into our date map and mark the date as unskipped
		skippedMap[utils.RoundTimeToDay(m)] = false

		// Format is "January 2020"
		month := fmt.Sprintf("%v %v", m.Month(), m.Year())
		_, ok := monthMap[month]

		// If !ok, we've never seen this month before, so initialize it to 0s
		if !ok {
			monthMap[month] = map[string]float64{"duration": 0, "miles": 0, "days_skipped": 0}
		}

		// Otherwise, add duration, and distance (if mileage)
		monthMap[month]["duration"] += activity.Duration
		if activity.Unit == "miles" {
			monthMap[month]["miles"] += activity.Distance
		}
	}

	// After going through every activity, we need to summarize the skipped days
	for month, payload := range monthMap {
		m := strings.Split(month, " ")[0]
		y := strings.Split(month, " ")[1]

		// For each skipped activity, see if it matches, and if so, add it to the tally
		for sA, wasSkipped := range skippedMap {
			if wasSkipped && sA.Month().String() == m && strconv.Itoa(sA.Year()) == y {
				newPayload := payload
				newPayload["days_skipped"]++
				monthMap[month] = newPayload
			}
		}
	}

	type monthlySum struct {
		Month       string  `json:"month"`
		Duration    float64 `json:"duration"`
		Miles       float64 `json:"miles"`
		DaysSkipped float64 `json:"days_skipped"`
	}

	response := []monthlySum{}
	for k, v := range monthMap {
		response = append(response, monthlySum{Month: k, Duration: v["duration"], Miles: v["miles"], DaysSkipped: v["days_skipped"]})
	}
	c.JSON(http.StatusOK, response)
}
