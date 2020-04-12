package utils

import (
	"time"
)

// DateEqual tells you if two time objects are on the same date
// Credit where credit is due: https://stackoverflow.com/a/21069048
func DateEqual(date1 time.Time, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// RoundTimeToDay takes in a time and returns a time with just that date
// as a unique ID for that day in history
func RoundTimeToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
