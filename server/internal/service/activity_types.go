package service

import "server/internal/models"

// GetActivityTypes returns all available activity types
func (svc *service) GetActivityTypes() (*[]models.ActivityType, error) {
	at, err := svc.db.GetActivityTypes()
	if err != nil {
		return nil, err
	}
	return &at, nil
}
