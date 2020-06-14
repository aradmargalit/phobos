package service

import (
	"errors"
	"server/internal/models"
	"server/internal/responsetypes"
	"server/mocks"
	"server/testdata"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddActivity(t *testing.T) {
	// Arrange
	dbActivity := testdata.GetTestActivity()
	postActivity := testdata.GetTestPostActivityRequest()

	// Set the time to be an RFC3339 time
	activityDate, _ := time.Parse("2006-01-02", postActivity.ActivityDate)
	postActivity.ActivityDate = activityDate.Format(time.RFC3339)
	dbActivity.ActivityDate = activityDate.Format(time.RFC3339)

	// Create a service that we'll use for testing, but with a mockDB
	mockDB := new(mocks.PhobosDB)
	mockDB.On("InsertActivity", mock.AnythingOfType("*models.Activity")).Return(dbActivity, nil)
	svc := New(mockDB)

	// Assert that this'll fail before acting
	assert.NotEqual(t, postActivity.ActivityDate, testdata.GetTestActivity().ActivityDate)

	// Act
	result, err := svc.AddActivity(postActivity, 1)

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, postActivity.ActivityDate, result.ActivityDate)

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}

func TestAddActivityReturnsError(t *testing.T) {
	// Arrange
	inputActivity := testdata.GetTestPostActivityRequest()

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
	// Arrange
	userID := 1
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
		assert.Equal(t, activity.LogicalIndex, len(*result)-idx)
	}

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}

func TestGetActivitiesWithError(t *testing.T) {
	// Arrange
	userID := 1
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

func TestDeleteActivity(t *testing.T) {
	// Arrange
	activityID := 1
	userID := 2
	mockDB := new(mocks.PhobosDB)
	mockDB.On("DeleteActivityByID", activityID, userID).Return(nil)

	svc := New(mockDB)

	// Act
	err := svc.DeleteActivity(userID, activityID)

	// Assert that if the DB throws no errors, neither do we
	assert.NoError(t, err)

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}

func TestDeleteActivityErrorWithWrongUserID(t *testing.T) {
	// Arrange
	activityID := 1
	userID := 2
	mockDB := new(mocks.PhobosDB)
	mockDB.On("DeleteActivityByID", activityID, userID).Return(errors.New("oh dear"))

	svc := New(mockDB)

	// Act
	err := svc.DeleteActivity(userID, activityID)

	// Assert that if the DB throws no errors, neither do we
	assert.Error(t, err)

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}

func TestGetIntervalSummary(t *testing.T) {
	// Arrange
	userID := 1
	mockDB := new(mocks.PhobosDB)
	mockDB.On("GetActivitiesByUser", userID).Return(testdata.GetActivityResponses(20, 24), nil)

	svc := New(mockDB)

	// Act
	result, err := svc.GetIntervalSummary(userID, "month", 0)

	// Assert that if the DB throws no errors, neither do we
	assert.NotNil(t, result)
	assert.NoError(t, err)

	want := responsetypes.IntervalSum{
		Interval:         "January 2001",    // Our generator only creates 20 days in January
		Duration:         20,                // Each activity is 1 minute (x20 => 20min)
		Miles:            20,                // Each activity is 1 mile (x20 => 20 miles)
		PercentageActive: 70.37037037037037, // January 2001 has 31 days, but our first activity was on the second, so we "skipped" the 1st
	}

	assert.Equal(t, want, (*result)[0])

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}

func TestGetIntervalSummaryWeekly(t *testing.T) {
	// Arrange
	userID := 1
	mockDB := new(mocks.PhobosDB)
	mockDB.On("GetActivitiesByUser", userID).Return(testdata.GetActivityResponses(22, 24), nil)

	svc := New(mockDB)

	// Act
	result, err := svc.GetIntervalSummary(userID, "week", 0)

	// Assert that if the DB throws no errors, neither do we
	assert.NotNil(t, result)
	assert.NoError(t, err)

	want := []responsetypes.IntervalSum{
		// Despite only working out 6 days, we "started" on the 2nd, meaning it wasn't skipped
		{Interval: "2001, week 4 (Jan)", Duration: 4, Miles: 4, PercentageActive: 57.14285714285714},
		{Interval: "2001, week 3 (Jan)", Duration: 7, Miles: 7, PercentageActive: 100},
		{Interval: "2001, week 2 (Jan)", Duration: 7, Miles: 7, PercentageActive: 100},
		{Interval: "2001, week 1 (Jan)", Duration: 4, Miles: 4, PercentageActive: 100},
	}

	assert.Equal(t, want[0], (*result)[0])
	assert.Equal(t, want[1], (*result)[1])
	assert.Equal(t, want[2], (*result)[2])
	assert.Equal(t, want[3], (*result)[3])

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}

func TestGetIntervalSummaryYearly(t *testing.T) {
	// Arrange
	userID := 1
	mockDB := new(mocks.PhobosDB)
	mockDB.On("GetActivitiesByUser", userID).Return(testdata.GetActivityResponses(20, 24), nil)

	svc := New(mockDB)

	// Act
	result, err := svc.GetIntervalSummary(userID, "year", 0)

	// Assert that if the DB throws no errors, neither do we
	assert.NotNil(t, result)
	assert.NoError(t, err)

	want := []responsetypes.IntervalSum{
		// 364 days (skipped the 1st) - 20 = 344
		{Interval: "2001", Duration: 20, Miles: 20, PercentageActive: 5.263157894736842},
	}

	assert.Equal(t, want[0], (*result)[0])

	// Sanity check: assert that our mock did everything we thought it would
	mockDB.AssertExpectations(t)
}

func TestGetIntervalSummaryInvalidInterval(t *testing.T) {
	// Arrange
	svc := New(new(mocks.PhobosDB))

	// Act
	_, err := svc.GetIntervalSummary(1, "blorgus", 0)

	// Assert
	assert.EqualError(t, err, "interval must be week, month, or year")
}

func TestGetIntervalSummaryNoActivities(t *testing.T) {
	// Arrange
	userID := 1
	mockDB := new(mocks.PhobosDB)
	mockDB.On("GetActivitiesByUser", userID).Return([]models.ActivityResponse{}, nil)

	svc := New(mockDB)

	// Act
	result, err := svc.GetIntervalSummary(userID, "year", 0)

	// Assert
	want := &[]responsetypes.IntervalSum{}

	assert.Equal(t, want, result)
	assert.NoError(t, err)
}

func TestGetIntervalSummaryDBError(t *testing.T) {
	// Arrange
	userID := 1
	mockDB := new(mocks.PhobosDB)
	mockDB.On("GetActivitiesByUser", userID).Return(nil, errors.New("oh no please no"))

	svc := New(mockDB)

	// Act
	_, err := svc.GetIntervalSummary(userID, "year", 0)

	// Assert
	assert.EqualError(t, err, "oh no please no")
}
