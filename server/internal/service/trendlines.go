package service

import "server/utils"

// GetTrendPoints returns an array of minutes worked out for each possible interval
func (svc *service) GetTrendPoints(uid int, interval string, utcOffset int) (*[]float64, error) {
	activities, err := svc.db.GetActivitiesByUser(uid)
	if err != nil {
		return nil, err
	}

	trend := utils.CalculateLastTenDays(activities, utcOffset)

	return &trend, nil
}
