package utils

import (
	"server/models"
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

func seedActivityTypes(db *models.DB) {
	// First, delete everything that existed already
	db.DeleteAllActivityTypes()

	// Then, seed the database
	for _, name := range ActivityTypes {
		db.InsertActivityType(models.ActivityType{Name: name})
	}
}

// Seed will seed the database with all application data
func Seed(db *models.DB) {
	seedActivityTypes(db)
}
