package controllers

import (
	"fmt"
	"net/http"

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

	fmt.Println(event)
	c.JSON(http.StatusOK, gin.H{})
}