package testdata

import (
	"database/sql"
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
		StravaID: sql.NullInt64{},
		CreatedAt: "2001-02-03",
		UpdatedAt: "2004-05-06",
	}
}
