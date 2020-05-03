package service

import (
	"server/internal/models"
)

// ActivityTypes is a list of possible activities
var ActivityTypes []string = []string{
	// Start Strava Types
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
	// Start Custom Types
	"Basketball",
	"Soccer",
	"Ultimate",
	"Tennis",
	"Volleyball",
}

func (svc *service) SeedActivityTypes() (err error) {
	// First, delete everything that existed already
	err = svc.db.DeleteAllActivityTypes()
	if err != nil {
		return
	}

	// Then, seed the database
	for _, name := range ActivityTypes {
		err = svc.db.InsertActivityType(models.ActivityType{Name: name})
		if err != nil {
			return
		}
	}
	return
}
