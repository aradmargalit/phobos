package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"server/internal/models"
	"server/internal/repository"
	"server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	// Autoloads .env file to supply environment variables
	_ "github.com/joho/godotenv/autoload"
)

const metersToMiles = 0.000621371
const metersToYards = 1.09361

// HandleStravaWebhookVerification responds with an OK so Strava knows we have a real server
func (svc *service) HandleStravaWebhookVerification(c *gin.Context) {
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

// HandleWebhookEvent will listen for webhook events and act accordingly
func (svc *service) HandleWebhookEvent(event models.StravaWebhookEvent) error {
	// Bail out if the subscription doesn't match Strava's ID, this may be a malicious POST!
	if ok := strconv.Itoa(event.SubscriptionID) == os.Getenv("STRAVA_WEBHOOK_SUB_ID"); !ok {
		return fmt.Errorf("unauthorized webhook POST! Tried to use %v as the subscription ID", event.SubscriptionID)
	}

	// Now we're off to actually process the event!
	// I tried doing this as a go routine, but it's useful to have Strava retry if we fail
	// Strava retries up to 3 times or until it gets a 200 in under 2s
	if event.ObjectType == "activity" {
		err := handleWebhookEvent(event, svc.db)
		if err != nil {
			return err
		}
	}
	return nil
}

func handleWebhookEvent(e models.StravaWebhookEvent, db repository.PhobosDB) (err error) {
	switch e.AspectType {
	case "create":
		err = fetchAndCreate(e.OwnerID, e.ObjectID, db)
	case "update":
		err = fetchAndUpdate(e.OwnerID, e.ObjectID, db)
	case "delete":
		err = eventDelete(e.OwnerID, e.ObjectID, db)
	}
	return
}

func fetchAndCreate(ownerID int, activityID int, db repository.PhobosDB) error {
	// 1. Fetch the activity from our application to check if we already have it
	dbActivity, err := db.GetActivityByStravaID(utils.MakeI64(activityID))
	if err == nil {
		// If we had no problems fetching this ID, we must already have it. No need to re-insert
		fmt.Printf("We already have this activity! Strava ID: %v | Phobos ID: %v\n", activityID, dbActivity.ID)
		return nil
	}

	// 2. If we -do- need to insert it, I need the verbose payload
	fetchedActivity, userID, err := fetchActivity(ownerID, activityID, db)
	if err != nil {
		return fmt.Errorf("failed to fetch activity ID %v from Strava: %v", activityID, err)
	}

	// 3. Convert the Strava Activity to a Phobos one and insert
	activity, err := convertStravaActivity(fetchedActivity, userID, db)
	if err != nil {
		return err
	}

	inserted, err := db.InsertActivity(activity)
	if err != nil {
		return fmt.Errorf("failed to insert activity ID %v from Strava: %v", activityID, err)
	}
	fmt.Printf("Successfully created activity! Strava ID: %v | Phobos ID: %v\n", activityID, inserted.ID)
	return nil
}

func fetchAndUpdate(ownerID int, activityID int, db repository.PhobosDB) error {
	// 1. We -always- need to fetch the full activity from Strava, since we always upsert
	fetchedActivity, userID, err := fetchActivity(ownerID, activityID, db)
	if err != nil {
		panic(err)
	}

	activity, err := convertStravaActivity(fetchedActivity, userID, db)
	if err != nil {
		return err
	}

	// Get the ID from our application
	dbActivity, err := db.GetActivityByStravaID(activity.StravaID)
	if err != nil {
		// We were unable to get this activity, so just insert it instead
		inserted, err := db.InsertActivity(activity)
		if err != nil {
			return fmt.Errorf("unable to insert activity from update action: %v", err)
		}
		fmt.Printf("Successfully created activity (from Strava Update)! Strava ID: %v | Phobos ID: %v\n", activityID, inserted.ID)
		return nil
	}

	// If we don't need to insert it, we'll just update it
	activity.ID = dbActivity.ID
	_, err = db.UpdateActivity(activity)
	if err != nil {
		return fmt.Errorf("unable to update activity: %v", err)
	}
	fmt.Printf("Successfully updated activity! Strava ID: %v | Phobos ID: %v\n", activityID, activity.ID)
	return nil
}

func eventDelete(ownerID int, activityID int, db repository.PhobosDB) error {
	// Get the ID from our application
	// TODO, there's no reason this shouldn't be a single DB call
	activity, err := db.GetActivityByStravaID(sql.NullInt64{Int64: int64(activityID), Valid: true})
	if err != nil {
		return fmt.Errorf("failed to fetch activity ID %v from Strava: %v", activityID, err)
	}

	err = db.DeleteActivityByID(activity.OwnerID, activity.ID)
	if err != nil {
		return fmt.Errorf("failed to delete activity ID %v from Strava: %v", activityID, err)
	}
	fmt.Printf("Successfully deleted activity! Strava ID: %v | Phobos ID: %v\n", activityID, activity.ID)
	return nil
}

func fetchActivity(ownerID int, activityID int, db repository.PhobosDB) (models.StravaActivity, int, error) {
	var fetchedActivity models.StravaActivity

	// We need to swap the ownerID for our user ID in order to generate an HTTP client with that user's token
	// so that we're authorized to get "private" activities for that user
	userID, err := db.GetUserIDByStravaID(ownerID)
	if err != nil {
		return models.StravaActivity{}, 0, err
	}

	client := getHTTPClient(userID, db)
	fmt.Println("Fetching: " + (baseURL + "/activities/" + strconv.Itoa(activityID)) + " from Strava...")

	resp, err := client.Get(baseURL + "/activities/" + strconv.Itoa(activityID))
	if err != nil {
		return models.StravaActivity{}, 0, err
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

	return fetchedActivity, userID, errors.New("could not fetch the activity from Strava")
}

func convertStravaActivity(fetchedActivity models.StravaActivity, userID int, db repository.PhobosDB) (*models.Activity, error) {
	// Convert the activity to our version of that activity
	typeID, err := db.GetActivityTypeIDByStravaType(fetchedActivity.Type)
	if err != nil {
		return nil, err
	}

	// Convert time to the correct format, using the provided timezone
	// Timezone is provided as (GMT-08:00) America/Los_Angeles, so split on the space to get the portion we need
	location, err := time.LoadLocation(strings.Split(fetchedActivity.Timezone, " ")[1])
	if err != nil {
		return nil, err
	}

	t, err := time.Parse("2006-01-02T15:04:05Z", fetchedActivity.StartDate)
	if err != nil {
		return nil, err
	}

	unit := "miles"
	// Convert Meters to Miles
	convertedDistance := fetchedActivity.Distance * metersToMiles
	if fetchedActivity.Type == "Swim" {
		unit = "yards"
		convertedDistance = fetchedActivity.Distance * metersToYards
	}

	return &models.Activity{
		Name:           fetchedActivity.Name,
		ActivityDate:   t.In(location).Format("2006-01-02"),
		ActivityTypeID: typeID,
		OwnerID:        userID,
		Duration:       (float64(fetchedActivity.MovingTime) / 60),
		Distance:       math.Floor(convertedDistance*100) / 100,
		Unit:           unit,
		StravaID:       sql.NullInt64{Int64: int64(fetchedActivity.ID), Valid: true},
	}, nil
}
