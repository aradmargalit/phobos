package service

import (
	"server/internal/models"
)

// AddQuickAdd adds a new activity to the database
func (svc *service) AddQuickAdd(uid int, quickAdd *models.QuickAdd) (*models.QuickAdd, error) {

	quickAdd.OwnerID = uid
	record, err := svc.db.InsertQuickAdd(quickAdd)
	if err != nil {
		return nil, err
	}
	return record, nil

}

// GetQuickAdds returns all the user's quick-adds
func (svc *service) GetQuickAdds(uid int) (*[]models.QuickAdd, error) {
	qa, err := svc.db.GetQuickAddsByUser(uid)
	if err != nil {
		return nil, err
	}

	return &qa, nil
}

// DeleteQuickAdd deletes a quick add by the user ID and quick add ID
func (svc *service) DeleteQuickAdd(uid int, quickAddID int) error {
	return svc.db.DeleteQuickAddByID(uid, quickAddID)
}
