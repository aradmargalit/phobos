package transport

import (
	"net/http"
	"server/internal/middleware"
	"server/internal/models"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

// Aggregator function to make it easy for the router builder to hit these all at once
func registerGoalHandlers(r *gin.Engine, svc service.PhobosAPI) {
	private := r.Group("/private")
	private.Use(middleware.AuthRequired)
	{
		private.GET("/goals", makeGetGoalsHandler(svc))
		private.POST("/goals", makeAddGoalHandler(svc))
		// private.DELETE("/goals/:id", makeDeleteGoalHandler(svc))
	}
}

func makeGetGoalsHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		// Pull user out of context to figure out which goals to grab
		uid := c.GetInt("user")
		goals, err := svc.GetGoals(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, *goals)
	}
}

func makeAddGoalHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		var goal models.Goal
		if err := c.ShouldBindJSON(&goal); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Add the owner ID to the activituy
		uid := c.GetInt("user")

		record, err := svc.AddGoal(uid, &goal)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, *record)
	}
}

// func makeDeleteGoalHandler(svc service.PhobosAPI) func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		// Pull user out of context to confirm it's safe to delete the activity
// 		uid := c.GetInt("user")

// 		quickAddID, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.AbortWithError(http.StatusBadRequest, errors.New("quick Add ID must be an int"))
// 			return
// 		}

// 		err = svc.DeleteQuickAdd(uid, quickAddID)
// 		if err != nil {
// 			c.AbortWithError(http.StatusInternalServerError, err)
// 			return
// 		}

// 		c.JSON(http.StatusOK, "Successfully deleted quick-add: "+c.Param("id"))
// 	}
// }
