package service

import (
	"server/internal/responsetypes"
)

// GetActivityTypes returns all available activity types
func (svc *service) GetActivityTypes() (*[]responsetypes.ActivityType, error) {
	at, err := svc.db.GetActivityTypes()
	if err != nil {
		return nil, err
	}
	return &at, nil
}
