package service

import (
	"errors"
	"fmt"
	"server/utils"
)

// GetTrendPoints returns an array of minutes worked out for each possible lookback window
func (svc *service) GetTrendPoints(uid int, lookback string, utcOffset int) (*[]float64, error) {
	activities, err := svc.db.GetActivitiesByUser(uid)
	if err != nil {
		return nil, err
	}
	fmt.Println(lookback)

	// We'll use a different utility function for each of these
	switch lookback {
	case "l10": // Last 10
		return utils.CalculateLastNDays(activities, utcOffset, 10), nil
	case "l7": // Last 7
		return utils.CalculateLastNDays(activities, utcOffset, 7), nil
	case "lw": // Last Week
		// For this, we need to use a slightly different approach than counting, since there can be sparse days
		return utils.CalculateLastWeek(activities, utcOffset), nil
	default:
		return nil, errors.New("lookback is invalid, must be one of [l10, l7, lw, lm]")
	}
}
