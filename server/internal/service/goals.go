package service

import "server/internal/models"

// AddGoal adds a new user workout goal to the database
func (svc *service) AddGoal(uid int, goal *models.Goal) (*models.Goal, error) {
	goal.UserID = uid
	record, err := svc.db.InsertGoal(goal)
	if err != nil {
		return nil, err
	}

	return record, nil
}

// GetGoals returns all the user's quick-adds
func (svc *service) GetGoals(uid int) (*[]models.Goal, error) {
	goals, err := svc.db.GetGoalsByUser(uid)
	if err != nil {
		return nil, err
	}

	return &goals, nil
}
