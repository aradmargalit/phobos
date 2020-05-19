package service

import (
	"errors"
	"fmt"
	"server/internal/models"
	"server/internal/responsetypes"
	"server/utils"
	"strconv"
	"strings"
	"time"
)

// AddActivity adds a new activity to the database
func (svc *service) AddActivity(activity *models.Activity, uid int) (*models.Activity, error) {
	// MySQL doesn't like RFC3339 times, so convert it to YYYY-MM-DD
	d, err := time.Parse(time.RFC3339, activity.ActivityDate)
	if err != nil {
		return nil, err
	}

	activity.ActivityDate = d.Format("2006-01-02")

	activity.OwnerID = uid
	fmt.Println(activity)

	record, err := svc.db.InsertActivity(activity)
	if err != nil {
		return nil, err
	}

	return record, nil
}

// UpdateActivity adds a new activity to the database
func (svc *service) UpdateActivity(activity *models.Activity) (*models.Activity, error) {

	// MySQL doesn't like RFC3339 times, so convert it to YYYY-MM-DD
	d, err := time.Parse(time.RFC3339, activity.ActivityDate)
	if err != nil {
		return nil, err
	}

	activity.ActivityDate = d.Format("2006-01-02")

	record, err := svc.db.UpdateActivity(activity)
	if err != nil {
		return nil, err
	}

	return record, nil
}

// GetActivities returns all the user's activities
func (svc *service) GetActivities(uid int) (*[]responsetypes.Activity, error) {
	a, err := svc.db.GetActivitiesByUser(uid)
	if err != nil {
		return nil, err
	}

	count := len(a)
	fmt.Printf("Found %v activities for user ID: %v...\n", count, uid)

	withIndices := make([]responsetypes.Activity, count)
	// No smart way to do this, add an decreasing logical index to each for the frontend's benefit
	// I also want to represent the date in an fast-to-sort fashion, so doing that here
	for idx, activity := range a {
		activity.LogicalIndex = count - idx

		// Convert date to seconds since epoch, much faster to sort ints in the UI than cast to Date objects
		t, _ := time.Parse("2006-01-02 15:04:05", activity.ActivityDate)
		activity.Epoch = t.Unix()
		withIndices[idx] = activity
	}

	return &withIndices, nil
}

// DeleteActivity returns all the user's activities
func (svc *service) DeleteActivity(activityID int, uid int) error {
	return svc.db.DeleteActivityByID(uid, activityID)
}

// GetIntervalSummary returns the user's aggregate data for the given interval
func (svc *service) GetIntervalSummary(uid int, interval string, offset int) (*[]responsetypes.IntervalSum, error) {
	// Validate the interval
	if interval != "week" && interval != "month" && interval != "year" {
		return nil, errors.New("interval must be week, month, or year")
	}

	a, err := svc.db.GetActivitiesByUser(uid)
	if err != nil {
		return nil, err
	}

	if len(a) < 1 {
		return &[]responsetypes.IntervalSum{}, nil
	}

	intervals := bucketIntoIntervals(a, interval)

	c1 := make(chan map[string]float64, 1)
	c2 := make(chan map[string]float64, 1)
	c3 := make(chan map[string]float64, 1)

	go makeDurationMap(a, intervals, interval, c1)
	go makeDistanceMap(a, intervals, interval, c2)
	go makeSkippedMap(a, intervals, interval, offset, c3)

	durationMap := <-c1
	distanceMap := <-c2
	skippedMap := <-c3

	/* This is in the format:
	{
		"January 2019": {``
			"duration": 12.123,
			"distance": 12.231256,
			"daysSkipped": 12
		}
	}
	*/

	response := []responsetypes.IntervalSum{}
	for _, itvl := range intervals {
		mSum := responsetypes.IntervalSum{Interval: itvl, Duration: durationMap[itvl], Miles: distanceMap[itvl], DaysSkipped: skippedMap[itvl]}
		response = append(response, mSum)
	}

	return &response, nil
}

func bucketIntoIntervals(activities []responsetypes.Activity, itvl string) []string {
	intervals := []string{}
	var prev string
	for _, a := range activities {
		t, _ := time.Parse("2006-01-02 15:04:05", a.ActivityDate)

		activityInterval := timeToIntervalString(t, itvl)
		if activityInterval != prev {
			intervals = append(intervals, activityInterval)
			prev = activityInterval
		}

	}
	return intervals
}

func makeDurationMap(activities []responsetypes.Activity, intervals []string, itvl string, c chan map[string]float64) {
	durationMap := map[string]float64{}
	for _, interval := range intervals {
		durationMap[interval] = 0
	}

	for _, a := range activities {
		t, _ := time.Parse("2006-01-02 15:04:05", a.ActivityDate)

		activityInterval := timeToIntervalString(t, itvl)
		durationMap[activityInterval] += a.Duration
	}

	c <- durationMap
}

func makeDistanceMap(activities []responsetypes.Activity, intervals []string, itvl string, c chan map[string]float64) {
	distanceMap := map[string]float64{}
	for _, interval := range intervals {
		distanceMap[interval] = 0
	}
	for _, a := range activities {
		t, _ := time.Parse("2006-01-02 15:04:05", a.ActivityDate)

		activityInterval := timeToIntervalString(t, itvl)

		if a.Unit == "miles" {
			distanceMap[activityInterval] += a.Distance
		}
	}
	c <- distanceMap
}

func makeSkippedMap(activities []responsetypes.Activity, intervals []string, itvl string, offset int, c chan map[string]float64) {
	// In order to figure out which days have no activities, we start with an array with every day from
	// their first activity until now
	firstActivityDate, _ := time.Parse("2006-01-02 15:04:05", activities[len(activities)-1].ActivityDate)

	// Make a map of days to whether or not they were skipped
	skippedMap := map[time.Time]bool{}

	// Start by assuming each date is skipped, we'll invalidate that assumption as we go
	// Need to use UTC for time.Now() since the server is deployed in UTC
	dur, _ := time.ParseDuration(fmt.Sprintf("%vh", offset))
	now := time.Now().UTC().Add(-dur)
	for d := firstActivityDate; !utils.DateEqual(d, now.AddDate(0, 0, 1)); d = d.AddDate(0, 0, 1) {
		skippedMap[utils.RoundTimeToDay(d)] = true
	}

	for _, a := range activities {
		t, _ := time.Parse("2006-01-02 15:04:05", a.ActivityDate)

		// Go into our date map and mark the date as unskipped
		skippedMap[utils.RoundTimeToDay(t)] = false
	}

	// After going through every activity, we need to summarize the skipped days
	groupedSkips := map[string]float64{}

	for _, interval := range intervals {
		groupedSkips[interval] = 0

		// For each skipped activity, see if it matches, and if so, add it to the tally
		for t, wasSkipped := range skippedMap {
			if wasSkipped && matchesIntervalDate(t, interval, itvl) {
				groupedSkips[interval]++
			}
		}
	}

	c <- groupedSkips
}

func timeToIntervalString(t time.Time, itvl string) string {
	switch itvl {
	case "year":
		return fmt.Sprintf("%v", t.Year())
	case "month":
		return fmt.Sprintf("%v %v", t.Month(), t.Year())
	case "week":
		year, week := t.ISOWeek()
		return fmt.Sprintf("%v, week %v", year, week)
	}
	// Theoretically this could happen, but we're bouncing requests that this switch wouldn't catch
	return ""
}

func matchesIntervalDate(t time.Time, interval string, itvl string) bool {
	switch itvl {
	case "year":
		return strconv.Itoa(t.Year()) == interval
	case "month":
		m := strings.Split(interval, " ")[0]
		y := strings.Split(interval, " ")[1]
		return t.Month().String() == m && strconv.Itoa(t.Year()) == y
	case "week":
		year, week := t.ISOWeek()
		return fmt.Sprintf("%v, week %v", year, week) == interval

	}
	return false
}
