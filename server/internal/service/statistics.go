package service

import (
	"server/internal/responsetypes"
	"server/utils"
)

// GetUserStatistics returns some fun user statistics for the frontend
func (svc *service) GetUserStatistics(uid int, offset int) (*responsetypes.Stats, error) {
	a, err := svc.db.GetActivitiesByUser(uid)
	if err != nil {
		return nil, err
	}

	// Now that we have activities, let's cronch the numbies
	response := responsetypes.Stats{
		Workouts:      len(a),
		Hours:         utils.CalculateTotalHours(a),
		Miles:         utils.CalculateMileage(a),
		TypeBreakdown: utils.CalculateTypeBreakdown(a),
		DayBreakdown:  utils.CalculateDayBreakdown(a),
	}

	return &response, nil
}
