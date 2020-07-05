package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const metersToMiles = 0.000621371
const metersToYards = 1.09361

func main() {

	// This sucks, and should really be using a database lookup table
	activityTypeMap := map[string]int{
		"Sail":             1,
		"Ultimate":         2,
		"Alpine Ski":       3,
		"Surf":             4,
		"Soccer":           5,
		"Golf":             6,
		"Walk":             7,
		"Inline Skate":     8,
		"Stand Up Paddle":  9,
		"Basketball":       10,
		"Kitesurf Session": 11,
		"Snowboard":        12,
		"Volleyball":       13,
		"Run":              14,
		"Hike":             15,
		"Stair Stepper":    16,
		"Row":              17,
		"Snowshoe":         18,
		"Virtual Ride":     19,
		"Wheelchair":       20,
		"Ride":             21,
		"Canoe":            22,
		"Elliptical":       23,
		"Handcycle":        24,
		"Skateboard":       25,
		"Rock Climb":       26,
		"Swim":             27,
		"Crossfit":         28,
		"Ice Skate":        29,
		"Nordic Ski":       30,
		"Backcountry Ski":  31,
		"Virtual Run":      32,
		"Weight Training":  33,
		"Workout":          34,
		"Yoga":             35,
		"Tennis":           36,
		"E-Bike Ride":      37,
		"Kayak":            38,
		"Roller Ski":       39,
		"Windsurf Session": 40,
	}

	// This sucks, and should be a CLI argument
	csvfile, err := os.Open("/Users/aradmargalit/Desktop/activities_jo.csv")
	if err != nil {
		log.Fatalln("Couldn't open the input", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	count := 0

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if count == 0 {
			count++
			continue
		}

		// This sucks, and the plucking from the CSV should be easier to read
		activityDate := record[1]
		activityType := record[3]
		duration := record[5]
		distance := record[6]
		name := record[2]
		stravaID := record[0]
		hr := record[29]

		if hr == "" {
			hr = "NULL"
		}

		activityTypeID, ok := activityTypeMap[activityType]
		if !ok {
			panic(activityType)
		}

		if distance == "" {
			distance = "0"
		}

		//Replace quotes in the name with single quotes
		name = strings.ReplaceAll(name, "\"", "'")

		// Downscale duration
		intDur, err := strconv.Atoi(duration)
		scaledDur := intDur / 60

		// MySQL doesn't like RFC3339 times, so convert it to YYYY-MM-DD
		d, err := time.Parse("Jan 2, 2006, 3:04:05 PM", activityDate)
		if err != nil {
			panic(err)
		}

		floatDistace, _ := strconv.ParseFloat(distance, 64)
		floatDistace = floatDistace * 1000
		// Convert Meters to Miles
		unit := "miles"
		convertedDistance := floatDistace * metersToMiles
		if activityType == "Swim" {
			unit = "yards"
			convertedDistance = floatDistace * metersToYards
		}
		convertedDistance = math.Floor(convertedDistance*100) / 100

		activityDate = d.Format("2006-01-02")
		// This sucks, and shouldn't hardcode the id of the athlete
		sql := fmt.Sprintf("INSERT INTO activities (name, activity_date, activity_type_id, owner_id, duration, distance, unit, heart_rate, meters, strava_id) VALUES (\"%v\", \"%v\", \"%v\", 9, \"%v\",  \"%v\",  \"%v\",  %v, %v, %v);", name, activityDate, activityTypeID, scaledDur, convertedDistance, unit, hr, floatDistace, stravaID)
		fmt.Println(sql)
	}

}
