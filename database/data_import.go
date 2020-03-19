package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	activityTypeMap := map[string]int{
		"Bike": 1,
		"Mountain Bike": 1,
		"Run": 2,
		"Swim": 3,
		"Walk": 4,
		"Hike": 5,
		"Climbing": 18,
		"Snowboard": 21,
		"Surf": 25,
		"Lift": 28,
		"Yoga": 32,
		"Basketball": 33,
		"Soccer": 34,
		"Ultimate": 35,		
		"Tennis": 36,
		"Volleyball": 37,
	}

	csvfile, err := os.Open("/Users/aradmargalit/Downloads/we.csv")
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

		activityDate := record[0]
		activityType := record[5]
		duration := record[6]
		distance := record[7]

		if activityType != "None" {
			if activityType == "Run - FS" {
				activityType = "Run"
			}
	
			activityTypeID, ok := activityTypeMap[activityType]
			if !ok {
				panic(activityType)
			}
	
			if distance == "" {
				distance = "0"
			}
	
			// MySQL doesn't like RFC3339 times, so convert it to YYYY-MM-DD
			d, err := time.Parse("1/2/06", activityDate)
			if err != nil {
				panic(err)
			}
			activityDate = d.Format("2006-01-02")

			sql := fmt.Sprintf("INSERT INTO activities (name, activity_date, activity_type_id, owner_id, duration, distance, unit) VALUES (\"\", \"%v\", \"%v\", 1, \"%v\",  \"%v\", \"miles\");", activityDate, activityTypeID, duration, distance)
			fmt.Println(sql)
		}
	}
}
