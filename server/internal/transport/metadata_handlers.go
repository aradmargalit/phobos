package transport

import (
	"net/http"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

// Aggregator function to make it easy for the router builder to hit these all at once
func registerMetadataHandlers(r *gin.Engine, svc service.PhobosAPI) {
	// Called by the UI when the user clicks the "Login with Google Button"
	metadata := r.Group("/metadata")
	{
		metadata.GET("/activity_types", makeGetActivityTypesHandler(svc))
	}
}

func makeGetActivityTypesHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		at, err := svc.GetActivityTypes()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"activity_types": *at})
	}
}
