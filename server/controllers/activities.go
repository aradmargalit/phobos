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

	months := bucketIntoMonths(a)

	c1 := make(chan map[string]float64, 1)
	c2 := make(chan map[string]float64, 1)
	c3 := make(chan map[string]float64, 1)

	go makeDurationMap(a, months, c1)
	go makeDistanceMap(a, months, c2)
	go makeSkippedMap(a, months, c3)

	durationMap := <-c1
	distanceMap := <-c2
	skippedMap := <-c3

	/* This is in the format:
	{
		"January 2019": {
			"duration": 12.123,
			"distance": 12.231256,
			"daysSkipped": 12
		}
	}
	*/

	response := []monthlySum{}
	for _, month := range months {
		mSum := monthlySum{Month: month, Duration: durationMap[month], Miles: distanceMap[month], DaysSkipped: skippedMap[month]}
		response = append(response, mSum)
	}

	c.JSON(http.StatusOK, response)
}

type monthlySum struct {
	Month       string  `json:"month"`
	Duration    float64 `json:"duration"`
	Miles       float64 `json:"miles"`
	DaysSkipped float64 `json:"days_skipped"`
}

func bucketIntoMonths(activities []responsetypes.ActivityResponse) []string {
	months := []string{}
	var previousMonth string
	for _, a := range activities {
		t, _ := time.Parse("2006-01-02 15:04:05", a.ActivityDate)

		// Format is "January 2020"
		month := fmt.Sprintf("%v %v", t.Month(), t.Year())
		if month != previousMonth {
			months = append(months, month)
			previousMonth = month
		}

	}
	return months
}

func makeDurationMap(activities []responsetypes.ActivityResponse, months []string, c chan map[string]float64) {
	durationMap := map[string]float64{}
	for _, month := range months {
		durationMap[month] = 0
	}

	for _, a := range activities {
		t, _ := time.Parse("2006-01-02 15:04:05", a.ActivityDate)

		// Format is "January 2020"
		month := fmt.Sprintf("%v %v", t.Month(), t.Year())
		durationMap[month] += a.Duration
	}
	c <- durationMap
}

func makeDistanceMap(activities []responsetypes.ActivityResponse, months []string, c chan map[string]float64) {
	distanceMap := map[string]float64{}
	for _, month := range months {
		distanceMap[month] = 0
	}
	for _, a := range activities {
		t, _ := time.Parse("2006-01-02 15:04:05", a.ActivityDate)

		// Format is "January 2020"
		month := fmt.Sprintf("%v %v", t.Month(), t.Year())
		if a.Unit == "miles" {
			distanceMap[month] += a.Distance
		}
	}
	c <- distanceMap
}

func makeSkippedMap(activities []responsetypes.ActivityResponse, months []string, c chan map[string]float64) {
	// In order to figure out which days have no activities, we start with an array with every day from
	// their first activity until now
	firstActivityDate, _ := time.Parse("2006-01-02 15:04:05", activities[len(activities)-1].ActivityDate)

	// Make a map of days to whether or not they were skipped
	skippedMap := map[time.Time]bool{}

	// Start by assuming each date is skipped, we'll invalidate that assumption as we go
	for d := firstActivityDate; !utils.DateEqual(d, time.Now().AddDate(0, 0, 1)); d = d.AddDate(0, 0, 1) {
		skippedMap[utils.RoundTimeToDay(d)] = true
	}

	for _, a := range activities {
		t, _ := time.Parse("2006-01-02 15:04:05", a.ActivityDate)

		// Go into our date map and mark the date as unskipped
		skippedMap[utils.RoundTimeToDay(t)] = false
	}

	// After going through every activity, we need to summarize the skipped days
	monthlySkips := map[string]float64{}

	for _, month := range months {
		monthlySkips[month] = 0
		m := strings.Split(month, " ")[0]
		y := strings.Split(month, " ")[1]

		// For each skipped activity, see if it matches, and if so, add it to the tally
		for sA, wasSkipped := range skippedMap {
			if wasSkipped && sA.Month().String() == m && strconv.Itoa(sA.Year()) == y {

				monthlySkips[month]++
			}
		}
	}

	c <- monthlySkips
}
