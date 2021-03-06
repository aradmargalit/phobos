// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "server/internal/models"

	mock "github.com/stretchr/testify/mock"

	responsetypes "server/internal/responsetypes"
)

// PhobosDB is an autogenerated mock type for the PhobosDB type
type PhobosDB struct {
	mock.Mock
}

// DeleteActivityByID provides a mock function with given fields: _a0, _a1
func (_m *PhobosDB) DeleteActivityByID(_a0 int, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAllActivityTypes provides a mock function with given fields:
func (_m *PhobosDB) DeleteAllActivityTypes() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteGoalByID provides a mock function with given fields: uid, GoalID
func (_m *PhobosDB) DeleteGoalByID(uid int, GoalID int) error {
	ret := _m.Called(uid, GoalID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(uid, GoalID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteQuickAddByID provides a mock function with given fields: uid, quickAddID
func (_m *PhobosDB) DeleteQuickAddByID(uid int, quickAddID int) error {
	ret := _m.Called(uid, quickAddID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(uid, quickAddID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteStravaTokenByUserID provides a mock function with given fields: uid
func (_m *PhobosDB) DeleteStravaTokenByUserID(uid int) error {
	ret := _m.Called(uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetActivitiesByUser provides a mock function with given fields: _a0
func (_m *PhobosDB) GetActivitiesByUser(_a0 int) ([]models.ActivityResponse, error) {
	ret := _m.Called(_a0)

	var r0 []models.ActivityResponse
	if rf, ok := ret.Get(0).(func(int) []models.ActivityResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ActivityResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActivityByID provides a mock function with given fields: _a0
func (_m *PhobosDB) GetActivityByID(_a0 int) (models.Activity, error) {
	ret := _m.Called(_a0)

	var r0 models.Activity
	if rf, ok := ret.Get(0).(func(int) models.Activity); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Activity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActivityByStravaID provides a mock function with given fields: _a0
func (_m *PhobosDB) GetActivityByStravaID(_a0 *int) (models.Activity, error) {
	ret := _m.Called(_a0)

	var r0 models.Activity
	if rf, ok := ret.Get(0).(func(*int) models.Activity); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Activity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActivityTypeIDByStravaType provides a mock function with given fields: _a0
func (_m *PhobosDB) GetActivityTypeIDByStravaType(_a0 string) (int, error) {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActivityTypes provides a mock function with given fields:
func (_m *PhobosDB) GetActivityTypes() ([]models.ActivityType, error) {
	ret := _m.Called()

	var r0 []models.ActivityType
	if rf, ok := ret.Get(0).(func() []models.ActivityType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ActivityType)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllActivities provides a mock function with given fields:
func (_m *PhobosDB) GetAllActivities() ([]models.Activity, error) {
	ret := _m.Called()

	var r0 []models.Activity
	if rf, ok := ret.Get(0).(func() []models.Activity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGoalByID provides a mock function with given fields: id
func (_m *PhobosDB) GetGoalByID(id int) (models.Goal, error) {
	ret := _m.Called(id)

	var r0 models.Goal
	if rf, ok := ret.Get(0).(func(int) models.Goal); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Goal)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGoalsByUser provides a mock function with given fields: uid
func (_m *PhobosDB) GetGoalsByUser(uid int) ([]models.Goal, error) {
	ret := _m.Called(uid)

	var r0 []models.Goal
	if rf, ok := ret.Get(0).(func(int) []models.Goal); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Goal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQuickAddByID provides a mock function with given fields: id
func (_m *PhobosDB) GetQuickAddByID(id int) (models.QuickAdd, error) {
	ret := _m.Called(id)

	var r0 models.QuickAdd
	if rf, ok := ret.Get(0).(func(int) models.QuickAdd); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.QuickAdd)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQuickAddsByUser provides a mock function with given fields: uid
func (_m *PhobosDB) GetQuickAddsByUser(uid int) ([]models.QuickAdd, error) {
	ret := _m.Called(uid)

	var r0 []models.QuickAdd
	if rf, ok := ret.Get(0).(func(int) []models.QuickAdd); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.QuickAdd)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStravaTokenByUserID provides a mock function with given fields: uid
func (_m *PhobosDB) GetStravaTokenByUserID(uid int) (models.StravaToken, error) {
	ret := _m.Called(uid)

	var r0 models.StravaToken
	if rf, ok := ret.Get(0).(func(int) models.StravaToken); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(models.StravaToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: _a0
func (_m *PhobosDB) GetUserByEmail(_a0 string) (models.User, error) {
	ret := _m.Called(_a0)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: _a0
func (_m *PhobosDB) GetUserByID(_a0 int) (responsetypes.User, error) {
	ret := _m.Called(_a0)

	var r0 responsetypes.User
	if rf, ok := ret.Get(0).(func(int) responsetypes.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(responsetypes.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserIDByStravaID provides a mock function with given fields: stravaID
func (_m *PhobosDB) GetUserIDByStravaID(stravaID int) (int, error) {
	ret := _m.Called(stravaID)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(stravaID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(stravaID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertActivity provides a mock function with given fields: _a0
func (_m *PhobosDB) InsertActivity(_a0 *models.Activity) (*models.Activity, error) {
	ret := _m.Called(_a0)

	var r0 *models.Activity
	if rf, ok := ret.Get(0).(func(*models.Activity) *models.Activity); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Activity) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertActivityType provides a mock function with given fields: _a0
func (_m *PhobosDB) InsertActivityType(_a0 models.ActivityType) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.ActivityType) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertGoal provides a mock function with given fields: a
func (_m *PhobosDB) InsertGoal(a *models.Goal) (*models.Goal, error) {
	ret := _m.Called(a)

	var r0 *models.Goal
	if rf, ok := ret.Get(0).(func(*models.Goal) *models.Goal); ok {
		r0 = rf(a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Goal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Goal) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertQuickAdd provides a mock function with given fields: a
func (_m *PhobosDB) InsertQuickAdd(a *models.QuickAdd) (*models.QuickAdd, error) {
	ret := _m.Called(a)

	var r0 *models.QuickAdd
	if rf, ok := ret.Get(0).(func(*models.QuickAdd) *models.QuickAdd); ok {
		r0 = rf(a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.QuickAdd)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.QuickAdd) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertStravaToken provides a mock function with given fields: tok
func (_m *PhobosDB) InsertStravaToken(tok models.StravaToken) (models.StravaToken, error) {
	ret := _m.Called(tok)

	var r0 models.StravaToken
	if rf, ok := ret.Get(0).(func(models.StravaToken) models.StravaToken); ok {
		r0 = rf(tok)
	} else {
		r0 = ret.Get(0).(models.StravaToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.StravaToken) error); ok {
		r1 = rf(tok)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUser provides a mock function with given fields: _a0
func (_m *PhobosDB) InsertUser(_a0 models.User) (*responsetypes.User, error) {
	ret := _m.Called(_a0)

	var r0 *responsetypes.User
	if rf, ok := ret.Get(0).(func(models.User) *responsetypes.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*responsetypes.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateActivity provides a mock function with given fields: _a0
func (_m *PhobosDB) UpdateActivity(_a0 *models.Activity) (*models.Activity, error) {
	ret := _m.Called(_a0)

	var r0 *models.Activity
	if rf, ok := ret.Get(0).(func(*models.Activity) *models.Activity); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Activity) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGoal provides a mock function with given fields: _a0
func (_m *PhobosDB) UpdateGoal(_a0 *models.Goal) (*models.Goal, error) {
	ret := _m.Called(_a0)

	var r0 *models.Goal
	if rf, ok := ret.Get(0).(func(*models.Goal) *models.Goal); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Goal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Goal) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStravaToken provides a mock function with given fields: tok
func (_m *PhobosDB) UpdateStravaToken(tok models.StravaToken) (models.StravaToken, error) {
	ret := _m.Called(tok)

	var r0 models.StravaToken
	if rf, ok := ret.Get(0).(func(models.StravaToken) models.StravaToken); ok {
		r0 = rf(tok)
	} else {
		r0 = ret.Get(0).(models.StravaToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.StravaToken) error); ok {
		r1 = rf(tok)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
