package service

import (
	"errors"
	"server/internal/models"
	"server/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetActivityTypes(t *testing.T) {
	// Arrange
	mockDB := new(mocks.PhobosDB)
	svc := New(mockDB)
	mockDB.On("GetActivityTypes").Return([]models.ActivityType{{ID: 1, Name: "Run"}}, nil)

	// Act
	result, err := svc.GetActivityTypes()

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)

	want := &[]models.ActivityType{{ID: 1, Name: "Run"}}

	assert.Equal(t, want, result)
}

func TestGetActivityTypesWithErr(t *testing.T) {
	// Arrange
	mockDB := new(mocks.PhobosDB)
	svc := New(mockDB)
	mockDB.On("GetActivityTypes").Return(nil, errors.New("oh dear"))

	// Act
	result, err := svc.GetActivityTypes()

	// Assert
	assert.Nil(t, result)
	assert.EqualError(t, err, "oh dear")
}
