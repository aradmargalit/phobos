package controllers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

// ActivityTypes is a list of possible activities
var ActivityTypes []string = []string{
	"Ride",
	"Run",
	"Swim",
	"Walk",
	"Hike",
	"Alpine Ski",
	"Backcountry Ski",
	"Canoe",
	"Crossfit",
	"E-Bike Ride",
	"Elliptical",
	"Handcycle",
	"Ice Skate",
	"Inline Skate",
	"Kayak",
	"Kitesurf Session",
	"Nordic Ski",
	"Rock Climb",
	"Roller Ski",
	"Row",
	"Snowboard",
	"Snowshoe",
	"Stair Stepper",
	"Stand Up Paddle",
	"Surf",
	"Virtual Ride",
	"Virtual Run",
	"Weight Training",
	"Windsurf Session",
	"Wheelchair",
	"Workout",
	"Yoga",
}

func seedActivityTypes(db *models.DB) (err error) {
	// First, delete everything that existed already
	err = db.DeleteAllActivityTypes()
	if err != nil {
		return
	}

	// Then, seed the database
	for _, name := range ActivityTypes {
		err = db.InsertActivityType(models.ActivityType{Name: name})
		if err != nil {
			return
		}
	}
	return
}

// SeedHander will seed the database with all application data
func (e *Env) SeedHandler(c *gin.Context) {
	err := seedActivityTypes(e.DB)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.String(http.StatusOK, "Successfully seeded database.")
}