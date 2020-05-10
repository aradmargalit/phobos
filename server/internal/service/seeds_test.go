package service

import (
	"errors"
	"server/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Confirm we don't reseed if the delete fails
func TestSeedActivityTypesWithFailedDelete(t *testing.T) {
	// Arrange
	mockDB := new(mocks.PhobosDB)
	mockDB.On("DeleteAllActivityTypes").Return(errors.New("failed to delete"))
	svc := New(mockDB)

	// Act
	err := svc.SeedActivityTypes()

	// Assert
	assert.EqualError(t, err, "failed to delete")
}

// Confirm we insert the right number of activities
func TestSeedActivityTypes(t *testing.T) {
	// Arrange
	mockDB := new(mocks.PhobosDB)
	mockDB.On("DeleteAllActivityTypes").Return(nil)
	mockDB.On("InsertActivityType", mock.Anything).Return(nil)

	svc := New(mockDB)

	// Act
	err := svc.SeedActivityTypes()

	// Assert
	assert.NoError(t, err)
	mockDB.AssertNumberOfCalls(t, "InsertActivityType", 37)
}
