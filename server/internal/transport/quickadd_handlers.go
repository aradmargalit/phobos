package transport

import (
	"errors"
	"net/http"
	"server/internal/middleware"
	"server/internal/models"
	"server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Aggregator function to make it easy for the router builder to hit these all at once
func registerQuickAddHandlers(r *gin.Engine, svc service.PhobosAPI) {
	private := r.Group("/private")
	private.Use(middleware.AuthRequired)
	{
		private.GET("/quick_adds", makeGetQuickAddsHandler(svc))
		private.POST("/quick_adds", makeAddQuickAddsHandler(svc))
		private.DELETE("/quick_adds/:id", makeDeleteQuickAddHandler(svc))
	}
}

func makeGetQuickAddsHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		// Pull user out of context to figure out which quick-adds to grab
		uid := c.GetInt("user")
		qas, err := svc.GetQuickAdds(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		c.JSON(http.StatusOK, *qas)
	}
}

func makeAddQuickAddsHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		var quickAdd models.QuickAdd
		if err := c.ShouldBindJSON(&quickAdd); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Add the owner ID to the activituy
		uid := c.GetInt("user")

		record, err := svc.AddQuickAdd(uid, &quickAdd)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, *record)
	}
}

func makeDeleteQuickAddHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		// Pull user out of context to confirm it's safe to delete the activity
		uid := c.GetInt("user")

		quickAddID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("quick Add ID must be an int"))
			return
		}

		err = svc.DeleteQuickAdd(uid, quickAddID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "Successfully deleted quick-add: "+c.Param("id"))
	}
}
