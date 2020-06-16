package service

import (
	"server/internal/models"
)

// ActivityTypes is a list of possible activities, mapped to their Strava names
var ActivityTypes map[string]string = map[string]string{
	// Start Strava Types
	"Ride":             "Ride",
	"Run":              "Run",
	"Swim":             "Swim",
	"Walk":             "Walk",
	"Hike":             "Hike",
	"Alpine Ski":       "AlpineSki",
	"Backcountry Ski":  "BackcountrySki",
	"Canoe":            "Canoeing",
	"Crossfit":         "Crossfit",
	"E-Bike Ride":      "EBikeRide",
	"Elliptical":       "Elliptical",
	"Handcycle":        "Handcycle",
	"Ice Skate":        "IceSkate",
	"Inline Skate":     "InlineSkate",
	"Kayak":            "Kayaking",
	"Kitesurf Session": "Kitesurf",
	"Nordic Ski":       "NordicSki",
	"Rock Climb":       "RockClimbing",
	"Roller Ski":       "RollerSki",
	"Row":              "Rowing",
	"Snowboard":        "Snowboard",
	"Snowshoe":         "Snowshoe",
	"Stair Stepper":    "StairStepper",
	"Stand Up Paddle":  "StandUpPaddling",
	"Surf":             "Surfing",
	"Virtual Ride":     "VirtualRide",
	"Virtual Run":      "VirtualRun",
	"Weight Training":  "WeightTraining",
	"Windsurf Session": "Windsurf",
	"Wheelchair":       "Wheelchair",
	"Workout":          "Workout",
	"Yoga":             "Yoga",
	"Soccer":           "Soccer",
	"Golf":             "Golf",
	"Sail":             "Sail",
	"Skateboard":       "Skateboard",
	// Start Custom Types
	"Basketball": "",
	"Ultimate":   "",
	"Tennis":     "",
	"Volleyball": "",
}

func (svc *service) SeedActivityTypes() (err error) {
	// First, delete everything that existed already
	err = svc.db.DeleteAllActivityTypes()
	if err != nil {
		return
	}

	// Then, seed the database
	for name, stravaName := range ActivityTypes {
		err = svc.db.InsertActivityType(models.ActivityType{Name: name, StravaName: stravaName})
		if err != nil {
			return
		}
	}
	return
}
