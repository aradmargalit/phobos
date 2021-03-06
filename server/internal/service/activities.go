package service

import (
	"errors"
	"fmt"
	"server/internal/models"
	"server/internal/responsetypes"
	"server/utils"
	"time"
)

// AddActivity adds a new activity to the database
func (svc *service) AddActivity(par *models.PostActivityRequest, uid int) (*models.Activity, error) {
	// Convert a PostActivityRequest to an Activity
	activity := models.Activity{
		Name:           par.Name,
		ActivityDate:   par.ActivityDate,
		ActivityTypeID: par.ActivityTypeID,
		OwnerID:        uid,
		Duration:       par.Duration,
		Meters:         utils.DistanceToMeters(par.Distance, par.Unit),
		Distance:       par.Distance,
		Unit:           par.Unit,
		HeartRate:      &par.HeartRate,
	}

	// MySQL doesn't like RFC3339 times, so convert it to YYYY-MM-DD
	d, err := time.Parse(time.RFC3339, activity.ActivityDate)
	if err != nil {
		return nil, err
	}

	activity.ActivityDate = d.Format("2006-01-02")

	activity.OwnerID = uid

	record, err := svc.db.InsertActivity(&activity)
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
func (svc *service) GetActivities(uid int) (*[]models.ActivityResponse, error) {
	a, err := svc.db.GetActivitiesByUser(uid)
	if err != nil {
		return nil, err
	}

	count := len(a)
	fmt.Printf("Found %v activities for user ID: %v...\n", count, uid)

	withIndices := make([]models.ActivityResponse, count)
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

	intervals := bucketIntoIntervals(a, interval, offset)

	c1 := make(chan map[string]float64, 1)
	c2 := make(chan map[string]float64, 1)
	c3 := make(chan map[string]float64, 1)

	go makeDurationMap(a, intervals, interval, c1)
	go makeDistanceMap(a, intervals, interval, c2)
	go makePercentageActiveMap(a, intervals, interval, offset, c3)

	durationMap := <-c1
	distanceMap := <-c2
	percentActiveMap := <-c3

	/* This is in the format:
	{
		"January 2019": {``
			"duration": 12.123,
			"distance": 12.231256,
			"percentage_active": 12.21
		}
	}
	*/

	sortIndex := 0
	response := []responsetypes.IntervalSum{}
	for _, itvl := range intervals {
		sortIndex++
		mSum := responsetypes.IntervalSum{Interval: itvl, Duration: durationMap[itvl], Miles: distanceMap[itvl], PercentageActive: percentActiveMap[itvl], SortIndex: sortIndex}
		response = append(response, mSum)
	}

	return &response, nil
}

func bucketIntoIntervals(activities []models.ActivityResponse, itvl string, offset int) []string {
	intervals := []string{}

	// First create a bucket for the current period
	dur, _ := time.ParseDuration(fmt.Sprintf("%vh", offset))
	now := time.Now().UTC().Add(-dur)
	nowString := timeToIntervalString(now, itvl)

	intervals = append(intervals, nowString)

	prev := nowString
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

func makeDurationMap(activities []models.ActivityResponse, intervals []string, itvl string, c chan map[string]float64) {
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

func makeDistanceMap(activities []models.ActivityResponse, intervals []string, itvl string, c chan map[string]float64) {
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

func makePercentageActiveMap(activities []models.ActivityResponse, intervals []string, itvl string, offset int, c chan map[string]float64) {
	// Used to keep track of an intervals "hits" (days worked out) and totals
	type hitCounter struct {
		hits  int
		total int
	}

	// ALGORITHM:
	// We want to go through every day since the user's first activity and, if they worked out on that date add to a count for that interval
	// Regardless of whether or not they worked out, mark it in a "total" field, to later define a percentage
	// At the end, go through the map of interval to hits and total and write the result percentage to the channel

	firstActivityDate, _ := time.Parse("2006-01-02 15:04:05", activities[len(activities)-1].ActivityDate)

	// Make a Set of all activity dates
	// Go doesn't have sets, but we can use a map["some type"]bool
	activityDateSet := map[time.Time]bool{}
	for _, activity := range activities {
		t, _ := time.Parse("2006-01-02 15:04:05", activity.ActivityDate)
		activityDateSet[utils.RoundTimeToDay(t)] = true
	}

	intervalToHitTotalMap := map[string]hitCounter{}

	// To make things easy, initialize the map with 0s
	for _, itvl := range intervals {
		intervalToHitTotalMap[itvl] = struct {
			hits  int
			total int
		}{hits: 0, total: 0}
	}

	// Go through each date and add to the total, conditionally add to the "hit"
	// Need to use UTC for time.Now() since the server is deployed in UTC
	dur, _ := time.ParseDuration(fmt.Sprintf("%vh", offset))
	now := time.Now().UTC().Add(-dur)

	// For each date from now to the first activity, working backwards,
	// check if any activities match that date
	dayBeforeFirst := firstActivityDate.AddDate(0, 0, -1)
	for d := now; !utils.DateEqual(d, dayBeforeFirst); d = d.AddDate(0, 0, -1) {
		dateToCheck := utils.RoundTimeToDay(d)
		intervalFromDate := timeToIntervalString(dateToCheck, itvl)
		currHits := intervalToHitTotalMap[intervalFromDate].hits
		currTotal := intervalToHitTotalMap[intervalFromDate].total

		if _, ok := activityDateSet[dateToCheck]; ok {
			currHits++
		}

		// Always add to the total, even if there's no hit
		currTotal++
		intervalToHitTotalMap[intervalFromDate] = hitCounter{hits: currHits, total: currTotal}
	}

	// Now that we have counts, convert those to percentages
	percentages := map[string]float64{}
	for interval, counts := range intervalToHitTotalMap {
		percentage := (float64(counts.hits) / float64(counts.total))
		percentages[interval] = utils.FloatTwoDecimals(percentage * 100)
	}

	c <- percentages
}

func timeToIntervalString(t time.Time, itvl string) string {
	switch itvl {
	case "year":
		return fmt.Sprintf("%v", t.Year())
	case "month":
		return fmt.Sprintf("%v %v", t.Month(), t.Year())
	case "week":
		year, week := t.ISOWeek()
		month := t.Month()
		return fmt.Sprintf("%v, week %v (%v)", year, week, month.String()[:3])
	}
	// Theoretically this could happen, but we're bouncing requests that this switch wouldn't catch
	return ""
}
