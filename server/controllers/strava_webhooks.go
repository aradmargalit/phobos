package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"server/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const metersToMiles = 0.000621371
const metersToYards = 1.09361

type stravaWebhookEvent struct {
	ObjectType string `json:"object_type"`
	ObjectID   int    `json:"object_id"`
	AspectType string `json:"aspect_type"`
	OwnerID    int    `json:"owner_id"`
	EventTime  int    `json:"event_time"`
}

type stravaActivity struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Distance    float64 `json:"distance"`
	ElapsedTime int     `json:"elapsed_time"`
	Type        string  `json:"type"`
	StartDate   string  `json:"start_date"`
	Timezone    string  `json:"timezone"`
}

// StravaWebookVerificationHandler responds with an OK so Strava knows we have a real server
func (e *Env) StravaWebookVerificationHandler(c *gin.Context) {
	// https://developers.strava.com/docs/webhooks/
	/* 	Your callback address must respond within two seconds to the GET request from Strava’s subscription service.
	The response should indicate status code 200 and should echo the hub.challenge field in the response body as application/json content type:
	{ “hub.challenge”:”15f7d1a91c1f40f8a748fd134752feb3” }
	*/

	// We should only need to run this once in production, and we should remember what our ID is
	challenge := c.Query("hub.challenge")
	c.JSON(http.StatusOK, gin.H{
		"hub.challenge": challenge,
	})
}

// StravaWebHookCatcher will listen for webhook events and act accordingly
func (e *Env) StravaWebHookCatcher(c *gin.Context) {
	var event stravaWebhookEvent

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We're going to respond immediately with an OK so the webhook doesn't retry
	c.JSON(http.StatusOK, gin.H{})

	// Now we're off to actually process the event!
	// Do this as a gorountine to free up the main webserver, and ignore non-activity updates for now
	if event.ObjectType == "activity" {
		go handleWebhookEvent(event, e.DB)
	}
}

func handleWebhookEvent(e stravaWebhookEvent, db *models.DB) {
	switch e.AspectType {
	case "create":
		fetchAndCreate(e.OwnerID, e.ObjectID, db)
	case "update":
		fetchAndUpdate(e.OwnerID, e.ObjectID, db)
	case "delete":
		eventDelete(e.OwnerID, e.ObjectID, db)
	}
}

func fetchAndCreate(ownerID int, activityID int, db *models.DB) {
	fetchedActivity, userID, err := fetchActivity(ownerID, activityID, db)
	if err != nil {
		panic(err)
	}

	activity := convertStravaActivity(fetchedActivity, userID, db)
	_, err = db.InsertActivity(activity)
	if err != nil {
		panic(err)
	}
}

func fetchAndUpdate(ownerID int, activityID int, db *models.DB) {
	fetchedActivity, userID, err := fetchActivity(ownerID, activityID, db)
	if err != nil {
		panic(err)
	}

	activity := convertStravaActivity(fetchedActivity, userID, db)

	// Get the ID from our application
	id, err := db.GetActivityIDByStravaID(activity.StravaID)
	activity.ID = id

	_, err = db.UpdateActivity(activity)
	if err != nil {
		panic(err)
	}
}

func eventDelete(ownerID int, activityID int, db *models.DB) {
	// Get the ID from our application
	id, err := db.GetActivityIDByStravaID(sql.NullInt64{Int64: int64(activityID), Valid: true})
	if err != nil {
		panic(err)
	}

	err = db.DeleteActivityByID(ownerID, id)
	if err != nil {
		panic(err)
	}
}

func fetchActivity(ownerID int, activityID int, db *models.DB) (stravaActivity, int, error) {
	var fetchedActivity stravaActivity

	// We need to swap the ownerID for our user ID
	userID, err := db.GetUserIDByStravaID(ownerID)
	if err != nil {
		panic(err)
	}

	client := getHTTPClient(userID, db)
	fmt.Println("Fetching: " + (baseURL + "/activities/" + strconv.Itoa(activityID)) + " from Strava...")
	resp, err := client.Get(baseURL + "/activities/" + strconv.Itoa(activityID))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fetchedActivity, userID, err
		}
		err = json.Unmarshal(bodyBytes, &fetchedActivity)
		if err != nil {
			panic(err)
		}
		return fetchedActivity, userID, nil
	}

	return fetchedActivity, userID, errors.New("Could not fetch the activity from Strava")
}

func convertStravaActivity(fetchedActivity stravaActivity, userID int, db *models.DB) models.Activity {
	// Convert the activity to our version of that activity
	typeID, err := db.GetActivityTypeIDByStravaType(fetchedActivity.Type)
	if err != nil {
		panic(err)
	}

	// Convert time to the correct format, using the provided timezone
	// Timezone is provided as (GMT-08:00) America/Los_Angeles, so split on the space to get the portion we need
	location, err := time.LoadLocation(strings.Split(fetchedActivity.Timezone, " ")[1])
	t, err := time.Parse("2006-01-02T15:04:05Z", fetchedActivity.StartDate)

	unit := "miles"
	// Convert Meters to Miles
	convertedDistance := fetchedActivity.Distance * metersToMiles
	if fetchedActivity.Type == "Swim" {
		unit = "yards"
		convertedDistance = fetchedActivity.Distance * metersToYards
	}

	return models.Activity{
		Name:           fetchedActivity.Name,
		ActivityDate:   t.In(location).Format("2006-01-02"),
		ActivityTypeID: typeID,
		OwnerID:        userID,
		Duration:       (float64(fetchedActivity.ElapsedTime) / 60),
		Distance:       math.Floor(convertedDistance*100) / 100,
		Unit:           unit,
		StravaID:       sql.NullInt64{Int64: int64(fetchedActivity.ID), Valid: true},
	}
}
