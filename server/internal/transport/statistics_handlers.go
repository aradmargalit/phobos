package transport

import (
	"net/http"
	"server/internal/middleware"
	"server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Aggregator function to make it easy for the router builder to hit these all at once
func registerStatisticsHandlers(r *gin.Engine, svc service.PhobosAPI) {
	private := r.Group("/private")
	private.Use(middleware.AuthRequired)
	{
		private.GET("/statistics", makeGetUserStatisticsHandler(svc))
	}
}

func makeGetUserStatisticsHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		// Pull user out of context to figure out which activities to grab
		uid := c.GetInt("user")

		// Pull the user's timezone out of the request
		utcOffset, err := strconv.Atoi(c.Query("utc_offset"))
		if err != nil {
			panic(err)
		}

		stats, err := svc.GetUserStatistics(uid, utcOffset)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, *stats)
	}
}
