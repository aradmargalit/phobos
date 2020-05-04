package service

import (
	"errors"
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

	// Create a service that we'll use for testing, but with a mockDB
	mockDB := new(mocks.PhobosDB)
	mockDB.On("InsertActivity", mock.AnythingOfType("*models.Activity")).Return(inputActivity, nil)
	svc := New(mockDB)
	
	// Assert that this'll fail before acting
	assert.NotEqual(t, inputActivity.ActivityDate, testdata.GetTestActivity().ActivityDate)

	// Act
	result, err := svc.AddActivity(inputActivity, 1)

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, inputActivity.ActivityDate, result.ActivityDate)

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}

func TestAddActivityReturnsError(t *testing.T) {
	// Arrange
	inputActivity := testdata.GetTestActivity()

	// Set the time to be an RFC3339 time
	activityDate, _ := time.Parse("2006-01-02", inputActivity.ActivityDate)
	inputActivity.ActivityDate = activityDate.Format(time.RFC3339)

	// Create a service that we'll use for testing, but with a mockDB
	mockDB := new(mocks.PhobosDB)
	mockDB.On("InsertActivity", mock.AnythingOfType("*models.Activity")).Return(nil, errors.New("Uh oh"))
	svc := New(mockDB)
	
	// Act
	result, err := svc.AddActivity(inputActivity, 1)

	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}

func TestUpdateActivity(t *testing.T) {
	// Arrange
	inputActivity := testdata.GetTestActivity()

	// Set the time to be an RFC3339 time
	activityDate, _ := time.Parse("2006-01-02", inputActivity.ActivityDate)
	inputActivity.ActivityDate = activityDate.Format(time.RFC3339)

	// Create a service that we'll use for testing, but with a mockDB
	mockDB := new(mocks.PhobosDB)
	mockDB.On("UpdateActivity", mock.AnythingOfType("*models.Activity")).Return(inputActivity, nil)
	svc := New(mockDB)
	
	// Assert that this'll fail before acting
	assert.NotEqual(t, inputActivity.ActivityDate, testdata.GetTestActivity().ActivityDate)

	// Act
	result, err := svc.UpdateActivity(inputActivity)

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, inputActivity.ActivityDate, result.ActivityDate)

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}

func TestUpdateActivityReturnsError(t *testing.T) {
	// Arrange
	inputActivity := testdata.GetTestActivity()

	// Set the time to be an RFC3339 time
	activityDate, _ := time.Parse("2006-01-02", inputActivity.ActivityDate)
	inputActivity.ActivityDate = activityDate.Format(time.RFC3339)

	// Create a service that we'll use for testing, but with a mockDB
	mockDB := new(mocks.PhobosDB)
	mockDB.On("UpdateActivity", mock.AnythingOfType("*models.Activity")).Return(nil, errors.New("Uh oh"))
	svc := New(mockDB)
	
	// Act
	result, err := svc.UpdateActivity(inputActivity)

	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}

func TestGetActivities(t *testing.T) {
	userID := 1
	// Arrange
	mockDB := new(mocks.PhobosDB)
	mockDB.On("GetActivitiesByUser", userID).Return(testdata.GetActivityResponses(10, 24), nil)
	
	svc := New(mockDB)
	
	// Act
	result, err := svc.GetActivities(userID)

	// Assert basic integrity of the response
	assert.NotNil(t, result)
	assert.NoError(t, err)

	// Assert that the activities have an embedded activity type
	assert.NotEmpty(t, (*result)[0].ActivityType.Name)

	// Assert that the activities have a logical index that makes sense
	for idx, activity := range *result {
		assert.Equal(t, activity.LogicalIndex, len(*result) - idx)
	}

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}

func TestGetActivitiesWithError(t *testing.T) {
	userID := 1
	// Arrange
	mockDB := new(mocks.PhobosDB)
	mockDB.On("GetActivitiesByUser", userID).Return(nil, errors.New("Oh dear"))
	
	svc := New(mockDB)
	
	// Act
	result, err := svc.GetActivities(userID)

	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}