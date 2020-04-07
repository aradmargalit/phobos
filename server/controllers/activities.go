package controllers

import (
	"fmt"
	"net/http"
	models "server/models"
	responsetypes "server/response_types"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// AddActivityHandler adds a new activity to the database
func (e *Env) AddActivityHandler(c *gin.Context) {
	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Currently not consumed by the UI, but echo back the record
	c.JSON(http.StatusOK, record)
}

// UpdateActivityHandler adds a new activity to the database
func (e *Env) UpdateActivityHandler(c *gin.Context) {
	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// MySQL doesn't like RFC3339 times, so convert it to YYYY-MM-DD
	d, err := time.Parse(time.RFC3339, activity.ActivityDate)
	activity.ActivityDate = d.Format("2006-01-02")

	record, err := e.DB.UpdateActivity(activity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		panic(err)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = e.DB.DeleteActivityByID(uid, activityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		panic(err)
	}

	monthMap := map[string]map[string]float64{}

	for _, activity := range a {
		m, _ := time.Parse("2006-01-02 15:04:05", activity.ActivityDate)
		month := fmt.Sprintf("%v %v", m.Month(), m.Year())
		_, ok := monthMap[month]
		if !ok {
			monthMap[month] = map[string]float64{"duration": 0, "miles": 0}
		}
		monthMap[month]["duration"] += activity.Duration
		if activity.Unit == "miles" {
			monthMap[month]["miles"] += activity.Distance
		}
	}

	type monthlySum struct {
		Month    string  `json:"month"`
		Duration float64 `json:"duration"`
		Miles    float64 `json:"miles"`
	}

	response := []monthlySum{}
	for k, v := range monthMap {
		response = append(response, monthlySum{Month: k, Duration: v["duration"], Miles: v["miles"]})
	}
	c.JSON(http.StatusOK, response)
}
