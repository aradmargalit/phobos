package utils

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var callCount int

func callCounter(shouldErr bool) func() error {
	return func() error {
		callCount++
		if shouldErr {
			return errors.New("Intentional error")
		}
		return nil
	}
}

func TestRetryWithNoTries(t *testing.T) {
	callCount = 0
	nanosec, _ := time.ParseDuration("1ns")
	err := Retry(callCounter(false), 0, nanosec)
	assert.NoError(t, err)
	assert.Equal(t, 0, callCount)
}

func TestRetryWithFiveTries(t *testing.T) {
	callCount = 0
	nanosec, _ := time.ParseDuration("1ns")
	err := Retry(callCounter(false), 5, nanosec)
	assert.NoError(t, err)
	assert.Equal(t, 1, callCount)
}

func TestRetryWithFiveFailedTries(t *testing.T) {
	callCount = 0
	nanosec, _ := time.ParseDuration("1ns")
	err := Retry(callCounter(true), 5, nanosec)
	assert.Error(t, err)
	assert.Equal(t, 5, callCount)
}
