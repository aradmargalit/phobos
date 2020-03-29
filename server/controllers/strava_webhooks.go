package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type stravaWebhookEvent struct {
	ObjectType string `json:"object_type"`
	ObjectID   int    `json:"object_id"`
	AspectType string `json:"aspect_type"`
	OwnerID    int    `json:"owner_id"`
	EventTime  int    `json:"event_time"`
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
		fmt.Println(e.ObjectID)
		fetchAndCreate(e.OwnerID, e.ObjectID, db)
	case "update":
		fmt.Println("Caught an update")
	case "delete":
		fmt.Println("Caught a delete!")
	}
}

func fetchAndCreate(ownerID int, activityID int, db *models.DB) {
	fetchedActivity, err := fetchActivity(ownerID, activityID, db)
	if err != nil {
		panic(err)
	}

	// Convert the activity to our version of that activity
	// TODO pickup here
	// activity := models.Activity{
	// 	Name: fetchedActivity.Name,
	// 	ActivityDate: fetchedActivity.StartDate,
	// }
}

type stravaActivity struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Distance    float64    `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type        string `json:"type"`
	StartDate   string `json:"start_date"`
	Timezone    string `json:"timezone"`
}

func fetchActivity(ownerID int, activityID int, db *models.DB) (stravaActivity, error) {
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
	fmt.Printf("%v\n", resp)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fetchedActivity, err
		}
		err = json.Unmarshal(bodyBytes, &fetchedActivity)
		if err != nil {
			panic(err)
		}
		return fetchedActivity, nil
	}

	return fetchedActivity, errors.New("Could not fetch the activity from Strava")
}
