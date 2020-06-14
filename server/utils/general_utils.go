package utils

import (
	"math"
	"time"
)

// Retry will retry the func "fn" as many times as requested
func Retry(fn func() error, tries int, duration time.Duration) (err error) {
	for i := 0; i < tries; i++ {
		if err = fn(); err == nil {
			return
		}
		time.Sleep(duration)
	}
	return
}

// FloatTwoDecimals takes a float and returns a float rounded to 2 decimal points
func FloatTwoDecimals(inFloat float64) float64 {
	return math.Round(inFloat*100) / 100
}
