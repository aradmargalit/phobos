package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const layout = "2006-01-02 15:04:05"

func TestDateEqual(t *testing.T) {
	// Assert the same date is...the same
	d := time.Now()
	assert.True(t, DateEqual(d, d))

	// Assert that within the same day, we still get a match
	fiveS, _ := time.ParseDuration("5s")
	d2 := d.Add(fiveS)
	assert.True(t, DateEqual(d, d2))

	d3, _ := time.Parse(layout, "2020-01-02 10:11:12")
	d4, _ := time.Parse(layout, "2020-01-02 20:12:11")
	assert.True(t, DateEqual(d3, d4))
}

func TestDatesUnequal(t *testing.T) {
	d := time.Now()

	// Assert that across days we don't get a match
	twoDays, _ := time.ParseDuration("48h")
	d2 := d.Add(twoDays)
	assert.False(t, DateEqual(d, d2))

	d3, _ := time.Parse(layout, "2020-01-02 10:11:12")
	d4, _ := time.Parse(layout, "2022-01-02 10:11:12")
	assert.False(t, DateEqual(d3, d4))
}

func TestRoundToDay(t *testing.T) {
	rawDate, _ := time.Parse(layout, "2020-01-02 10:11:12")
	want, _ := time.Parse(layout, "2020-01-02 00:00:00")
	assert.Equal(t, want, RoundTimeToDay(rawDate))
}
