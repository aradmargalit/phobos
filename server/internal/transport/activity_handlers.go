package transport

import (
	"net/http"
	"server/internal/middleware"
	"server/internal/models"
	"server/internal/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Aggregator function to make it easy for the router builder to hit these all at once
func registerActivityHandlers(r *gin.Engine, svc service.PhobosAPI) {
	private := r.Group("/private")
	private.Use(middleware.AuthRequired)
	{
		private.GET("/users/current", makeCurrentUserHandler(svc))
		private.POST("/activities", makeAddActivityHandler(svc))
		private.GET("/activities", makeGetActivitiesHandler(svc))
		private.DELETE("/activities/:id", makeDeleteActivityHandler(svc))
		private.PUT("/activities/:id", makeUpdateActivityHandler(svc))
		private.GET("/activities/interval_summary", makeGetIntervalSummaryHandler(svc))
	}
}

// TODO move this to its own area
func makeCurrentUserHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		user := svc.GetCurrentUser(c)
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func makeAddActivityHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		var par models.PostActivityRequest
		if err := c.ShouldBindJSON(&par); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Add the owner ID to the activity
		uid := c.GetInt("user")

		record, err := svc.AddActivity(&par, uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		// Currently not consumed by the UI, but echo back the record
		c.JSON(http.StatusOK, *record)
	}
}

func makeGetActivitiesHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		uid := c.GetInt("user")
		activities, err := svc.GetActivities(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"activities": *activities, "timestamp": time.Now()})
	}
}

func makeDeleteActivityHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		// Pull user out of context to confirm it's safe to delete the activity
		uid := c.GetInt("user")

		// Convert the request param to an activity ID
		activityID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		err = svc.DeleteActivity(activityID, uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, "Successfully deleted activity: "+c.Param("id"))
	}
}

func makeUpdateActivityHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		var activity models.Activity
		if err := c.ShouldBindJSON(&activity); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		record, err := svc.UpdateActivity(&activity)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, *record)
	}
}

func makeGetIntervalSummaryHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		// Pull user out of context to figure out which activities to grab
		uid := c.GetInt("user")
		interval := c.Query("interval")
		// Pull the user's timezone out of the request
		utcOffset, err := strconv.Atoi(c.Query("utc_offset"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		activities, err := svc.GetIntervalSummary(uid, interval, utcOffset)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, *activities)
	}
}
