package transport

import (
	"net/http"
	"server/internal/middleware"
	"server/internal/models"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

// Aggregator function to make it easy for the router builder to hit these all at once

func registerStravaHandlers(r *gin.Engine, svc service.PhobosAPI) {
	r.GET("/public/strava/callback", makeStravaCallBackHandler(svc))
	r.GET("/public/strava/webhook", makeStravaWebookVerificationHandler(svc))
	r.POST("/public/strava/webhook", makeStravaWebhookCatcher(svc))

	strava := r.Group("/strava")
	strava.Use(middleware.AuthRequired)
	{
		strava.GET("/auth", makeStravaLoginHandler(svc))
		strava.GET("/deauth", makeStravaDeAuthHandler(svc))
	}
}

func makeStravaCallBackHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		svc.HandleStravaCallback(c)
	}
}

func makeStravaLoginHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		svc.HandleStravaLogin(c)
	}
}

func makeStravaDeAuthHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		uid := c.GetInt("user")
		err := svc.HandleStravaDeauthorization(uid)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, gin.H{"deleted": true})
	}
}

func makeStravaWebookVerificationHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		svc.HandleStravaWebhookVerification(c)
	}
}

func makeStravaWebhookCatcher(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		var event models.StravaWebhookEvent

		if err := c.ShouldBindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := svc.HandleWebhookEvent(event)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		// We're going to respond with an OK so the webhook doesn't retry if this succeeds
		c.JSON(http.StatusOK, gin.H{})
	}
}
