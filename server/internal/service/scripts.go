package service

import (
	"server/utils"

	"github.com/sirupsen/logrus"
)

// BackfillMeters adds the correct "meters" value to every activity entry
func (svc *service) BackfillMeters() {
	activities, err := svc.db.GetAllActivities()
	if err != nil {
		logrus.Fatal("Could not get activities")
	}

	// TODO: Batch these, it's really slow to issue an update per record
	for _, activity := range activities {
		if activity.Distance == 0 {
			continue
		}

		activity.Meters = utils.DistanceToMeters(activity.Distance, activity.Unit)
		logrus.Infof("Updating activity: %v with distance : %v %v and meters: %v...\n", activity.ID, activity.Distance, activity.Unit, activity.Meters)
		_, err := svc.db.UpdateActivity(&activity)
		if err != nil {
			logrus.Error(err)
		}
	}
}
