package service

import (
	"server/mocks"
	"server/testdata"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddActivity(t *testing.T) {
	// Arrange
	inputActivity := testdata.GetTestActivity()

	// Set the time to be an RFC3339 time
	activityDate, _ := time.Parse("2006-01-02", inputActivity.ActivityDate)
	inputActivity.ActivityDate = activityDate.Format(time.RFC3339)

	// Create a service that we'll use for testing, but with a mockD
	mockDB := new(mocks.PhobosDB)
	mockDB.On("InsertActivity", mock.AnythingOfType("*models.Activity")).Return(inputActivity, nil)
	svc := New(mockDB)
	
	// Act
	result, err := svc.AddActivity(inputActivity, 1)

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, inputActivity.ActivityDate, result.ActivityDate)

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}
