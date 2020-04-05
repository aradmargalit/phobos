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

	"github.com/tmdvs/Go-Emoji-Utils"
)

const metersToMiles = 0.000621371
const metersToYards = 1.09361

func main() {
	activityTypeMap := map[string]int{
		"Ride":            1,
		"Mountain Bike":   1,
		"Run":             2,
		"Swim":            3,
		"Walk":            4,
		"Hike":            5,
		"Rock Climbing":   18,
		"Snowboard":       21,
		"Surf":            25,
		"Workout":         28,
		"Yoga":            32,
		"Basketball":      33,
		"Soccer":          34,
		"Ultimate":        35,
		"Tennis":          36,
		"Volleyball":      37,
		"Elliptical":      11,
		"Kayaking":        15,
		"Canoeing":        8,
		"Weight Training": 28,
		"Virtual Ride":    26,
	}

	csvfile, err := os.Open("/Users/aradmargalit/Downloads/perry.csv")
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

		activityDate := record[1]
		activityType := record[3]
		duration := record[5]
		distance := record[6]
		name := record[2]
		stravaID := record[0]

		cleanName := emoji.RemoveAll(name)
		cleanName = strings.ReplaceAll(cleanName, `"`, "'")

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

		sql := fmt.Sprintf("INSERT INTO activities (name, activity_date, activity_type_id, owner_id, duration, distance, unit, strava_id) VALUES (\"%v\", \"%v\", \"%v\", 4, \"%v\",  \"%v\",  \"%v\",  \"%v\" );", cleanName, activityDate, activityTypeID, scaledDur, convertedDistance, unit, stravaID)
		fmt.Println(sql)
	}

}
