package testdata

import (
	"server/internal/models"
)

// GetTestActivity returns a boilerplate activity for testing
func GetTestActivity() *models.Activity {
	return &models.Activity{
		ID: 1,
		Name: "Fun activity",
		ActivityDate: "2020-01-05",
		ActivityTypeID: 1,
		OwnerID: 1,
		Duration: 1,
		Distance: 1,
		Unit: "miles",
		HeartRate: MakeIntPointer(200),
		StravaID: nil,
		CreatedAt: "2001-02-03",
		UpdatedAt: "2004-05-06",
	}
}

// GetTestPostActivityRequest returns a POST request object to use for testing
func GetTestPostActivityRequest() *models.PostActivityRequest {
	return &models.PostActivityRequest{
		ID: 1,
		Name: "Fun activity",
		ActivityDate: "2020-01-05",
		ActivityTypeID: 1,
		OwnerID: 1,
		Duration: 1,
		Distance: 1,
		Unit: "miles",
		HeartRate: 200,
		CreatedAt: "2001-02-03",
		UpdatedAt: "2004-05-06",
	}
}
